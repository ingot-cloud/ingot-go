package app

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/container/injector"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot"
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// Run start app
func Run(ctx context.Context, options *config.Options) error {
	// 初始化模块
	factory := func(ctx context.Context) (bootContainer.Container, func(), error) {
		return initContainer(ctx, options)
	}

	return boot.Run(ctx, factory)
}

func initContainer(ctx context.Context, options *config.Options) (bootContainer.Container, func(), error) {

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

	containerCombine, appCleanup, err := injector.BuildContainerCombine(config, options)

	// boot 容器
	container, containerCleanupFunc, err := injector.BuildContainer(config, options, containerCombine)
	if err != nil {
		return nil, nil, err
	}

	return container, func() {
		containerCleanupFunc()
		appCleanup()
		loggerCleanFunc()
	}, nil
}
