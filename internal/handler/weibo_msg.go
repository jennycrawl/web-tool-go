package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/service"
)

type WeiboMsgHandler struct {
    *Handler
    weiboMsgService service.WeiboMsgService
}

func NewWeiboMsgHandler(
    handler *Handler,
    weiboMsgService service.WeiboMsgService,
) *WeiboMsgHandler {
    return &WeiboMsgHandler{
        Handler:         handler,
        weiboMsgService: weiboMsgService,
    }
}

//GetWeiboMsgList godoc
//@Summary 获取微博文章列表
//@Schemes
//@Description
//@Tags 微博模块
//@Accept json
//@Produce json
//@Security Bearer
//@Success 200 {object} v1.GetWeiboMsgListResponseData
//@Router /weibo/msg [get]
func (h *WeiboMsgHandler) GetWeiboMsgList(ctx *gin.Context) {
    accountId, _ := strconv.Atoi(ctx.DefaultQuery("account_id", "0"))
    startTime := dateStrToTime(ctx.DefaultQuery("start_date", ""))

    endTime := dateStrToTime(ctx.DefaultQuery("end_date", ""))
    if !endTime.IsZero() {
        endTime = endTime.AddDate(0, 0, 1)
    }

    sortField := ctx.DefaultQuery("sort_field", "id")

    sortOrder := ctx.DefaultQuery("sort_order", "asc")
    if sortOrder != "asc" && sortOrder != "desc" {
        sortOrder = "asc"
    }

    page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    if page <= 0 {
        page = 1
    }

    perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "20"))
    if perPage <= 0 {
        perPage = 20
    }

    h.logger.Debug(fmt.Sprintf("query params, accountId:%d, startTime:%s, endTime:%s, sortField:%s, sortOrder:%s, page:%d, perPage:%d", accountId, startTime, endTime, sortField, sortOrder, page, perPage))

    msgList, err := h.weiboMsgService.GetWeiboMsgList(ctx, accountId, startTime, endTime, sortField, sortOrder, page, perPage)
    if err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    v1.HandleSuccess(ctx, msgList)
}
