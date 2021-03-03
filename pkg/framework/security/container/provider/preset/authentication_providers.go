package preset

import (
	"github.com/google/wire"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/basic"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// AuthProvidersContainer 容器
var AuthProvidersContainer = wire.NewSet(wire.Struct(new(container.AuthProvidersContainer), "*"))

// AuthProvidersContainerFields 所有provider
var AuthProvidersContainerFields = wire.NewSet(
	Providers,
	DaoAuthenticationProvider,
	BasicAuthenticationProvider,
)

// Providers 所有认证提供者
func Providers(dao *dao.AuthenticationProvider) coreAuth.Providers {
	var providers coreAuth.Providers
	providers = append(providers, dao)
	return providers
}

// DaoAuthenticationProvider UsernamePasswordAuthenticationToken 认证提供者
var DaoAuthenticationProvider = wire.NewSet(wire.Struct(new(dao.AuthenticationProvider), "*"))

// BasicAuthenticationProvider 认证提供者，其中注入了 ClientDetailsUserDetailsService
func BasicAuthenticationProvider(common *container.CommonContainer) *basic.AuthenticationProvider {
	return basic.NewProvider(common.PasswordEncoder, common.ClientDetailsService, common.UserCache, common.PreChecker, common.PostChecker)
}