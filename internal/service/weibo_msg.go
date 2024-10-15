package service

import (
    "context"
    "fmt"
    "time"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/model"
    "web-tool-go/internal/repository"
)

type WeiboMsgService interface {
    GetWeiboMsgList(ctx context.Context, accountId int, startTime time.Time, endTime time.Time, sortField string, sortOrder string, page int, perPage int) (*v1.GetWeiboMsgListResponseData, error)
}

func NewWeiboMsgService(
    service *Service,
    weiboMsgRepository repository.WeiboMsgRepository,
) WeiboMsgService {
    return &weiboMsgService{
        Service:            service,
        weiboMsgRepository: weiboMsgRepository,
    }
}

type weiboMsgService struct {
    *Service
    weiboMsgRepository repository.WeiboMsgRepository
}

func (s *weiboMsgService) GetWeiboMsgList(ctx context.Context, accountId int, startTime time.Time, endTime time.Time, sortField string, sortOrder string, page int, perPage int) (*v1.GetWeiboMsgListResponseData, error) {
    paginator, err := s.weiboMsgRepository.GetWeiboMsgList(ctx, accountId, startTime, endTime, sortField, sortOrder, page, perPage)
    if err != nil {
        return nil, err
    }

    responseData := v1.GetWeiboMsgListResponseData{
        Total:       paginator.Total,
        TotalPage:   paginator.TotalPages,
        CurrentPage: paginator.CurrentPage,
        MsgList:     []v1.WeiboMsg{},
    }
    if data, ok := paginator.Data.([]model.WeiboMsg); ok {
        for _, msg := range data {
            responseData.MsgList = append(responseData.MsgList, v1.WeiboMsg{
                AccountName: msg.Account.Name,
                Id:          msg.ID,
                Mid:         msg.Mid,
                AccountId:   msg.AccountID,
                Forward:     msg.Forward,
                Comment:     msg.Comment,
                Like:        msg.Like,
                Pubtime:     msg.Pubtime.Format("2006-01-02"),
                CrawlTime:   msg.UpdateTime.Format("2006-01-02 15:04:05"),
                Url:         fmt.Sprintf("https://m.weibo.cn/detail/%s", msg.Mid),
            })
        }
    }
    return &responseData, nil
}
