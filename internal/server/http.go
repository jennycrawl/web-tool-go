package server

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    apiV1 "web-tool-go/api/v1"
    "web-tool-go/docs"
    "web-tool-go/internal/handler"
    "web-tool-go/internal/middleware"
    "web-tool-go/pkg/jwt"
    "web-tool-go/pkg/log"
    "web-tool-go/pkg/server/http"
)

func NewHTTPServer(
    logger *log.Logger,
    conf *viper.Viper,
    jwt *jwt.JWT,
    userHandler *handler.UserHandler,
    weiboMsgHandler *handler.WeiboMsgHandler,
    weiboAccountHandler *handler.WeiboAccountHandler,
    weiboStatistics *handler.WeiboStatisticsHandler,
) *http.Server {
    gin.SetMode(gin.DebugMode)
    s := http.NewServer(
        gin.Default(),
        logger,
        http.WithServerHost(conf.GetString("http.host")),
        http.WithServerPort(conf.GetInt("http.port")),
    )

    // swagger doc
    docs.SwaggerInfo.BasePath = "/v1"
    s.GET("/swagger/*any", ginSwagger.WrapHandler(
        swaggerfiles.Handler,
        //ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
        ginSwagger.DefaultModelsExpandDepth(-1),
        ginSwagger.PersistAuthorization(true),
    ))

    s.Use(
        middleware.CORSMiddleware(),
        middleware.ResponseLogMiddleware(logger),
        middleware.RequestLogMiddleware(logger),
        //middleware.SignMiddleware(log),
    )
    s.GET("/", func(ctx *gin.Context) {
        logger.WithContext(ctx).Info("hello")
        apiV1.HandleSuccess(ctx, map[string]interface{}{
            ":)": "Thank you for using nunu!",
        })
    })

    //v1 := s.Group("/v1")
    //{
    //    // No route group has permission
    //    noAuthRouter := v1.Group("/")
    //    {
    //        noAuthRouter.POST("/register", userHandler.Register)
    //        noAuthRouter.POST("/login", userHandler.Login)
    //    }
    //    // Non-strict permission routing group
    //    noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
    //    {
    //        noStrictAuthRouter.GET("/user", userHandler.GetProfile)
    //    }
    //
    //    // Strict permission routing group
    //    strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
    //    {
    //        strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
    //    }
    //}

    v1 := s.Group("/v1")
    {
        weibo := v1.Group("/weibo")
        {
            weiboAccount := weibo.Group("/account")
            {
                weiboAccount.GET("/list", weiboAccountHandler.GetWeiboAccountList)
                weiboAccount.POST("/", weiboAccountHandler.CreateWeiboAccount)
                weiboAccount.PATCH("/:id", weiboAccountHandler.UpdateWeiboAccount)
                weiboAccount.DELETE("/:id", weiboAccountHandler.DeleteWeiboAccount)
            }

            weibo.GET("/msg", weiboMsgHandler.GetWeiboMsgList)
            weibo.GET("/statistics", weiboStatistics.GetWeiboStatisticsList)
        }
    }

    return s
}
