package v1

type CreateWeiboAccountRequest struct {
    Name string `json:"name" binding:"required" example:"my name"`
    Uid  string `json:"uid" binding:"required" example:"123456"`
}

type UpdateWeiboAccountRequest CreateWeiboAccountRequest

type WeiboAccount struct {
    ID         int    `json:"id"`
    Name       string `json:"name"`
    Uid        string `json:"uid"`
    Status     int    `json:"status"`
    Attention  int    `json:"attention"`
    Fans       int    `json:"fans"`
    Feed       int    `json:"feed"`
    CreateTime string `json:"create_time"`
    UpdateTime string `json:"update_time"`
    Url        string `json:"url"`
    CrawlTime  string `json:"crawl_time"`
}

type GetWeiboAccountListResponseData []WeiboAccount

type GetWeiboAccountResponseData WeiboAccount

type GetWeiboAccountListResponse struct {
    Response
    Data GetWeiboAccountListResponseData
}
