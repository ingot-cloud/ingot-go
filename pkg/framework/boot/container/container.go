package container

import (
	"context"

	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// Container for app
type Container struct {
	HTTPConfig                   config.HTTPConfig
	HTTPConfigurer               config.HTTPConfigurer
	SecurityContainer            *container.SecurityContainer
	OAuth2Container              *container.OAuth2Container
	ResourceServerContainer      *container.ResourceServerContainer
	AuthorizationServerContainer *container.AuthorizationServerContainer
}

// Factory 容器工厂
type Factory func(context.Context) (*Container, func(), error)

// ContainerSet 容器
var ContainerSet = wire.NewSet(wire.Struct(new(Container), "*"))
