// +build wireinject
// The build tag makes sure the stub is not built in the final build

package injector

import (
	"ingot/core/container"
	"ingot/core/provider"

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
		container.ContainerSet,
	)
	return new(container.Container), nil, nil
}
