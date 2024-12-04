package repository

import (
	"context"
	"errors"
	"time"
	v1 "web-tool-go/api/v1"
	"web-tool-go/internal/model"

	"gorm.io/gorm"
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
	db := r.DB(ctx).Table(modelMsg.TableName() + " as wf")
	db.Select("wa.name, wa.id, wa.attention, wa.fans, wa.feed, wa.update_time," +
		"COUNT(wf.id) AS count," +
		"SUM(wf.forward) AS forward_sum," +
		"SUM(wf.comment) AS comment_sum," +
		"SUM(wf.like) AS like_sum," +
		"AVG(wf.forward) AS forward_avg," +
		"AVG(wf.comment) AS comment_avg," +
		"AVG(wf.like) AS like_avg," +
		"MAX(wf.forward) AS forward_max," +
		"MAX(wf.comment) AS comment_max," +
		"MAX(wf.like) AS like_max," +
		"MIN(wf.forward) AS forward_min," +
		"MIN(wf.comment) AS comment_min," +
		"MIN(wf.like) AS like_min")
	db.Joins("JOIN weibo_account as wa ON wa.id = wf.account_id")
	db.Where("wa.status = ?", model.WeiboAccountStatusValid)
	if accountId > 0 {
		db.Where("wf.account_id = ?", accountId)
	}
	if !startTime.IsZero() {
		db.Where("wf.pubtime >= ?", startTime.Format("2006-01-02 15:04:05"))
	}
	if !endTime.IsZero() {
		db.Where("wf.pubtime < ?", endTime.Format("2006-01-02 15:04:05"))
	}
	db.Group("wa.id, wa.name, wa.attention, wa.fans, wa.feed, wa.update_time").Order("wa.id asc")

	if err := db.Scan(&weiboStatisticsList).Error; err != nil {
		r.logger.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}

	return &weiboStatisticsList, nil
}
