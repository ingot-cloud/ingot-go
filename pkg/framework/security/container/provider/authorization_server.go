package provider

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// TokenEnhancer token增强，默认使用增强链
func TokenEnhancer(enhancers token.Enhancers, jwtEnhancer *store.JwtAccessTokenConverter) token.Enhancer {
	chain := &token.EnhancerChain{}
	// 默认追加 jwt enhancer
	enhancers = append(enhancers, jwtEnhancer)
	chain.SetTokenEnhancers(enhancers)
	return chain
}

// AuthorizationServerTokenServices 授权服务器 token 服务
func AuthorizationServerTokenServices(tokenServices *token.DefaultTokenServices, client clientdetails.Service, enhancer token.Enhancer, manager authentication.Manager) token.AuthorizationServerTokenServices {
	if _, ok := client.(*clientdetails.NilClientdetails); !ok {
		tokenServices.ClientDetailsService = client
	}
	tokenServices.TokenEnhancer = enhancer
	tokenServices.AuthenticationManager = manager
	return tokenServices
}

// ConsumerTokenServices 令牌撤销
func ConsumerTokenServices(tokenServices *token.DefaultTokenServices) token.ConsumerTokenServices {
	return tokenServices
}

// AuthorizationAuthenticationManager 授权服务器中的认证管理器
func AuthorizationAuthenticationManager(providers authentication.Providers) authentication.Manager {
	return authentication.NewProviderManager(providers)
}

// todo 提供 Providers
