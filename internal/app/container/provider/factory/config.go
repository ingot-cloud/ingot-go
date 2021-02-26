package factory

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	httpConfig "github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
)

// HTTPConfigSet 单独注入 http config
func HTTPConfigSet(config *config.Config) (httpConfig.HTTPConfig, error) {
	return config.Server, nil
}

// SecurityConfigSet 单独注入 Security config
func SecurityConfigSet(config *config.Config) (config.Security, error) {
	return config.Security, nil
}

// Config 需要单独注入的配置
var Config = wire.NewSet(
	HTTPConfigSet,
	SecurityConfigSet,
)
