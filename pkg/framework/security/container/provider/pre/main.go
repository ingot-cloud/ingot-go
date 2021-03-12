package pre

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/process"
)

// All 所有实例
var All = wire.NewSet(
	CommonContainerFields,
	CommonContainer,
	OAuth2ContainerFields,
	OAuth2Container,
	AuthorizationServerContainerFields,
	AuthorizationServerContainer,
	ResourceServerContainerFields,
	ResourceServerContainer,
	AuthProvidersContainer,
	AuthProvidersContainerFields,
	Security,
	InjectCustomInstance,
)

// Security 安全容器
var Security = wire.NewSet(
	wire.NewSet(wire.Struct(new(container.NilSecurityInjector), "*")),

	wire.Struct(new(container.SecurityContainerImpl), "*"),
	wire.Bind(new(container.SecurityContainer), new(*container.SecurityContainerImpl)),

	wire.Struct(new(container.SecurityContainerProxyImpl), "*"),
	wire.Bind(new(container.SecurityContainerProxy), new(*container.SecurityContainerProxyImpl)),
)

// InjectCustomInstance 注入自定义实例
func InjectCustomInstance(proxy container.SecurityContainerProxy) container.SecurityContainerCombine {
	injector := proxy.GetSecurityInjector()
	sc := proxy.GetSecurityContainer()
	return process.DoPre(injector, sc)
}
