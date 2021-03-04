package configurer

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
	anonymous "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/anoymous"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/basic"
)

// AuthorizationWebSecurityConfigurer 授权服务器配置
type AuthorizationWebSecurityConfigurer struct {
	authenticationManager authentication.Manager
}

// NewAuthorizationServerWebSecurityConfigurer 实例化
func NewAuthorizationServerWebSecurityConfigurer(authenticationManager coreAuth.Manager) security.AuthorizationServerWebSecurityConfigurer {
	configurer := &AuthorizationWebSecurityConfigurer{
		authenticationManager: authenticationManager,
	}
	return config.NewWebSecurityConfigurerAdapter(nil, configurer)
}

// Configure 配置
func (oa *AuthorizationWebSecurityConfigurer) Configure(http security.HTTPSecurityBuilder) error {
	http.Apply(basic.NewSecurityConfigurer(oa.authenticationManager))
	http.Apply(anonymous.NewSecurityConfigurer())
	return nil
}
