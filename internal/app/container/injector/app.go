// +build wireinject
// The build tag makes sure the stub is not built in the final build

package injector

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"

	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	securityProvider "github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/preset"
)

func BuildConfiguration(options *config.Options) (*config.Config, error) {
	wire.Build(provider.NewConfig)
	return nil, nil
}

func BuildContainerInjector(config *config.Config, options *config.Options) (securityContainer.SecurityInjector, func(), error) {
	wire.Build(
		provider.AllSet,
		provider.AllFactory,
		provider.SecurityInjector,
		provider.CustomSecurityAll,

		preset.All,

		container.BootContainerPre,
	)
	return nil, nil, nil
}

func BuildContainer(config *config.Config, options *config.Options, securityInjector securityContainer.SecurityInjector) (container.Container, func(), error) {
	wire.Build(
		provider.AllSet,
		provider.AllFactory,

		securityProvider.All,

		container.BootContainer,
	)
	return nil, nil, nil
}
