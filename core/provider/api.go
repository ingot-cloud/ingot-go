package provider

import (
	"ingot/api"

	"github.com/google/wire"
)

var apiTestSet = wire.NewSet(wire.Struct(new(api.Test), "*"))

// APISet inject
var APISet = wire.NewSet(apiTestSet)
