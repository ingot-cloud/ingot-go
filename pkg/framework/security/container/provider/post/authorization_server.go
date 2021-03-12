package post

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/pre"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/granter"
)

func enableAuthorizationServer(sc container.SecurityContainerCombine) bool {
	oauth2Container := sc.GetOAuth2Container()
	return oauth2Container.OAuth2Config.AuthorizationServer.Enable
}

// AuthorizationServerContainer 授权服务器容器
var AuthorizationServerContainer = wire.NewSet(wire.Struct(new(container.AuthorizationServerContainer), "*"))

// AuthorizationServerContainerFields 授权服务器容器所有字段
var AuthorizationServerContainerFields = wire.NewSet(
	AuthorizationAuthenticationManager,
	AuthorizationServerConfigurer,
	AuthorizationServerTokenServices,
	ConsumerTokenServices,
	TokenEndpoint,
	TokenEndpointHTTPConfigurer,
	TokenEnhancer,
	TokenGranter,
	PasswordTokenGranter,
)

// AuthorizationAuthenticationManager 授权服务器中的认证管理器
func AuthorizationAuthenticationManager(sc container.SecurityContainerCombine) authentication.AuthorizationManager {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	return pre.AuthorizationAuthenticationManager(sc.GetAuthProvidersContainer())
}

// AuthorizationServerConfigurer 授权服务器配置
func AuthorizationServerConfigurer(sc container.SecurityContainerCombine) security.AuthorizationServerConfigurer {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	return pre.AuthorizationServerConfigurer(sc.GetAuthorizationServerContainer().AuthenticationManager)
}

// AuthorizationServerTokenServices 授权服务器 token 服务
func AuthorizationServerTokenServices(sc container.SecurityContainerCombine) token.AuthorizationServerTokenServices {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	config := sc.GetOAuth2Container().OAuth2Config
	tokenStore := sc.GetOAuth2Container().TokenStore
	common := sc.GetCommonContainer()
	enhancer := sc.GetAuthorizationServerContainer().TokenEnhancer
	manager := sc.GetAuthorizationServerContainer().AuthenticationManager
	return pre.AuthorizationServerTokenServices(config, tokenStore, common, enhancer, manager)
}

// ConsumerTokenServices 令牌撤销
func ConsumerTokenServices(sc container.SecurityContainerCombine) token.ConsumerTokenServices {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	return pre.ConsumerTokenServices(sc.GetOAuth2Container().TokenStore)
}

// TokenEndpoint 端点
func TokenEndpoint(sc container.SecurityContainerCombine) *endpoint.TokenEndpoint {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	return pre.TokenEndpoint(sc.GetAuthorizationServerContainer().TokenGranter, sc.GetCommonContainer())
}

// TokenEndpointHTTPConfigurer 端点配置
func TokenEndpointHTTPConfigurer(sc container.SecurityContainerCombine) endpoint.OAuth2HTTPConfigurer {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	return pre.TokenEndpointHTTPConfigurer(sc.GetAuthorizationServerContainer().TokenEndpoint)
}

// TokenEnhancer token增强，默认使用增强链
func TokenEnhancer(sc container.SecurityContainerCombine) token.Enhancer {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	return pre.TokenEnhancer(sc.GetOAuth2Container())
}

// TokenGranter token 授权
func TokenGranter(sc container.SecurityContainerCombine) token.Granter {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	return pre.TokenGranter(sc.GetAuthorizationServerContainer().PasswordTokenGranter)
}

// PasswordTokenGranter 密码模式授权
func PasswordTokenGranter(sc container.SecurityContainerCombine) *granter.PasswordTokenGranter {
	if !enableAuthorizationServer(sc) {
		return nil
	}
	authServerContainer := sc.GetAuthorizationServerContainer()
	return pre.PasswordTokenGranter(authServerContainer.AuthorizationServerTokenServices, authServerContainer.AuthenticationManager)
}
