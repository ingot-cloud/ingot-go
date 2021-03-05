package set

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"

	"github.com/google/wire"
)

// DaoSet inject
var DaoSet = wire.NewSet(
	wire.Struct(new(dao.User), "*"),
	wire.Struct(new(dao.Authority), "*"),
	wire.Struct(new(dao.Role), "*"),
	wire.Struct(new(dao.RoleUser), "*"),
	wire.Struct(new(dao.RoleAuthority), "*"),
	wire.Struct(new(dao.OauthClientDetails), "*"),
)
