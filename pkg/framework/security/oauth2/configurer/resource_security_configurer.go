package configurer

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
	anonymous "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/configurers/anoymous"
)

// ResourceWebSecurityConfigurer 资源服务器安全配置
type ResourceWebSecurityConfigurer struct {
	tokenExtractor        authentication.TokenExtractor
	authenticationManager coreAuth.Manager
}

// NewResourceServerWebSecurityConfigurer 实例化
func NewResourceServerWebSecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.Manager) security.ResourceServerWebSecurityConfigurer {
	configurer := &ResourceWebSecurityConfigurer{
		tokenExtractor:        tokenExtractor,
		authenticationManager: authenticationManager,
	}
	return config.NewWebSecurityConfigurerAdapter(nil, configurer)
}

// Configure 配置
func (oa *ResourceWebSecurityConfigurer) Configure(http security.HTTPSecurityBuilder) error {
	http.AddFilter(authentication.NewOAuth2ProcessingFilter(oa.tokenExtractor, oa.authenticationManager))
	http.Apply(anonymous.NewSecurityConfigurer())
	return nil
}
