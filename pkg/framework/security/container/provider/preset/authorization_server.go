package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/granter"
)

// AuthorizationServerContainer 授权服务器容器
var AuthorizationServerContainer = wire.NewSet(wire.Struct(new(container.AuthorizationServerContainer), "*"))

// AuthorizationServerContainerFields 授权服务器容器所有字段
var AuthorizationServerContainerFields = wire.NewSet(
	AuthorizationAuthenticationManager,
	AuthorizationServerTokenServices,
	ConsumerTokenServices,
	TokenEndpoint,
	TokenEndpointHTTPConfigurer,
	TokenEnhancer,
	TokenEnhancers,
	TokenGranters,
	TokenGranter,
	PasswordTokenGranter,
)

// AuthorizationAuthenticationManager 授权服务器中的认证管理器
func AuthorizationAuthenticationManager(providerContainer *container.AuthProvidersContainer) authentication.AuthorizationManager {
	return authentication.NewProviderManager(providerContainer.Providers)
}

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

// TokenEndpoint 端点
func TokenEndpoint(granter token.Granter, securityContainer *container.SecurityContainer) *endpoint.TokenEndpoint {
	return endpoint.NewTokenEndpoint(granter, securityContainer.ClientDetailsService)
}

// TokenEndpointHTTPConfigurer 端点配置
func TokenEndpointHTTPConfigurer(tokenEndpoint *endpoint.TokenEndpoint) endpoint.OAuth2HTTPConfigurer {
	return endpoint.NewOAuth2ApiConfig(tokenEndpoint)
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

// TokenGranters 自定义授权
func TokenGranters() token.Granters {
	return nil
}

// TokenGranter token 授权
func TokenGranter(granters token.Granters, password *granter.PasswordTokenGranter) token.Granter {
	result := granter.NewCompositeTokenGranter()
	for _, g := range granters {
		result.AddTokenGranter(g)
	}

	result.AddTokenGranter(password)
	return result
}

// PasswordTokenGranter 密码模式授权
func PasswordTokenGranter(tokenServices token.AuthorizationServerTokenServices, manager authentication.AuthorizationManager) *granter.PasswordTokenGranter {
	return granter.NewPasswordTokenGranter(tokenServices, manager)
}
