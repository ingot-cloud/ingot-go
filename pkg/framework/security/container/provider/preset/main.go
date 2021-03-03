package preset

import "github.com/google/wire"

// All 所有实例
var All = wire.NewSet(
	NilSecurityInjector,
	SecurityContainerFields,
	SecurityContainer,
	OAuth2ContainerFields,
	OAuth2Container,
	AuthorizationServerContainerFields,
	AuthorizationServerContainer,
	ResourceServerContainerFields,
	ResourceServerContainer,
	AuthProvidersContainer,
	AuthProvidersContainerFields,
)
