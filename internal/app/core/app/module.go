package app

import (
	"context"
	"net/http"
	"time"

	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/injector"
	"github.com/ingot-cloud/ingot-go/internal/app/core/server"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

func initModule(ctx context.Context, options *Options) (func(), error) {

	// 初始化 config
	configCleanFunc, err := config.LoadConfig(options.ConfigFile)
	if err != nil {
		return nil, err
	}

	// 赋值 casbin 模型路径
	config.CONFIG.Casbin.ModelPath = options.CasbinModelFile

	loggerCleanFunc, err := log.InitLogger(config.CONFIG.Log)
	if err != nil {
		return nil, err
	}

	container, containerCleanupFunc, err := injector.BuildContainer()
	if err != nil {
		return nil, err
	}

	httpServerCleanupFunc := runHTTPServer(ctx, config.CONFIG.Server, container.Engine)

	return func() {
		httpServerCleanupFunc()
		containerCleanupFunc()
		loggerCleanFunc()
		configCleanFunc()
	}, nil
}

func runHTTPServer(ctx context.Context, conf config.Server, handler http.Handler) func() {
	httpServer := server.HTTPServer(handler, conf)

	go func() {
		log.WithContext(ctx).Infof("=== HTTP server started successfully, address=%s ===", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithContext(ctx).Fatalf("listen: %s\n", err)
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(5*time.Second))
		defer cancel()

		httpServer.SetKeepAlivesEnabled(false)
		if err := httpServer.Shutdown(ctx); err != nil {
			log.WithContext(ctx).Errorf(err.Error())
		}
	}
}
