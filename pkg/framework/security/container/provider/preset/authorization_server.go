package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer"
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
	AuthorizationServerWebSecurityConfigurer,
	AuthorizationServerTokenServices,
	ConsumerTokenServices,
	TokenEndpoint,
	TokenEndpointHTTPConfigurer,
	TokenEnhancer,
	TokenGranter,
	PasswordTokenGranter,
)

// AuthorizationAuthenticationManager 授权服务器中的认证管理器
func AuthorizationAuthenticationManager(providerContainer *container.AuthProvidersContainer) authentication.AuthorizationManager {
	return authentication.NewProviderManager(providerContainer.Providers)
}

// AuthorizationServerWebSecurityConfigurer 授权服务器配置
func AuthorizationServerWebSecurityConfigurer(manager authentication.AuthorizationManager) security.AuthorizationServerWebSecurityConfigurer {
	return configurer.NewAuthorizationServerWebSecurityConfigurer(manager)
}

// AuthorizationServerTokenServices 授权服务器 token 服务
func AuthorizationServerTokenServices(oauth2Container *container.OAuth2Container, common *container.CommonContainer, enhancer token.Enhancer, manager authentication.AuthorizationManager) token.AuthorizationServerTokenServices {
	tokenServices := oauth2Container.DefaultTokenServices
	client := common.ClientDetailsService
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
func TokenEndpoint(granter token.Granter, common *container.CommonContainer) *endpoint.TokenEndpoint {
	return endpoint.NewTokenEndpoint(granter, common.ClientDetailsService)
}

// TokenEndpointHTTPConfigurer 端点配置
func TokenEndpointHTTPConfigurer(tokenEndpoint *endpoint.TokenEndpoint) endpoint.OAuth2HTTPConfigurer {
	return endpoint.NewOAuth2ApiConfig(tokenEndpoint)
}

// TokenEnhancer token增强，默认使用增强链
func TokenEnhancer(oauth2Container *container.OAuth2Container) token.Enhancer {
	chain := token.NewEnhancerChain()
	var enhancers []token.Enhancer
	// 默认追加 jwt enhancer
	enhancers = append(enhancers, oauth2Container.JwtAccessTokenConverter)
	chain.SetTokenEnhancers(enhancers)
	return chain
}

// TokenGranter token 授权
func TokenGranter(password *granter.PasswordTokenGranter) token.Granter {
	result := granter.NewCompositeTokenGranter()
	result.AddTokenGranter(password)
	return result
}

// PasswordTokenGranter 密码模式授权
func PasswordTokenGranter(tokenServices token.AuthorizationServerTokenServices, manager authentication.AuthorizationManager) *granter.PasswordTokenGranter {
	return granter.NewPasswordTokenGranter(tokenServices, manager)
}
