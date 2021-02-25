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
	ResourceServerTokenServices token.ResourceServerTokenServices
	OAuth2SecurityConfigurer    *config.OAuth2SecurityConfigurer
	TokenExtractor              authentication.TokenExtractor
	AuthenticationManager       coreAuth.Manager
	UserAuthenticationConverter token.UserAuthenticationConverter
}

// AuthorizationServerContainer 授权服务器容器
type AuthorizationServerContainer struct {
	AuthorizationServerTokenServices token.AuthorizationServerTokenServices
	ConsumerTokenServices            token.ConsumerTokenServices
	TokenEnhancer                    token.Enhancer
	ClientDetailsService             clientdetails.Service
	AuthenticationManager            coreAuth.Manager
	UserAuthenticationConverter      token.UserAuthenticationConverter
}
