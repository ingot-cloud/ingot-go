package post

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
	SecurityPost,
	process.PrintInjectInstance,
)

// SecurityPost 安全容器
var SecurityPost = wire.NewSet(
	wire.NewSet(wire.Struct(new(container.NilSecurityInjector), "*")),

	wire.Struct(new(container.SecurityContainerImpl), "*"),
	wire.Bind(new(container.SecurityContainerPost), new(*container.SecurityContainerImpl)),

	wire.Struct(new(container.SecurityContainerPostProxyImpl), "*"),
	wire.Bind(new(container.SecurityContainerPostProxy), new(*container.SecurityContainerPostProxyImpl)),

	InjectCustomInstance,
)

// InjectCustomInstance 注入自定义实例
func InjectCustomInstance(proxy container.SecurityContainerPostProxy) container.SecurityContainer {
	injector := proxy.GetSecurityInjector()
	sc := proxy.GetSecurityContainer()
	return process.DoPost(injector, sc)
}
