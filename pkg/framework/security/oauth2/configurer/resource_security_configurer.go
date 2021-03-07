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
}

// NewResourceServerWebSecurityConfigurer 实例化
func NewResourceServerWebSecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.Manager) security.ResourceServerWebSecurityConfigurer {
	pre := &resourceHTTPSecurityConfigurer{
		tokenExtractor:        tokenExtractor,
		authenticationManager: authenticationManager,
	}
	return &ResourceWebSecurityConfigurer{
		WebSecurityConfigurerAdapter: config.NewWebSecurityConfigurerAdapter(nil, pre),
	}
}

type resourceHTTPSecurityConfigurer struct {
	tokenExtractor        authentication.TokenExtractor
	authenticationManager coreAuth.Manager
}

// Configure 配置
func (oa *resourceHTTPSecurityConfigurer) Configure(http security.HTTPSecurityBuilder) error {
	http.RequestMatcher(oa.requestMatcher)
	http.AddFilter(authentication.NewOAuth2ProcessingFilter(oa.tokenExtractor, oa.authenticationManager))
	http.Apply(anonymous.NewSecurityConfigurer())
	http.Apply(authresult.NewSecurityConfigurer())
	return nil
}

func (oa *resourceHTTPSecurityConfigurer) requestMatcher(ctx *ingot.Context) bool {
	current := ctx.Request.RequestURI
	// 当前url不能匹配token授权url
	for _, p := range endpoint.Paths {
		if p == current {
			return false
		}
	}
	return true
}
