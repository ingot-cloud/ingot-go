package basic

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
)

// SecurityConfigurer basic 验证
type SecurityConfigurer struct {
	AuthenticationManager authentication.Manager
}

// NewSecurityConfigurer 配置
func NewSecurityConfigurer(manager authentication.Manager) *SecurityConfigurer {
	return &SecurityConfigurer{
		AuthenticationManager: manager,
	}
}

// Configure 配置
func (b *SecurityConfigurer) Configure(http security.HTTPSecurityBuilder) error {
	http.AddFilter(NewFilter(b.AuthenticationManager))
	return nil
}
