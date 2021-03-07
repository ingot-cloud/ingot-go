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

// AuthorizationWebSecurityConfigurer 授权服务器配置
type AuthorizationWebSecurityConfigurer struct {
	*config.WebSecurityConfigurerAdapter
}

// NewAuthorizationServerWebSecurityConfigurer 实例化
func NewAuthorizationServerWebSecurityConfigurer(authenticationManager coreAuth.Manager) security.AuthorizationServerWebSecurityConfigurer {
	pre := &authorizationHTTPSecurityConfigurer{
		authenticationManager: authenticationManager,
	}
	return &AuthorizationWebSecurityConfigurer{
		WebSecurityConfigurerAdapter: config.NewWebSecurityConfigurerAdapter(nil, pre),
	}
}

type authorizationHTTPSecurityConfigurer struct {
	authenticationManager authentication.Manager
}

// Configure 配置
func (oa *authorizationHTTPSecurityConfigurer) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	http.RequestMatcher(oa.requestMatcher)
	http.Apply(basic.NewSecurityConfigurer(oa.authenticationManager))
	http.Apply(anonymous.NewSecurityConfigurer())
	return nil
}

func (oa *authorizationHTTPSecurityConfigurer) requestMatcher(ctx *ingot.Context) bool {
	current := ctx.Request.RequestURI
	for _, p := range endpoint.Paths {
		if p == current {
			return true
		}
	}
	return false
}
