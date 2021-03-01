package preset

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// NilSecurityInjector 默认实例
var NilSecurityInjector = wire.NewSet(wire.Struct(new(container.NilSecurityInjector), "*"))
