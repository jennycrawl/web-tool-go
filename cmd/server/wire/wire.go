//go:build wireinject
// +build wireinject

package wire

import (
    "github.com/google/wire"
    "github.com/spf13/viper"
    "web-tool-go/internal/handler"
    "web-tool-go/internal/repository"
    "web-tool-go/internal/server"
    "web-tool-go/internal/service"
    "web-tool-go/pkg/app"
    "web-tool-go/pkg/jwt"
    "web-tool-go/pkg/log"
    "web-tool-go/pkg/server/http"
    "web-tool-go/pkg/sid"
)

var repositorySet = wire.NewSet(
    repository.NewDB,
    //repository.NewRedis,
    repository.NewRepository,
    repository.NewTransaction,
    repository.NewUserRepository,
    repository.NewWeiboMsgRepository,
    repository.NewWeiboAccountRepository,
    repository.NewWeiboStatisticsRepository,
)

var serviceSet = wire.NewSet(
    service.NewService,
    service.NewUserService,
    service.NewWeiboMsgService,
    service.NewWeiboAccountService,
    service.NewWeiboStatisticsService,
)

var handlerSet = wire.NewSet(
    handler.NewHandler,
    handler.NewUserHandler,
    handler.NewWeiboMsgHandler,
    handler.NewWeiboAccountHandler,
    handler.NewWeiboStatisticsHandler,
)

var serverSet = wire.NewSet(
    server.NewHTTPServer,
    server.NewJob,
)

// build App
func newApp(
    httpServer *http.Server,
    job *server.Job,
// task *server.Task,
) *app.App {
    return app.NewApp(
        app.WithServer(httpServer, job),
        app.WithName("demo-server"),
    )
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
    panic(wire.Build(
        repositorySet,
        serviceSet,
        handlerSet,
        serverSet,
        sid.NewSid,
        jwt.NewJwt,
        newApp,
    ))
}
