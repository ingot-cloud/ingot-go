package provider

import (
	"ingot/service"

	"github.com/google/wire"
)

var serviceTestSet = wire.NewSet(wire.Struct(new(service.Test), "*"))

// ServiceSet inject
var ServiceSet = wire.NewSet(serviceTestSet)
