package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "time"
    "web-tool-go/pkg/jwt"
    "web-tool-go/pkg/log"
)

type Handler struct {
    logger *log.Logger
}

func NewHandler(
    logger *log.Logger,
) *Handler {
    return &Handler{
        logger: logger,
    }
}
func GetUserIdFromCtx(ctx *gin.Context) string {
    v, exists := ctx.Get("claims")
    if !exists {
        return ""
    }
    return v.(*jwt.MyCustomClaims).UserId
}

func dateStrToTime(str string) time.Time {
    if str == "" {
        return time.Time{}
    }
    if len(str) > len("2006-01-02") {
        str = str[:len("2006-01-02")]
    }
    date, err := time.Parse("2006-01-02", str)
    if err != nil {
        fmt.Println("Error parsing date:", err)
        return time.Time{}
    }

    return date
}
