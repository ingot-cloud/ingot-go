package provider

import (
	"ingot/core"
	"ingot/router"

	"github.com/google/wire"
)

// RouterSet for router
var RouterSet = wire.NewSet(wire.Struct(new(router.Router), "*"), wire.Bind(new(core.IRouter), new(*router.Router)))
