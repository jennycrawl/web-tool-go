package repository

import (
    "context"
    "errors"
    "gorm.io/gorm"
    "time"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/model"
)

type WeiboStatisticsRepository interface {
    GetWeiboStatisticsList(ctx context.Context, accountId int, startTime time.Time, endTime time.Time) (*[]model.WeiboStatistics, error)
}

func NewWeiboStatisticsRepository(
    repository *Repository,
) WeiboStatisticsRepository {
    return &weiboStatisticsRepository{
        Repository: repository,
    }
}

type weiboStatisticsRepository struct {
    *Repository
}

func (r *weiboStatisticsRepository) GetWeiboStatisticsList(ctx context.Context, accountId int, startTime time.Time, endTime time.Time) (*[]model.WeiboStatistics, error) {
    var weiboStatisticsList []model.WeiboStatistics
    modelMsg := model.WeiboMsg{}
    db := r.DB(ctx).Table(modelMsg.TableName())
    db.Select("weibo_account.name, weibo_account.id, weibo_account.attention, weibo_account.fans, weibo_account.feed, weibo_account.update_time," +
        "COUNT(weibo_feed.id) AS count," +
        "SUM(weibo_feed.forward) AS forward_sum," +
        "SUM(weibo_feed.comment) AS comment_sum," +
        "SUM(weibo_feed.like) AS like_sum," +
        "AVG(weibo_feed.forward) AS forward_avg," +
        "AVG(weibo_feed.comment) AS comment_avg," +
        "AVG(weibo_feed.like) AS like_avg," +
        "MAX(weibo_feed.forward) AS forward_max," +
        "MAX(weibo_feed.comment) AS comment_max," +
        "MAX(weibo_feed.like) AS like_max," +
        "MIN(weibo_feed.forward) AS forward_min," +
        "MIN(weibo_feed.comment) AS comment_min," +
        "MIN(weibo_feed.like) AS like_min")
    db.Joins("JOIN weibo_account ON weibo_account.id = weibo_feed.account_id")
    db.Where("weibo_account.status = ?", model.WeiboAccountStatusValid)
    if accountId > 0 {
        db.Where("account_id = ?", accountId)
    }
    if !startTime.IsZero() {
        db.Where("pubtime >= ?", startTime.Format("2006-01-02 15:04:05"))
    }
    if !endTime.IsZero() {
        db.Where("pubtime < ?", endTime.Format("2006-01-02 15:04:05"))
    }
    db.Group("weibo_feed.account_id").Order("weibo_feed.id asc")

    if err := db.Scan(&weiboStatisticsList).Error; err != nil {
        r.logger.Error(err.Error())
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, v1.ErrNotFound
        }
        return nil, err
    }

    return &weiboStatisticsList, nil
}
