// +build wireinject
// The build tag makes sure the stub is not built in the final build

package injector

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"

	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/post"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/pre"
)

func BuildConfiguration(options *config.Options) (*config.Config, error) {
	wire.Build(provider.NewConfig)
	return nil, nil
}

func BuildContainerPre(config *config.Config, options *config.Options) (securityContainer.SecurityContainerCombine, func(), error) {
	wire.Build(
		provider.AllSet,
		provider.AllFactory,
		provider.SecurityInjector,

		pre.All,
	)
	return nil, nil, nil
}

func BuildContainerPost(config *config.Config, options *config.Options, combine securityContainer.SecurityContainerCombine) (container.Container, func(), error) {
	wire.Build(
		provider.AllSet,
		provider.AllFactory,
		provider.SecurityInjector,

		post.All,

		container.BootContainer,
	)
	return nil, nil, nil
}
