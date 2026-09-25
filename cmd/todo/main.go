package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/BraveTonni/gotodo/internal/core/logger"
	core_postgres_pool "github.com/BraveTonni/gotodo/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/BraveTonni/gotodo/internal/core/transport/http/middleware"
	core_http_server "github.com/BraveTonni/gotodo/internal/core/transport/http/server"
	users_transport_http "github.com/BraveTonni/gotodo/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("failed to init app logger: %w", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initialize postgres conn pool")
	pool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConfigMust())

	if err != nil {
		logger.Fatal("failed to init postgres conn pool: %w", zap.Error(err))
	}

	defer pool.Close()

	logger.Debug("Starting app")

	usersTransportHttp := users_transport_http.NewUsersHTTPHandler(nil)

	usersRoutes := usersTransportHttp.Routes()
	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)

	apiVersionRouter.RegisterRoutes(usersRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Metrics(),
	)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("http server run error", zap.Error(err))
	}

}
