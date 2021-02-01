package provider

import (
	"github.com/google/wire"
	innerSecurity "github.com/ingot-cloud/ingot-go/internal/app/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// NewWebSecurityConfigurers 注入 web 安全配置
func NewWebSecurityConfigurers() (security.WebSecurityConfigurers, error) {
	config := &innerSecurity.CustomHTTPConfigurer{}
	return config.Build()
}

// Security 安全注入
var Security = wire.NewSet(
	NewWebSecurityConfigurers,
	config.BuildWebSecurityFilter,
)
