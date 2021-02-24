package provider

import (
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// OAuth2SecurityConfigurer 实例化 OAuth2 安全配置
func OAuth2SecurityConfigurer(tokenExtractor authentication.TokenExtractor, authenticationManager coreAuth.Manager) *config.OAuth2SecurityConfigurer {
	return config.NewOAuth2SecurityConfigurer(tokenExtractor, authenticationManager)
}

// TokenExtractor TokenExtrator接口默认实现
func TokenExtractor() authentication.TokenExtractor {
	return authentication.NewBearerTokenExtractor()
}

// ResourceAuthenticationManager OAuth2 实现
func ResourceAuthenticationManager(config config.OAuth2, tokenService token.ResourceServerTokenServices, clientDetailsService clientdetails.Service) coreAuth.Manager {
	manager := authentication.NewOAuth2AuthenticationManager(tokenService, clientDetailsService)
	manager.ResourceID = config.ResourceServer.ResourceID
	return manager
}

// ResourceServerTokenServices 资源服务器 token 服务
func ResourceServerTokenServices(tokenServices *token.DefaultTokenServices) token.ResourceServerTokenServices {
	return tokenServices
}
