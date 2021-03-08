package oauth

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"

	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
)

// SecurityConfigurer basic 验证
type SecurityConfigurer struct {
	tokenExtractor        authentication.TokenExtractor
	authenticationManager coreAuth.Manager
}

// NewSecurityConfigurer 配置
func NewSecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.Manager) *SecurityConfigurer {
	return &SecurityConfigurer{
		tokenExtractor:        tokenExtractor,
		authenticationManager: authenticationManager,
	}
}

// HTTPConfigure 配置
func (c *SecurityConfigurer) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	http.AddFilter(authentication.NewOAuth2ProcessingFilter(c.tokenExtractor, c.authenticationManager))
	return nil
}
