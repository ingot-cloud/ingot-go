package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/container"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// AppContainer 容器
var AppContainer = wire.NewSet(
	wire.Struct(new(container.AppContainer), "*"),
	wire.Bind(new(securityContainer.SecurityInjector), new(*container.AppContainer)),
)
