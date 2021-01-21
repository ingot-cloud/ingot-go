package provider

import (
	"github.com/ingot-cloud/ingot-go/internal/app/core"
	"github.com/ingot-cloud/ingot-go/internal/app/router"

	"github.com/google/wire"
)

// RouterSet for router
var RouterSet = wire.NewSet(wire.Struct(new(router.Router), "*"), wire.Bind(new(core.IRouter), new(*router.Router)))
