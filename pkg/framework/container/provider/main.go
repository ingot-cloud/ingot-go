package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
)

// BootContainer 构建 Container
var BootContainer = wire.NewSet(

	wire.Struct(new(container.DefaultContainerPre), "*"),
	wire.Bind(new(container.ContainerPre), new(*container.DefaultContainerPre)),

	wire.Struct(new(container.DefaultContainerInjector), "*"),

	BuildContainerProcess,

	PrintInjectInstance,
)
