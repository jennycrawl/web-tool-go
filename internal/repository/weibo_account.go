package repository

import (
    "context"
    "errors"
    "gorm.io/gorm"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/model"
)

type WeiboAccountRepository interface {
    GetWeiboAccountList(ctx context.Context) (*[]model.WeiboAccount, error)
    GetWeiboAccountByID(ctx context.Context, id int) (*model.WeiboAccount, error)
    GetWeiboAccountByUid(ctx context.Context, uid string) (*model.WeiboAccount, error)
    CreateWeiboAccount(ctx context.Context, weiboAccount *model.WeiboAccount) error
    UpdateWeiboAccount(ctx context.Context, weiboAccount *model.WeiboAccount) error
    DeleteWeiboAccount(ctx context.Context, weiboAccount *model.WeiboAccount) error
}

func NewWeiboAccountRepository(
    repository *Repository,
) WeiboAccountRepository {
    return &weiboAccountRepository{
        Repository: repository,
    }
}

type weiboAccountRepository struct {
    *Repository
}

func (r *weiboAccountRepository) GetWeiboAccountList(ctx context.Context) (*[]model.WeiboAccount, error) {
    var weiboAccountList []model.WeiboAccount
    db := r.DB(ctx)
    db = db.Where("status = ?", model.WeiboAccountStatusValid)
    if err := db.Find(&weiboAccountList).Error; err != nil {
        r.logger.Debug(err.Error())
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, v1.ErrNotFound
        }
        return nil, err
    }
    return &weiboAccountList, nil
}

func (r *weiboAccountRepository) GetWeiboAccountByID(ctx context.Context, id int) (*model.WeiboAccount, error) {
    var weiboAccount model.WeiboAccount
    if err := r.DB(ctx).Where("id = ?", id).First(&weiboAccount).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, v1.ErrNotFound
        }
        return nil, err
    }
    return &weiboAccount, nil
}

func (r *weiboAccountRepository) GetWeiboAccountByUid(ctx context.Context, uid string) (*model.WeiboAccount, error) {
    var weiboAccount model.WeiboAccount
    if err := r.DB(ctx).Where("uid = ?", uid).First(&weiboAccount).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, v1.ErrNotFound
        }
        return nil, err
    }
    return &weiboAccount, nil
}

func (r *weiboAccountRepository) CreateWeiboAccount(ctx context.Context, weiboAccount *model.WeiboAccount) error {
    if err := r.DB(ctx).Create(weiboAccount).Error; err != nil {
        return err
    }
    return nil
}

func (r *weiboAccountRepository) UpdateWeiboAccount(ctx context.Context, weiboAccount *model.WeiboAccount) error {
    if err := r.DB(ctx).Save(weiboAccount).Error; err != nil {
        return err
    }
    return nil
}

func (r *weiboAccountRepository) DeleteWeiboAccount(ctx context.Context, weiboAccount *model.WeiboAccount) error {
    if err := r.DB(ctx).Delete(weiboAccount).Error; err != nil {
        return err
    }
    return nil
}
