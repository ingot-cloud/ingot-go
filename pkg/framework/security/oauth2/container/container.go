package container

import (
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// OAuth2Container 公共容器
type OAuth2Container struct {
	Config                  config.OAuth2
	TokenStore              token.Store
	JwtAccessTokenConverter *store.JwtAccessTokenConverter
	AccessTokenConverter    token.AccessTokenConverter
}

// ResourceServerContainer 资源服务器容器
type ResourceServerContainer struct {
	OAuth2SecurityConfigurer    *config.OAuth2SecurityConfigurer
	TokenExtractor              authentication.TokenExtractor
	AuthenticationManager       coreAuth.Manager
	ResourceServerTokenServices token.ResourceServerTokenServices
	UserAuthenticationConverter token.UserAuthenticationConverter
}

// AuthorizationServerContainer 授权服务器容器
type AuthorizationServerContainer struct {
	TokenEnhancer               token.Enhancer
	ClientDetailsService        clientdetails.Service
	AuthenticationManager       coreAuth.Manager
	UserAuthenticationConverter token.UserAuthenticationConverter
}

// todo AuthenticationManager 资源服务器和授权服务器使用的对象应该不同
// todo 需要确定资源服务器中的 AuthenticationManager 是否需要 ClientDetailsService
