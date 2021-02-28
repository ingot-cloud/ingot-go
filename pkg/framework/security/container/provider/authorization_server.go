package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// AuthorizationServerContainerSet 授权服务器容器
var AuthorizationServerContainerSet = wire.NewSet(wire.Struct(new(container.AuthorizationServerContainer), "*"))

// AuthorizationServerTokenServices 授权服务器 token 服务
func AuthorizationServerTokenServices(oauth2Container *container.OAuth2Container, securityContainer *container.SecurityContainer, enhancer token.Enhancer, manager authentication.Manager, injector container.SecurityInjector) token.AuthorizationServerTokenServices {
	if injector.GetAuthorizationServerTokenServices() != nil {
		return injector.GetAuthorizationServerTokenServices()
	}
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
func ConsumerTokenServices(oauth2Container *container.OAuth2Container, injector container.SecurityInjector) token.ConsumerTokenServices {
	if injector.GetConsumerTokenServices() != nil {
		return injector.GetConsumerTokenServices()
	}
	return oauth2Container.DefaultTokenServices
}

// TokenEnhancer token增强，默认使用增强链
func TokenEnhancer(enhancers token.Enhancers, oauth2Container *container.OAuth2Container, injector container.SecurityInjector) token.Enhancer {
	if injector.GetTokenEnhancer() != nil {
		return injector.GetTokenEnhancer()
	}
	chain := &token.EnhancerChain{}
	// 默认追加 jwt enhancer
	enhancers = append(enhancers, oauth2Container.JwtAccessTokenConverter)
	chain.SetTokenEnhancers(enhancers)
	return chain
}

// AuthorizationAuthenticationManager 授权服务器中的认证管理器
func AuthorizationAuthenticationManager(securityContainer *container.SecurityContainer, injector container.SecurityInjector) authentication.Manager {
	if injector.GetAuthorizationAuthenticationManager() != nil {
		return injector.GetAuthorizationAuthenticationManager()
	}
	return authentication.NewProviderManager(securityContainer.Providers)
}
