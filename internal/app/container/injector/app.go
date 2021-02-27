// +build wireinject
// The build tag makes sure the stub is not built in the final build

package injector

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	app "github.com/ingot-cloud/ingot-go/internal/app/container"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

func BuildConfiguration(options *config.Options) (*config.Config, error) {
	wire.Build(provider.NewConfig)
	return nil, nil
}

func BuildAppContainer(config *config.Config, options *config.Options) (*app.AppContainer, func(), error) {
	wire.Build(
		provider.AllSet,
		provider.AllFactory,
		provider.AppContainer,
	)
	return nil, nil, nil
}

func BuildContainer(httpInjector container.HTTPInjector, securityContainer securityContainer.SecurityAllContainer) (*container.Container, func(), error) {
	wire.Build(
		container.BuildContainer,
	)
	return nil, nil, nil
}
