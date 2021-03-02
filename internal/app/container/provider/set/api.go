package set

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/core/http"

	coreApi "github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
)

var apiConfigSet = wire.NewSet(wire.Struct(new(http.APIConfig), "*"), wire.Bind(new(coreApi.HTTPConfigurer), new(*http.APIConfig)))
var apiOAuth2 = wire.NewSet(wire.Struct(new(api.OAuth2), "*"))

// APISet api注入
var APISet = wire.NewSet(
	apiConfigSet,
	apiOAuth2,
)
