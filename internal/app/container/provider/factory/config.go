package factory

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	httpConfig "github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	oauth2Config "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
)

// HTTPConfig 单独注入 http config
func HTTPConfig(config *config.Config) (httpConfig.HTTPConfig, error) {
	return config.Server, nil
}

// SecurityConfig 单独注入 Security config
func SecurityConfig(config *config.Config) (config.Security, error) {
	return config.Security, nil
}

// OAuth2Config 单独注入 OAuth2 config
func OAuth2Config(config *config.Config) (oauth2Config.OAuth2, error) {
	return config.Security.OAuth2, nil
}

// Config 需要单独注入的配置
var Config = wire.NewSet(
	HTTPConfig,
	SecurityConfig,
	OAuth2Config,
)
