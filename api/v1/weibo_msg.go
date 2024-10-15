package v1

type WeiboMsg struct {
    AccountName string `json:"account_name"`
    Id          int    `json:"id"`
    Mid         string `json:"mid"`
    AccountId   int    `json:"account_id"`
    Forward     int    `json:"forward"`
    Comment     int    `json:"comment"`
    Like        int    `json:"like"`
    Pubtime     string `json:"pubtime"`
    CrawlTime   string `json:"crawl_time"`
    Url         string `json:"url"`
}

type GetWeiboMsgListResponseData struct {
    Total       int64      `json:"total"`
    TotalPage   int        `json:"total_page"`
    CurrentPage int        `json:"current_page"`
    MsgList     []WeiboMsg `json:"msg_list"`
}

type GetWeiboMsgListResponse struct {
    Response
    Data GetWeiboMsgListResponseData
}
