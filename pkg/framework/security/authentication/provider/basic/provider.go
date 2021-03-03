package basic

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/crypto/password"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// AuthenticationProvider basic
type AuthenticationProvider struct {
	*dao.AuthenticationProvider
}

// NewProvider 实例化
func NewProvider(encoder password.Encoder, service clientdetails.Service, cache userdetails.UserCache, preChecker userdetails.Checker, postChecker userdetails.Checker) *AuthenticationProvider {
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