package container

import (
	"context"

	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// Factory 容器工厂
type Factory func(context.Context) (Container, func(), error)

// BootContainerPre 构建 Container
var BootContainerPre = wire.NewSet(
	wire.Struct(new(DefaultPre), "*"),
	wire.Bind(new(Container), new(*DefaultPre)),
)

// BootContainer 构建 Container
var BootContainer = wire.NewSet(
	wire.Struct(new(DefaultContainer), "*"),
	wire.Bind(new(Container), new(*DefaultContainer)),
)

// Container boot 容器
type Container interface {
	GetHTTPConfig() config.HTTPConfig
	GetHTTPConfigurer() api.HTTPConfigurer
	GetSecurityInjector() container.SecurityInjector
	GetSecurityContainer() *container.Common
	GetOAuth2Container() *container.OAuth2Container
	GetResourceServerContainer() *container.ResourceServerContainer
	GetAuthorizationServerContainer() *container.AuthorizationServerContainer
}
