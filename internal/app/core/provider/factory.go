package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/core/provider/factory"
)

var AllFactory = wire.NewSet(
	factory.Config,
	factory.NewAuthentication,
	factory.NewCasbin,
	factory.NewGorm,
	factory.NewPasswordEncoder,
)
