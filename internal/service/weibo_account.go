package service

import (
    "context"
    "errors"
    "fmt"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/model"
    "web-tool-go/internal/repository"
)

type WeiboAccountService interface {
    GetWeiboAccountList(ctx context.Context) (*v1.GetWeiboAccountListResponseData, error)
    CreateWeiboAccount(ctx context.Context, req *v1.CreateWeiboAccountRequest) error
    UpdateWeiboAccount(ctx context.Context, id int, req *v1.UpdateWeiboAccountRequest) error
    DeleteWeiboAccountById(ctx context.Context, id int) error
}

func NewWeiboAccountService(
    service *Service,
    weiboAccountRepository repository.WeiboAccountRepository,
) WeiboAccountService {
    return &weiboAccountService{
        Service:                service,
        weiboAccountRepository: weiboAccountRepository,
    }
}

type weiboAccountService struct {
    *Service
    weiboAccountRepository repository.WeiboAccountRepository
}

func (s *weiboAccountService) GetWeiboAccountList(ctx context.Context) (*v1.GetWeiboAccountListResponseData, error) {
    accountList, err := s.weiboAccountRepository.GetWeiboAccountList(ctx)
    if err != nil {
        return nil, err
    }

    var responseData v1.GetWeiboAccountListResponseData
    for _, weiboAccount := range *accountList {
        responseData = append(responseData, v1.WeiboAccount{
            ID:         weiboAccount.ID,
            Name:       weiboAccount.Name,
            Uid:        weiboAccount.Uid,
            Status:     weiboAccount.Status,
            Attention:  weiboAccount.Attention,
            Fans:       weiboAccount.Fans,
            Feed:       weiboAccount.Feed,
            CreateTime: weiboAccount.CreateTime.Format("2006-01-02 15:04:05"),
            UpdateTime: weiboAccount.UpdateTime.Format("2006-01-02 15:04:05"),
            Url:        fmt.Sprintf("https://m.weibo.cn/u/%s", weiboAccount.Uid),
            CrawlTime:  weiboAccount.UpdateTime.Format("2006-01-02"),
        })
    }
    return &responseData, nil
}

func (s *weiboAccountService) CreateWeiboAccount(ctx context.Context, req *v1.CreateWeiboAccountRequest) error {
    // check uid
    weiboAccount, err := s.weiboAccountRepository.GetWeiboAccountByUid(ctx, req.Uid)
    if weiboAccount != nil {
        return v1.ErrWeiboAccountUidExist
    }
    if err == nil {
        return v1.ErrWeiboAccountUidExist
    } else if !errors.Is(err, v1.ErrNotFound) {
        return v1.ErrInternalServerError
    }

    weiboAccount = &model.WeiboAccount{
        Name: req.Name,
        Uid:  req.Uid,
    }
    if err = s.weiboAccountRepository.CreateWeiboAccount(ctx, weiboAccount); err != nil {
        return err
    }
    return nil
}

func (s *weiboAccountService) UpdateWeiboAccount(ctx context.Context, id int, req *v1.UpdateWeiboAccountRequest) error {
    weiboAccount, err := s.weiboAccountRepository.GetWeiboAccountByID(ctx, id)
    if err != nil {
        return err
    }
    weiboAccount.Name = req.Name
    weiboAccount.Uid = req.Uid

    if err = s.weiboAccountRepository.UpdateWeiboAccount(ctx, weiboAccount); err != nil {
        return err
    }

    return nil
}

func (s *weiboAccountService) DeleteWeiboAccountById(ctx context.Context, id int) error {
    weiboAccount, err := s.weiboAccountRepository.GetWeiboAccountByID(ctx, id)
    if err != nil {
        return err
    }

    if err = s.weiboAccountRepository.DeleteWeiboAccount(ctx, weiboAccount); err != nil {
        return err
    }

    return nil
}
