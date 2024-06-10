package app

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/container/injector"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// Run start app
func Run(ctx context.Context, options *config.Options) error {
	// 初始化模块
	factory := func(ctx context.Context) (container.Container, func(), error) {
		return initContainer(ctx, options)
	}

	return boot.Run(ctx, factory)
}

func initContainer(_ context.Context, options *config.Options) (container.Container, func(), error) {

	// 初始化 config
	config, err := injector.BuildConfiguration(options)
	if err != nil {
		return nil, nil, err
	}

	// 初始化 log
	loggerCleanFunc, err := log.InitLogger(config.Log)
	if err != nil {
		return nil, nil, err
	}

	container, containerCleanupFunc, err := injector.BuildContainer(config, options)
	if err != nil {
		return nil, nil, err
	}

	return container, func() {
		containerCleanupFunc()
		loggerCleanFunc()
	}, nil
}
