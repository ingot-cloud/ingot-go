package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

// SecurityContainer 安全容器
type SecurityContainer struct {
	WebSecurityConfigurers security.WebSecurityConfigurers
	Providers              coreAuth.Providers
	PasswordEncoder        password.Encoder
	UserCache              userdetails.UserCache
	PreChecker             userdetails.PreChecker
	PostChecker            userdetails.PostChecker
	UserDetailsService     userdetails.Service
	ClientDetailsService   clientdetails.Service
}

// OAuth2Container OAuth2 容器
type OAuth2Container struct {
	Config                      config.OAuth2
	DefaultTokenServices        *token.DefaultTokenServices
	TokenStore                  token.Store
	JwtAccessTokenConverter     *store.JwtAccessTokenConverter
	AccessTokenConverter        token.AccessTokenConverter
	UserAuthenticationConverter token.UserAuthenticationConverter
}

// ResourceServerContainer 资源服务器容器
type ResourceServerContainer struct {
	ResourceServerTokenServices token.ResourceServerTokenServices
	OAuth2SecurityConfigurer    *config.OAuth2SecurityConfigurer
	TokenExtractor              authentication.TokenExtractor
	AuthenticationManager       coreAuth.Manager
}

// AuthorizationServerContainer 授权服务器容器
type AuthorizationServerContainer struct {
	AuthorizationServerTokenServices token.AuthorizationServerTokenServices
	ConsumerTokenServices            token.ConsumerTokenServices
	TokenEnhancer                    token.Enhancer
	AuthenticationManager            coreAuth.Manager
}
