package provider

import (
	"github.com/google/wire"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// AuthProvidersContainer 容器
var AuthProvidersContainer = wire.NewSet(wire.Struct(new(container.AuthProvidersContainer), "*"))

// AuthProvidersContainerFields 所有provider
var AuthProvidersContainerFields = wire.NewSet(
	wire.Struct(new(ProvidersImpl), "*"),
	wire.Bind(new(coreAuth.Providers), new(*ProvidersImpl)),
)

// ProvidersImpl 接口实现
type ProvidersImpl struct {
	SC container.SecurityContainerCombine
}

// Add 追加provider
func (p *ProvidersImpl) Add(item coreAuth.Provider) {
	p.SC.GetAuthProvidersContainer().Providers.Add(item)
}

// Get 获取所有Provider
func (p *ProvidersImpl) Get() []coreAuth.Provider {
	return p.SC.GetAuthProvidersContainer().Providers.Get()
}
