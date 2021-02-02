package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider/factory"
)

// AllFactory 所有的
var AllFactory = wire.NewSet(
	factory.Config,
	factory.NewAuthentication,
	factory.NewCasbin,
	factory.NewGorm,
	factory.NewPasswordEncoder,
	factory.NewIDGenerator,
)
