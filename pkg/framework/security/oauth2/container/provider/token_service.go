package provider

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// DefaultTokenServices 默认的服务
func DefaultTokenServices(config config.OAuth2, tokenStore token.Store) *token.DefaultTokenServices {
	service := token.NewDefaultTokenServices(tokenStore)
	service.ReuseRefreshToken = config.AuthorizationServer.ReuseRefreshToken
	service.SupportRefreshToken = config.AuthorizationServer.SupportRefreshToken
	return service
}

// ResourceServerTokenServices 资源服务器 token 服务
func ResourceServerTokenServices(tokenServices *token.DefaultTokenServices) token.ResourceServerTokenServices {
	return tokenServices
}
