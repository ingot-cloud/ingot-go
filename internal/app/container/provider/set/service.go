package set

import (
	"github.com/casbin/casbin/v2/persist"
	"github.com/ingot-cloud/ingot-go/internal/app/service"

	"github.com/google/wire"
)

var servicePermissionSet = wire.NewSet(wire.Struct(new(service.Permission), "*"))
var serviceCasbinAdapterSet = wire.NewSet(wire.Struct(new(service.CasbinAdapterService), "*"), wire.Bind(new(persist.Adapter), new(*service.CasbinAdapterService)))

// ServiceSet inject
var ServiceSet = wire.NewSet(
	servicePermissionSet,
	serviceCasbinAdapterSet,
)
