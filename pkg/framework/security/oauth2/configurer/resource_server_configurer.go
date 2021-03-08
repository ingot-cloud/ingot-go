package configurer

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer/oauth"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
	anonymous "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/anoymous"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/authresult"
)

// ResourceServerConfigurerAdapter 资源服务器安全配置
type ResourceServerConfigurerAdapter struct {
	*config.WebSecurityConfigurerAdapter

	tokenExtractor        authentication.TokenExtractor
	authenticationManager coreAuth.Manager
}

// NewResourceServerConfigurer 实例化
func NewResourceServerConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.Manager) security.ResourceServerConfigurer {
	instance := &ResourceServerConfigurerAdapter{
		tokenExtractor:        tokenExtractor,
		authenticationManager: authenticationManager,
	}

	instance.WebSecurityConfigurerAdapter = config.NewWebSecurityConfigurerAdapter(instance)
	return instance
}

// HTTPConfigure 配置
func (a *ResourceServerConfigurerAdapter) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	http.RequestMatcher(a.RequestMatcher)
	http.Apply(oauth.NewSecurityConfigurer(a.tokenExtractor, a.authenticationManager))
	http.Apply(anonymous.NewSecurityConfigurer())
	http.Apply(authresult.NewSecurityConfigurer())
	return nil
}

// RequestMatcher 请求匹配器
func (a *ResourceServerConfigurerAdapter) RequestMatcher(ctx *ingot.Context) bool {
	current := ctx.Request.RequestURI
	// 当前url不能匹配token授权url
	for _, p := range endpoint.Paths {
		if p == current {
			return false
		}
	}
	return true
}
