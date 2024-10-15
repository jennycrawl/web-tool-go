package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/service"
)

type WeiboStatisticsHandler struct {
    *Handler
    weiboStatisticsService service.WeiboStatisticsService
}

func NewWeiboStatisticsHandler(
    handler *Handler,
    weiboStatisticsService service.WeiboStatisticsService,
) *WeiboStatisticsHandler {
    return &WeiboStatisticsHandler{
        Handler:                handler,
        weiboStatisticsService: weiboStatisticsService,
    }
}

//GetWeiboStatisticsList godoc
//@Summary 获取微博统计信息
//@Schemes
//@Description
//@Tags 微博模块
//@Accept json
//@Produce json
//@Security Bearer
//@Success 200 {object} v1.GetWeiboStatisticsListResponseData
//@Router /weibo/statistics [get]
func (h *WeiboStatisticsHandler) GetWeiboStatisticsList(ctx *gin.Context) {
    accountId, _ := strconv.Atoi(ctx.DefaultQuery("account_id", "0"))
    startTime := dateStrToTime(ctx.DefaultQuery("start_date", ""))

    endTime := dateStrToTime(ctx.DefaultQuery("end_date", ""))
    if !endTime.IsZero() {
        endTime = endTime.AddDate(0, 0, 1)
    }

    statisticsList, err := h.weiboStatisticsService.GetWeiboStatisticsList(ctx, accountId, startTime, endTime)
    if err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, []v1.WeiboStatistics{})
        return
    }
    if *statisticsList == nil {
        v1.HandleSuccess(ctx, []v1.WeiboStatistics{})
    } else {
        v1.HandleSuccess(ctx, statisticsList)
    }
}
