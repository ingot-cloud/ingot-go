package configurer

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
	anonymous "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/anoymous"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/authresult"
)

// ResourceWebSecurityConfigurer 资源服务器安全配置
type ResourceWebSecurityConfigurer struct {
	*config.WebSecurityConfigurerAdapter

	tokenExtractor        authentication.TokenExtractor
	authenticationManager coreAuth.Manager
}

// NewResourceServerWebSecurityConfigurer 实例化
func NewResourceServerWebSecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.Manager) security.ResourceServerWebSecurityConfigurer {
	pre := &ResourceWebSecurityConfigurer{
		tokenExtractor:        tokenExtractor,
		authenticationManager: authenticationManager,
	}

	return &ResourceWebSecurityConfigurer{
		WebSecurityConfigurerAdapter: config.NewWebSecurityConfigurerAdapter(pre),
	}
}

// todo 自定义过滤器，如何加入到资源服务器中

// HTTPConfigure 配置
func (c *ResourceWebSecurityConfigurer) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	http.RequestMatcher(c.requestMatcher)
	http.AddFilter(authentication.NewOAuth2ProcessingFilter(c.tokenExtractor, c.authenticationManager))
	http.Apply(anonymous.NewSecurityConfigurer())
	http.Apply(authresult.NewSecurityConfigurer())
	return nil
}

func (c *ResourceWebSecurityConfigurer) requestMatcher(ctx *ingot.Context) bool {
	current := ctx.Request.RequestURI
	// 当前url不能匹配token授权url
	for _, p := range endpoint.Paths {
		if p == current {
			return false
		}
	}
	return true
}
