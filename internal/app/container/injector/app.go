// +build wireinject
// The build tag makes sure the stub is not built in the final build

package injector

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	securityProvider "github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider"
)

func BuildConfiguration(options *config.Options) (*config.Config, error) {
	wire.Build(provider.NewConfig)
	return nil, nil
}

// func BuildAppContainer(config *config.Config, options *config.Options) (*app.AppContainer, func(), error) {
// 	wire.Build(
// 		provider.AllSet,
// 		provider.AllFactory,
// 		provider.AppContainer,
// 	)
// 	return nil, nil, nil
// }

func BuildContainer(config *config.Config, options *config.Options) (*container.Container, func(), error) {
	wire.Build(
		provider.AllSet,
		provider.AllFactory,
		provider.AppContainer,
		securityProvider.SecurityContainerFields,
		securityProvider.SecurityContainer,
		securityProvider.OAuth2ContainerFields,
		securityProvider.OAuth2Container,
		securityProvider.AuthorizationServerContainerFields,
		securityProvider.AuthorizationServerContainer,
		securityProvider.ResourceServerContainerFields,
		securityProvider.ResourceServerContainer,

		container.BootContainer,
	)
	return nil, nil, nil
}
