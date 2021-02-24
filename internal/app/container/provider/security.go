package provider

import (
	"github.com/google/wire"
	innerSecurity "github.com/ingot-cloud/ingot-go/internal/app/core/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
)

// NewWebSecurityConfigurers 注入 web 安全配置
func NewWebSecurityConfigurers() (security.WebSecurityConfigurers, error) {
	return innerSecurity.Build()
}

// Security 安全注入
var Security = wire.NewSet(
	NewWebSecurityConfigurers,
)
