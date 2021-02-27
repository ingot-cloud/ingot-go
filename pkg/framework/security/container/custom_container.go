package container

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// CustomContainer 自定义注入参数，app端实现
type CustomContainer struct {
	WebSecurityConfigurers security.WebSecurityConfigurers
	Providers              coreAuth.Providers
	PasswordEncoder        password.Encoder
	UserCache              userdetails.UserCache
	PreChecker             userdetails.PreChecker
	PostChecker            userdetails.PostChecker
	UserDetailsService     userdetails.Service
	ClientDetailsService   clientdetails.Service

	TokenStore                  token.Store
	AccessTokenConverter        token.AccessTokenConverter
	UserAuthenticationConverter token.UserAuthenticationConverter

	ResourceServerTokenServices   token.ResourceServerTokenServices
	TokenExtractor                authentication.TokenExtractor
	ResourceAuthenticationManager coreAuth.Manager

	AuthorizationServerTokenServices   token.AuthorizationServerTokenServices
	ConsumerTokenServices              token.ConsumerTokenServices
	TokenEnhancer                      token.Enhancer
	AuthorizationAuthenticationManager coreAuth.Manager
}
