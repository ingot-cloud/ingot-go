package factory

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
)

// HTTPConfigSet 单独注入 http config
func HTTPConfigSet(config *config.Config) (server.Config, error) {
	return config.Server, nil
}

// AuthConfigSet 单独注入 auth config
func AuthConfigSet(config *config.Config) (config.Auth, error) {
	return config.Auth, nil
}

// Config 需要单独注入的配置
var Config = wire.NewSet(
	HTTPConfigSet,
	AuthConfigSet,
)
