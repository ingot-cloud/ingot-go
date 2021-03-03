package provider

import (
	"github.com/google/wire"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/basic"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/preset"
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
func Providers(dao *dao.AuthenticationProvider, injector container.SecurityInjector) coreAuth.Providers {
	if len(injector.GetProviders()) != 0 {
		return injector.GetProviders()
	}
	return preset.Providers(dao)
}

// DaoAuthenticationProvider UsernamePasswordAuthenticationToken 认证提供者
var DaoAuthenticationProvider = wire.NewSet(wire.Struct(new(dao.AuthenticationProvider), "*"))

// BasicAuthenticationProvider 认证提供者，其中注入了 ClientDetailsUserDetailsService
func BasicAuthenticationProvider(common *container.CommonContainer) *basic.AuthenticationProvider {
	return preset.BasicAuthenticationProvider(common)
}
