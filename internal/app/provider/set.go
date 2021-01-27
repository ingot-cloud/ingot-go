package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/provider/set"
)

// AllSet 所有的结构集
var AllSet = wire.NewSet(
	set.APISet,
	set.DaoSet,
	set.RouterSet,
	set.ServiceSet,
)
