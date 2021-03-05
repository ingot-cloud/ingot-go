package set

import (
	"github.com/casbin/casbin/v2/persist"
	"github.com/ingot-cloud/ingot-go/internal/app/service"
	"github.com/ingot-cloud/ingot-go/internal/app/service/impl"

	"github.com/google/wire"
)

// ServiceSet inject
var ServiceSet = wire.NewSet(
	wire.Struct(new(impl.CasbinAdapterService), "*"),
	wire.Bind(new(persist.Adapter), new(*impl.CasbinAdapterService)),

	wire.Struct(new(impl.Permission), "*"),
	wire.Bind(new(service.Permission), new(*impl.Permission)),
)
