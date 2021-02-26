package provider

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// AuthorizationServerTokenServices 授权服务器 token 服务
func AuthorizationServerTokenServices(oauth2Container *container.OAuth2Container, securityContainer *container.SecurityContainer, enhancer token.Enhancer, manager authentication.Manager) token.AuthorizationServerTokenServices {
	tokenServices := oauth2Container.DefaultTokenServices
	client := securityContainer.ClientDetailsService
	if _, ok := client.(*clientdetails.NilClientdetails); !ok {
		tokenServices.ClientDetailsService = client
	}
	tokenServices.TokenEnhancer = enhancer
	tokenServices.AuthenticationManager = manager
	return tokenServices
}

// ConsumerTokenServices 令牌撤销
func ConsumerTokenServices(oauth2Container *container.OAuth2Container) token.ConsumerTokenServices {
	return oauth2Container.DefaultTokenServices
}

// TokenEnhancer token增强，默认使用增强链
func TokenEnhancer(enhancers token.Enhancers, jwtEnhancer *store.JwtAccessTokenConverter) token.Enhancer {
	chain := &token.EnhancerChain{}
	// 默认追加 jwt enhancer
	enhancers = append(enhancers, jwtEnhancer)
	chain.SetTokenEnhancers(enhancers)
	return chain
}

// AuthorizationAuthenticationManager 授权服务器中的认证管理器
func AuthorizationAuthenticationManager(securityContainer *container.SecurityContainer) authentication.Manager {
	return authentication.NewProviderManager(securityContainer.Providers)
}
