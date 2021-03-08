package configurer

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
	anonymous "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/anoymous"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/basic"
)

// AuthorizationServerConfigurerAdapter 授权服务器配置
type AuthorizationServerConfigurerAdapter struct {
	*config.WebSecurityConfigurerAdapter

	authenticationManager authentication.Manager
}

// NewAuthorizationServerConfigurer 实例化
func NewAuthorizationServerConfigurer(authenticationManager coreAuth.Manager) security.AuthorizationServerConfigurer {
	instance := &AuthorizationServerConfigurerAdapter{
		authenticationManager: authenticationManager,
	}
	instance.WebSecurityConfigurerAdapter = config.NewWebSecurityConfigurerAdapter(instance)
	return instance
}

// HTTPConfigure 配置
func (a *AuthorizationServerConfigurerAdapter) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	http.RequestMatcher(a.RequestMatcher)
	http.Apply(basic.NewSecurityConfigurer(a.authenticationManager))
	http.Apply(anonymous.NewSecurityConfigurer())
	return nil
}

// RequestMatcher 请求匹配器
func (a *AuthorizationServerConfigurerAdapter) RequestMatcher(ctx *ingot.Context) bool {
	current := ctx.Request.RequestURI
	for _, p := range endpoint.Paths {
		if p == current {
			return true
		}
	}
	return false
}
