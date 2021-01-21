// +build wireinject
// The build tag makes sure the stub is not built in the final build

package injector

import (
	"github.com/ingot-cloud/ingot-go/internal/app/core/container"
	"github.com/ingot-cloud/ingot-go/internal/app/core/provider"

	"github.com/google/wire"
)

func BuildContainer() (*container.Container, func(), error) {
	wire.Build(
		provider.APISet,
		provider.RouterSet,
		provider.ServiceSet,
		provider.DaoSet,
		provider.BuildHTTPHandler,
		provider.BuildGorm,
		provider.BuildAuthentication,
		provider.BuildCasbin,
		provider.CasbinAdapterSet,
		provider.BuildPasswordEncoder,
		container.ContainerSet,
	)
	return new(container.Container), nil, nil
}
