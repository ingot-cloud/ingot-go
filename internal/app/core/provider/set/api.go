package set

import (
	"github.com/ingot-cloud/ingot-go/internal/app/api"

	"github.com/google/wire"
)

var apiLoginSet = wire.NewSet(wire.Struct(new(api.Auth), "*"))

// APISet inject
var APISet = wire.NewSet(
	apiLoginSet,
)
