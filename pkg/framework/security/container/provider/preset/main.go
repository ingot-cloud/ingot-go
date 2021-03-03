package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// All 所有实例
var All = wire.NewSet(
	NilSecurityInjector,
	CommonFields,
	Common,
	OAuth2ContainerFields,
	OAuth2Container,
	AuthorizationServerContainerFields,
	AuthorizationServerContainer,
	ResourceServerContainerFields,
	ResourceServerContainer,
	AuthProvidersContainer,
	AuthProvidersContainerFields,
)

// SecurityContainer 安全容器
var SecurityContainer = wire.NewSet(new(container.SecurityContainer), "*")
