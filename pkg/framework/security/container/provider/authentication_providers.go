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
	DaoAuthenticationProvider,
	BasicAuthenticationProvider,
	wire.Struct(new(ProvidersImpl), "Basic", "Dao"),
	wire.Bind(new(coreAuth.Providers), new(*ProvidersImpl)),
)

// ProvidersImpl 接口实现
type ProvidersImpl struct {
	providers []coreAuth.Provider

	Injector container.SecurityInjector
	Basic    *basic.AuthenticationProvider
	Dao      *dao.AuthenticationProvider
}

// Add 追加provider
func (p *ProvidersImpl) Add(item coreAuth.Provider) {
	p.providers = append(p.providers, item)
}

// Get 获取所有Provider
func (p *ProvidersImpl) Get() []coreAuth.Provider {
	if p.Injector.GetProviders() != nil && len(p.Injector.GetProviders().Get()) != 0 {
		return p.Injector.GetProviders().Get()
	}
	p.providers = append(p.providers, p.Basic)
	p.providers = append(p.providers, p.Dao)
	return p.providers
}

// DaoAuthenticationProvider UsernamePasswordAuthenticationToken 认证提供者
var DaoAuthenticationProvider = wire.NewSet(wire.Struct(new(dao.AuthenticationProvider), "*"))

// BasicAuthenticationProvider 认证提供者，其中注入了 ClientDetailsUserDetailsService
func BasicAuthenticationProvider(common *container.CommonContainer) *basic.AuthenticationProvider {
	return preset.BasicAuthenticationProvider(common)
}
