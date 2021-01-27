// +build wireinject
// The build tag makes sure the stub is not built in the final build

package injector

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/provider"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
)

func BuildConfiguration(options *config.Options) (*config.Config, error) {
	wire.Build(provider.LoadConfig)
	return nil, nil
}

func BuildContainer(config *config.Config, options *config.Options) (*container.Container, func(), error) {
	wire.Build(
		provider.APISet,
		provider.RouterSet,
		provider.ServiceSet,
		provider.DaoSet,
		provider.BuildGorm,
		provider.BuildAuthentication,
		provider.BuildCasbin,
		provider.CasbinAdapterSet,
		provider.BuildPasswordEncoder,
		provider.ConfigSet,
		container.ContainerSet,
	)
	return nil, nil, nil
}
