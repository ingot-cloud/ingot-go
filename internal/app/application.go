package app

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/injector"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// Options 配置
type Options struct {
	ConfigFile      string
	CasbinModelFile string
}

// Run start app
func Run(context context.Context, options *Options) error {
	// 初始化模块
	container, cleanupFunc, err := initModule(context, options)
	if err != nil {
		return err
	}

	boot.Run(context, container, cleanupFunc)
	return nil
}

func initModule(ctx context.Context, options *Options) (*container.Container, func(), error) {

	// 初始化 config
	configCleanFunc, err := config.LoadConfig(options.ConfigFile)
	if err != nil {
		return nil, nil, err
	}

	// 赋值 casbin 模型路径
	config.CONFIG.Casbin.ModelPath = options.CasbinModelFile

	loggerCleanFunc, err := log.InitLogger(config.CONFIG.Log)
	if err != nil {
		return nil, nil, err
	}

	container, containerCleanupFunc, err := injector.BuildContainer(config.CONFIG.Server)
	if err != nil {
		return nil, nil, err
	}

	return container, func() {
		containerCleanupFunc()
		loggerCleanFunc()
		configCleanFunc()
	}, nil
}
