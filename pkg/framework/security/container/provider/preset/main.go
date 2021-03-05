package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// All 所有实例
var All = wire.NewSet(
	NilSecurityInjector,
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
	SecurityContainer,
)

// SecurityContainer 安全容器
var SecurityContainer = wire.NewSet(
	wire.Struct(new(container.SecurityContainerImpl), "*"),
	wire.Bind(new(container.SecurityContainer), new(*container.SecurityContainerImpl)),
)
