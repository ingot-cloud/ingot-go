package app

import (
	"context"
	"github.com/ingot-cloud/ingot-go/internal/app/common/log"
	"os"
	"os/signal"
)

// Options 配置
type Options struct {
	ConfigFile      string
	CasbinModelFile string
}

// Run start app
func Run(context context.Context, options *Options) error {
	banner()

	// 初始化模块
	cleanupFunc, err := initModule(context, options)
	if err != nil {
		return err
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.WithContext(context).Info("Server exiting")

	cleanupFunc()
	return nil
}
