package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// NewWebSecurityConfigurers 注入 web 安全配置
func NewWebSecurityConfigurers() (config.WebSecurityConfigurers, error) {
	config := &security.Configurer{}
	return config.Build()
}

// Security 安全注入
var Security = wire.NewSet(
	NewWebSecurityConfigurers,
	config.BuildWebSecurityFilter,
)
