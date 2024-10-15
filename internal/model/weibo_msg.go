package model

import "time"

type WeiboMsg struct {
    ID         int          `gorm:"primaryKey;autoIncrement"`                            // 主键，自增
    Mid        string       `gorm:"size:255;not null"`                                   // 微博 ID
    AccountID  int          `gorm:"not null;index"`                                      // 账号 ID
    Forward    int          `gorm:"not null;default:0;index"`                            // 转发数
    Comment    int          `gorm:"not null;default:0;index"`                            // 评论数
    Like       int          `gorm:"not null;default:0;index"`                            // 点赞数
    Pubtime    time.Time    `gorm:"not null"`                                            // 发布时间
    CreateTime time.Time    `gorm:"default:current_timestamp();not null"`                // 创建时间
    UpdateTime time.Time    `gorm:"default:current_timestamp();not null;autoUpdateTime"` // 更新时间
    Account    WeiboAccount `gorm:"foreignKey:AccountID"`                                // 反向关联
}

func (m *WeiboMsg) TableName() string {
    return "weibo_feed"
}
