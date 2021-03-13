package basic

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// AuthenticationProvider basic
type AuthenticationProvider struct {
	*dao.AuthenticationProvider
}

// NewProvider 实例化
func NewProvider(encoder password.Encoder, service clientdetails.Service, cache userdetails.UserCache, preChecker userdetails.PreChecker, postChecker userdetails.PostChecker) *AuthenticationProvider {
	return &AuthenticationProvider{
		AuthenticationProvider: &dao.AuthenticationProvider{
			PasswordEncoder:          encoder,
			UserDetailsService:       clientdetails.NewClientDetailsUserDetailsService(service),
			UserCache:                cache,
			PreAuthenticationChecks:  preChecker,
			PostAuthenticationChecks: postChecker,
		},
	}
}

// Supports 该身份验证提供者是否支持指定的认证信息
func (p *AuthenticationProvider) Supports(auth interface{}) bool {
	_, ok := auth.(*authentication.ClientUsernamePasswordAuthenticationToken)
	return ok
}

// Authenticate 身份验证
func (p *AuthenticationProvider) Authenticate(auth core.Authentication) (core.Authentication, error) {
	userAuth, _ := auth.(*authentication.ClientUsernamePasswordAuthenticationToken)
	return p.AuthenticationProvider.Authenticate(userAuth.UsernamePasswordAuthenticationToken)
}
