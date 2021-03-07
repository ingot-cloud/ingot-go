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

	authenticationManager authentication.Manager
}

// NewAuthorizationServerWebSecurityConfigurer 实例化
func NewAuthorizationServerWebSecurityConfigurer(authenticationManager coreAuth.Manager) security.AuthorizationServerWebSecurityConfigurer {
	pre := &AuthorizationWebSecurityConfigurer{
		authenticationManager: authenticationManager,
	}
	return &AuthorizationWebSecurityConfigurer{
		WebSecurityConfigurerAdapter: config.NewWebSecurityConfigurerAdapter(pre),
	}
}

// HTTPConfigure 配置
func (c *AuthorizationWebSecurityConfigurer) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	http.RequestMatcher(c.requestMatcher)
	http.Apply(basic.NewSecurityConfigurer(c.authenticationManager))
	http.Apply(anonymous.NewSecurityConfigurer())
	return nil
}

func (c *AuthorizationWebSecurityConfigurer) requestMatcher(ctx *ingot.Context) bool {
	current := ctx.Request.RequestURI
	for _, p := range endpoint.Paths {
		if p == current {
			return true
		}
	}
	return false
}
