package v1

type WeiboStatistics struct {
    Name       string `json:"name"`
    ID         int    `json:"id"`
    Attention  int    `json:"attention"`
    Fans       int    `json:"fans"`
    Feed       int    `json:"feed"`
    CrawlTime  string `json:"crawl_time"`
    Count      int    `json:"count"`
    ForwardSum string `json:"forward_sum"` // 转发总数
    CommentSum string `json:"comment_sum"` // 评论总数
    LikeSum    string `json:"like_sum"`    // 点赞总数
    ForwardAvg string `json:"forward_avg"` // 转发平均数
    CommentAvg string `json:"comment_avg"` // 评论平均数
    LikeAvg    string `json:"like_avg"`    // 点赞平均数
    ForwardMax int    `json:"forward_max"` // 转发最大值
    CommentMax int    `json:"comment_max"` // 评论最大值
    LikeMax    int    `json:"like_max"`    // 点赞最大值
    ForwardMin int    `json:"forward_min"` // 转发最小值
    CommentMin int    `json:"comment_min"` // 评论最小值
    LikeMin    int    `json:"like_min"`    // 点赞最小值
}

type GetWeiboStatisticsListResponseData []WeiboStatistics

type GetWeiboStatisticsResponse struct {
    Response
    Data GetWeiboStatisticsListResponseData
}
