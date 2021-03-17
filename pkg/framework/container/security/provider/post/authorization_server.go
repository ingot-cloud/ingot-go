package post

// import (
// 	"github.com/google/wire"
// 	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
// 	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
// 	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
// 	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/pre"
// 	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
// 	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/endpoint"
// 	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
// 	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/granter"
// )

// func enableAuthorizationServer(sc container.SecurityContainerCombine) bool {
// 	oauth2Container := sc.GetOAuth2Container()
// 	return enableAuthorizationServerWithConfig(oauth2Container.OAuth2Config)
// }

// func enableAuthorizationServerWithConfig(oc config.OAuth2) bool {
// 	return oc.AuthorizationServer.Enable
// }

// // AuthorizationServerContainer 授权服务器容器
// var AuthorizationServerContainer = wire.NewSet(
// 	wire.Struct(new(container.AuthorizationServerContainer), "*"),

// 	// Fields
// 	AuthorizationAuthenticationManager,
// 	AuthorizationServerConfigurer,
// 	AuthorizationServerTokenServices,
// 	ConsumerTokenServices,
// 	TokenEndpoint,
// 	TokenEndpointHTTPConfigurer,
// 	TokenEnhancer,
// 	TokenGranter,
// 	PasswordTokenGranter,
// )

// // AuthorizationAuthenticationManager 授权服务器中的认证管理器
// func AuthorizationAuthenticationManager(pc *container.AuthProvidersContainer, sc container.SecurityContainerCombine) authentication.AuthorizationManager {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.AuthorizationAuthenticationManager(pc)
// }

// // AuthorizationServerConfigurer 授权服务器配置
// func AuthorizationServerConfigurer(manager authentication.AuthorizationManager, sc container.SecurityContainerCombine) security.AuthorizationServerConfigurer {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.AuthorizationServerConfigurer(manager)
// }

// // AuthorizationServerTokenServices 授权服务器 token 服务
// func AuthorizationServerTokenServices(config config.OAuth2, tokenStore token.Store, common *container.CommonContainer, enhancer token.Enhancer, manager authentication.AuthorizationManager, sc container.SecurityContainerCombine) token.AuthorizationServerTokenServices {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.AuthorizationServerTokenServices(config, tokenStore, common, enhancer, manager)
// }

// // ConsumerTokenServices 令牌撤销
// func ConsumerTokenServices(tokenStore token.Store, sc container.SecurityContainerCombine) token.ConsumerTokenServices {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.ConsumerTokenServices(tokenStore)
// }

// // TokenEndpoint 端点
// func TokenEndpoint(granter token.Granter, common *container.CommonContainer, sc container.SecurityContainerCombine) *endpoint.TokenEndpoint {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.TokenEndpoint(granter, common)
// }

// // TokenEndpointHTTPConfigurer 端点配置
// func TokenEndpointHTTPConfigurer(tokenEndpoint *endpoint.TokenEndpoint, sc container.SecurityContainerCombine) endpoint.OAuth2HTTPConfigurer {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.TokenEndpointHTTPConfigurer(tokenEndpoint)
// }

// // TokenEnhancer token增强，默认使用增强链
// func TokenEnhancer(oauth2Container *container.OAuth2Container, sc container.SecurityContainerCombine) token.Enhancer {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.TokenEnhancer(oauth2Container)
// }

// // TokenGranter token 授权
// func TokenGranter(password *granter.PasswordTokenGranter, sc container.SecurityContainerCombine) token.Granter {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.TokenGranter(password)
// }

// // PasswordTokenGranter 密码模式授权
// func PasswordTokenGranter(tokenServices token.AuthorizationServerTokenServices, manager authentication.AuthorizationManager, sc container.SecurityContainerCombine) *granter.PasswordTokenGranter {
// 	if !enableAuthorizationServer(sc) {
// 		return nil
// 	}
// 	return pre.PasswordTokenGranter(tokenServices, manager)
// }
