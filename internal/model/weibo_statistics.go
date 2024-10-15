package model

import (
    "time"
)

type WeiboStatistics struct {
    ID         int       `gorm:"primaryKey"`         // 主键
    Name       string    `gorm:"column:name"`        // 名称
    Attention  int       `gorm:"column:attention"`   // 关注者数
    Fans       int       `gorm:"column:fans"`        // 粉丝数
    Feed       int       `gorm:"column:feed"`        // 文章数
    UpdateTime time.Time `gorm:"column:update_time"` // 更新时间
    Count      int       `gorm:"column:count"`       // 抓取的文章数
    ForwardSum int64     `gorm:"column:forward_sum"` // 转发总数
    CommentSum int64     `gorm:"column:comment_sum"` // 评论总数
    LikeSum    int64     `gorm:"column:like_sum"`    // 点赞总数
    ForwardAvg float64   `gorm:"column:forward_avg"` // 转发平均数
    CommentAvg float64   `gorm:"column:comment_avg"` // 评论平均数
    LikeAvg    float64   `gorm:"column:like_avg"`    // 点赞平均数
    ForwardMax int       `gorm:"column:forward_max"` // 转发最大值
    CommentMax int       `gorm:"column:comment_max"` // 评论最大值
    LikeMax    int       `gorm:"column:like_max"`    // 点赞最大值
    ForwardMin int       `gorm:"column:forward_min"` // 转发最小值
    CommentMin int       `gorm:"column:comment_min"` // 评论最小值
    LikeMin    int       `gorm:"column:like_min"`    // 点赞最小值
}
