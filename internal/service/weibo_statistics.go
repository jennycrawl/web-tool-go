package service

import (
    "context"
    "fmt"
    "time"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/repository"
)

type WeiboStatisticsService interface {
    GetWeiboStatisticsList(ctx context.Context, accountId int, startTime time.Time, endTime time.Time) (*v1.GetWeiboStatisticsListResponseData, error)
}

func NewWeiboStatisticsService(
    service *Service,
    weiboStatisticsRepository repository.WeiboStatisticsRepository,
) WeiboStatisticsService {
    return &weiboStatisticsService{
        Service:                   service,
        weiboStatisticsRepository: weiboStatisticsRepository,
    }
}

type weiboStatisticsService struct {
    *Service
    weiboStatisticsRepository repository.WeiboStatisticsRepository
}

func (s *weiboStatisticsService) GetWeiboStatisticsList(ctx context.Context, accountId int, startTime time.Time, endTime time.Time) (*v1.GetWeiboStatisticsListResponseData, error) {
    weiboStatisticsList, err := s.weiboStatisticsRepository.GetWeiboStatisticsList(ctx, accountId, startTime, endTime)
    if err != nil {
        return nil, err
    }
    var responseData v1.GetWeiboStatisticsListResponseData
    for _, weiboStatistics := range *weiboStatisticsList {
        responseData = append(responseData, v1.WeiboStatistics{
            Name:       weiboStatistics.Name,
            ID:         weiboStatistics.ID,
            Attention:  weiboStatistics.Attention,
            Fans:       weiboStatistics.Fans,
            Feed:       weiboStatistics.Feed,
            CrawlTime:  weiboStatistics.UpdateTime.Format("2006-01-02 15:04:05"),
            Count:      weiboStatistics.Count,
            ForwardSum: fmt.Sprintf("%d", weiboStatistics.ForwardSum),
            CommentSum: fmt.Sprintf("%d", weiboStatistics.CommentSum),
            LikeSum:    fmt.Sprintf("%d", weiboStatistics.LikeSum),
            ForwardAvg: fmt.Sprintf("%.4f", weiboStatistics.ForwardAvg),
            CommentAvg: fmt.Sprintf("%.4f", weiboStatistics.CommentAvg),
            LikeAvg:    fmt.Sprintf("%.4f", weiboStatistics.LikeAvg),
            ForwardMax: weiboStatistics.ForwardMax,
            CommentMax: weiboStatistics.CommentMax,
            LikeMax:    weiboStatistics.LikeMax,
            ForwardMin: weiboStatistics.ForwardMin,
            CommentMin: weiboStatistics.CommentMin,
            LikeMin:    weiboStatistics.LikeMin,
        })
    }
    return &responseData, nil
}
