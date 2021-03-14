package pre

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/process"
)

// All 所有实例
var All = wire.NewSet(
	CommonContainer,
	OAuth2Container,
	AuthorizationServerContainer,
	ResourceServerContainer,
	AuthProvidersContainer,
	SecurityPre,
)

// Security 安全容器
var SecurityPre = wire.NewSet(
	wire.NewSet(wire.Struct(new(container.NilSecurityInjector), "*")),

	wire.Struct(new(container.SecurityContainerImpl), "*"),
	wire.Bind(new(container.SecurityContainerPre), new(*container.SecurityContainerImpl)),

	wire.Struct(new(container.SecurityContainerPreProxyImpl), "*"),
	wire.Bind(new(container.SecurityContainerPreProxy), new(*container.SecurityContainerPreProxyImpl)),

	InjectCustomInstance,
)

// InjectCustomInstance 注入自定义实例
func InjectCustomInstance(proxy container.SecurityContainerPreProxy) container.SecurityContainerCombine {
	injector := proxy.GetSecurityInjector()
	sc := proxy.GetSecurityContainer()
	return process.DoPre(injector, sc)
}
