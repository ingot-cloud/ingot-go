package container

import (
	"context"

	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
)

// Container for app
type Container struct {
	Router     server.Router
	HTTPConfig server.Config
}

// Factory 容器工厂
type Factory func(context.Context) (*Container, func(), error)

// ContainerSet 容器
var ContainerSet = wire.NewSet(wire.Struct(new(Container), "*"))
