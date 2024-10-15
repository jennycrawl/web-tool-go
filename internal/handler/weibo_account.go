package handler

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "net/http"
    "strconv"
    v1 "web-tool-go/api/v1"
    "web-tool-go/internal/service"
)

type WeiboAccountHandler struct {
    *Handler
    weiboAccountService service.WeiboAccountService
}

func NewWeiboAccountHandler(
    handler *Handler,
    weiboAccountService service.WeiboAccountService,
) *WeiboAccountHandler {
    return &WeiboAccountHandler{
        Handler:             handler,
        weiboAccountService: weiboAccountService,
    }
}

//GetWeiboAccountList godoc
//@Summary 获取微博号列表
//@Schemes
//@Description
//@Tags 微博模块
//@Accept json
//@Produce json
//@Security Bearer
//@Success 200 {object} v1.GetWeiboAccountListResponseData
//@Router /weibo/account/list [get]
func (h *WeiboAccountHandler) GetWeiboAccountList(ctx *gin.Context) {
    accountList, err := h.weiboAccountService.GetWeiboAccountList(ctx)
    if err != nil {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, []v1.WeiboAccount{})
        return
    }
    if *accountList == nil {
        v1.HandleSuccess(ctx, []v1.WeiboAccount{})
    } else {
        v1.HandleSuccess(ctx, accountList)
    }
}

// CreateWeiboAccount godoc
// @Summary 创建微博号
// @Schemes
// @Description
// @Tags 微博模块
// @Accept json
// @Produce json
// @Param request body v1.CreateWeiboAccountRequest true "params"
// @Success 200 {object} v1.Response
// @Router /weibo/account [post]
func (h *WeiboAccountHandler) CreateWeiboAccount(ctx *gin.Context) {
    req := new(v1.CreateWeiboAccountRequest)
    if err := ctx.ShouldBindJSON(req); err != nil {
        h.logger.WithContext(ctx).Error("weiboAccountService.Create error", zap.Error(err))
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    if err := h.weiboAccountService.CreateWeiboAccount(ctx, req); err != nil {
        h.logger.WithContext(ctx).Error("weiboAccountService.Create error", zap.Error(err))
        v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
        return
    }

    v1.HandleSuccess(ctx, nil)
}

// UpdateWeiboAccount godoc
// @Summary 修改微博号
// @Schemes
// @Description
// @Tags 微博模块
// @Accept json
// @Produce json
// @Param request body v1.UpdateWeiboAccountRequest true "params"
// @Success 200 {object} v1.Response
// @Router /weibo/account/{id} [patch]
func (h *WeiboAccountHandler) UpdateWeiboAccount(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, idErr := strconv.Atoi(idStr)
    if idErr != nil || id <= 0 {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    req := new(v1.UpdateWeiboAccountRequest)
    if err := ctx.ShouldBindJSON(req); err != nil {
        h.logger.WithContext(ctx).Error("weiboAccountService.Update error", zap.Error(err))
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    if err := h.weiboAccountService.UpdateWeiboAccount(ctx, id, req); err != nil {
        h.logger.WithContext(ctx).Error("weiboAccountService.Create error", zap.Error(err))
        v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
        return
    }

    v1.HandleSuccess(ctx, nil)
}

// DeleteWeiboAccount godoc
// @Summary 删除微博号
// @Schemes
// @Description
// @Tags 微博模块
// @Accept json
// @Produce json
// @Success 200 {object} v1.Response
// @Router /weibo/account/{id} [delete]
func (h *WeiboAccountHandler) DeleteWeiboAccount(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, idErr := strconv.Atoi(idStr)
    if idErr != nil || id <= 0 {
        v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
        return
    }

    if err := h.weiboAccountService.DeleteWeiboAccountById(ctx, id); err != nil {
        h.logger.WithContext(ctx).Error("weiboAccountService.Delete error", zap.Error(err))
        v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
        return
    }

    v1.HandleSuccess(ctx, nil)
}
