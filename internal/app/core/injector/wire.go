// +build wireinject
// The build tag makes sure the stub is not built in the final build

package injector

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/core/provider"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
)

func BuildContainer(config server.Config) (*container.Container, func(), error) {
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
		container.ContainerSet,
	)
	return nil, nil, nil
}
