package repository

import (
    "context"
    "errors"
    "fmt"
    "gorm.io/gorm"
    "math"
    "time"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/model"
    "web-tool-go/internal/utils"
)

type WeiboMsgRepository interface {
    GetWeiboMsgList(ctx context.Context, accountId int, startTime time.Time, endTime time.Time, sortField string, sortOrder string, page int, perPage int) (*utils.Paginator, error)
}

func NewWeiboMsgRepository(
    repository *Repository,
) WeiboMsgRepository {
    return &weiboMsgRepository{
        Repository: repository,
    }
}

type weiboMsgRepository struct {
    *Repository
}

func (r *weiboMsgRepository) GetWeiboMsgList(ctx context.Context, accountId int, startTime time.Time, endTime time.Time, sortField string, sortOrder string, page int, perPage int) (*utils.Paginator, error) {
    var weiboMsgList []model.WeiboMsg
    db := r.DB(ctx).Model(&model.WeiboMsg{}).Preload("Account")
    if accountId > 0 {
        db.Where("account_id = ?", accountId)
    }
    if !startTime.IsZero() {
        db.Where("pubtime >= ?", startTime.Format("2006-01-02 15:04:05"))
    }
    if !endTime.IsZero() {
        db.Where("pubtime < ?", endTime.Format("2006-01-02 15:04:05"))
    }
    if sortField != "" {
        //加上转义，防止字段是mysql关键词的情况，比如like
        db.Order(fmt.Sprintf("`%s` %s", sortField, sortOrder))
    }
    var total int64
    db.Count(&total)
    if page > 0 && perPage > 0 {
        db.Offset((page - 1) * perPage).Limit(perPage)
    }
    if err := db.Find(&weiboMsgList).Error; err != nil {
        r.logger.Error(err.Error())
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, v1.ErrNotFound
        }
        return nil, err
    }
    paginator := utils.Paginator{
        CurrentPage: page,
        PerPage:     perPage,
        Total:       total,
        TotalPages:  int(math.Ceil(float64(total) / float64(perPage))),
        Data:        weiboMsgList,
    }
    return &paginator, nil
}
