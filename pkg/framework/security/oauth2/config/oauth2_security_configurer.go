package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
)

// OAuth2SecurityConfigurer OAuth2安全配置
type OAuth2SecurityConfigurer struct {
	tokenExtractor        authentication.TokenExtractor
	authenticationManager coreAuth.Manager
}

// NewOAuth2SecurityConfigurer 实例化
func NewOAuth2SecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.Manager) *OAuth2SecurityConfigurer {
	return &OAuth2SecurityConfigurer{
		tokenExtractor:        tokenExtractor,
		authenticationManager: authenticationManager,
	}
}

// Configure 配置
func (oa *OAuth2SecurityConfigurer) Configure(http security.HTTPSecurityBuilder) error {
	http.AddFilter(authentication.NewOAuth2ProcessingFilter(oa.tokenExtractor, oa.authenticationManager))
	return nil
}
