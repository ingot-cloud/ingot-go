package set

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/core/http"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
)

var apiConfigSet = wire.NewSet(wire.Struct(new(http.APIConfig), "*"), wire.Bind(new(config.HTTPConfigurer), new(*http.APIConfig)))
var apiAuthSet = wire.NewSet(wire.Struct(new(api.Auth), "*"))

// APISet api注入
var APISet = wire.NewSet(
	apiConfigSet,
	apiAuthSet,
)
