package preset

import (
	"github.com/google/wire"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
)

// ProviderAll 所有provider
var ProviderAll = wire.NewSet(
	DaoAuthenticationProvider,
	Providers,
)

// DaoAuthenticationProvider UsernamePasswordAuthenticationToken 认证提供者
var DaoAuthenticationProvider = wire.NewSet(wire.Struct(new(dao.AuthenticationProvider), "*"))

// // BasicAuthenticationProvider 认证提供者，其中注入了 ClientDetailsUserDetailsService
// func BasicAuthenticationProvider(encoder password.Encoder, service clientdetails.Service, cache userdetails.UserCache, preChecker userdetails.Checker, postChecker userdetails.Checker) *basic.AuthenticationProvider {
// 	return basic.NewProvider(encoder, service, cache, preChecker, postChecker)
// }

// Providers 所有认证提供者
func Providers(dao *dao.AuthenticationProvider) coreAuth.Providers {
	var providers coreAuth.Providers
	providers = append(providers, dao)
	return providers
}
