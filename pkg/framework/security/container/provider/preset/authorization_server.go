package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// AuthorizationServerContainer 授权服务器容器
var AuthorizationServerContainer = wire.NewSet(wire.Struct(new(container.AuthorizationServerContainer), "*"))

// AuthorizationServerContainerFields 授权服务器容器所有字段
var AuthorizationServerContainerFields = wire.NewSet(
	AuthorizationServerTokenServices,
	ConsumerTokenServices,
	TokenEnhancer,
	TokenEnhancers,
	AuthorizationAuthenticationManager,
)

// AuthorizationServerTokenServices 授权服务器 token 服务
func AuthorizationServerTokenServices(oauth2Container *container.OAuth2Container, securityContainer *container.SecurityContainer, enhancer token.Enhancer, manager authentication.AuthorizationManager) token.AuthorizationServerTokenServices {
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
func TokenEnhancer(enhancers token.Enhancers, oauth2Container *container.OAuth2Container) token.Enhancer {
	chain := &token.EnhancerChain{}
	// 默认追加 jwt enhancer
	enhancers = append(enhancers, oauth2Container.JwtAccessTokenConverter)
	chain.SetTokenEnhancers(enhancers)
	return chain
}

// TokenEnhancers 自定义增强
func TokenEnhancers() token.Enhancers {
	return nil
}

// AuthorizationAuthenticationManager 授权服务器中的认证管理器
func AuthorizationAuthenticationManager(securityContainer *container.SecurityContainer) authentication.AuthorizationManager {
	return authentication.NewProviderManager(securityContainer.Providers)
}
