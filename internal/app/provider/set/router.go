package set

import (
	"github.com/ingot-cloud/ingot-go/internal/app/router"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"

	"github.com/google/wire"
)

// RouterSet for router
var RouterSet = wire.NewSet(wire.Struct(new(router.Router), "*"), wire.Bind(new(server.Router), new(*router.Router)))
