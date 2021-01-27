package container

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
)

// Container for app
type Container struct {
	Router     server.Router
	HTTPConfig server.Config
}

// ContainerSet 容器
var ContainerSet = wire.NewSet(wire.Struct(new(Container), "*"))
