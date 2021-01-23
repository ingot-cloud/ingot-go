package provider

import (
	"github.com/ingot-cloud/ingot-go/internal/app/service"

	"github.com/google/wire"
)

var serviceAuthSet = wire.NewSet(wire.Struct(new(service.Auth), "*"))
var servicePermissionSet = wire.NewSet(wire.Struct(new(service.Permission), "*"))

// ServiceSet inject
var ServiceSet = wire.NewSet(
	serviceAuthSet,
	servicePermissionSet,
)
