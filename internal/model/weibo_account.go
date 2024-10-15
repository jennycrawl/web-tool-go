package model

import "time"

const (
    WeiboAccountStatusValid   = 1
    WeiboAccountStatusInvalid = 2
)

type WeiboAccount struct {
    ID         int        `gorm:"primaryKey;autoIncrement"`                                           // 主键，自增
    Name       string     `gorm:"type:varchar(255);not null"`                                         // 名称，非空
    Uid        string     `gorm:"type:varchar(255);not null"`                                         // 用户ID，非空
    Status     int        `gorm:"type:int unsigned;default:1;not null;comment:'1有效，0无效'"`         // 状态
    Attention  int        `gorm:"type:int unsigned;default:0;not null;comment:'关注数'"`              // 关注数
    Fans       int        `gorm:"type:int unsigned;default:0;not null;comment:'粉丝数'"`              // 粉丝数
    Feed       int        `gorm:"type:int unsigned;default:0;not null;comment:'微博数'"`              // 微博数
    CreateTime time.Time  `gorm:"type:timestamp;default:current_timestamp();not null"`                // 创建时间
    UpdateTime time.Time  `gorm:"type:timestamp;default:current_timestamp();not null;autoUpdateTime"` // 更新时间
    Msgs       []WeiboMsg `gorm:"foreignKey:AccountID"`                                               //一对多关联
}

func (m *WeiboAccount) TableName() string {
    return "weibo_account"
}
