package set

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"

	"github.com/google/wire"
)

var daoUserSet = wire.NewSet(wire.Struct(new(dao.User), "*"))
var daoAuthoritySet = wire.NewSet(wire.Struct(new(dao.Authority), "*"))
var daoRoleSet = wire.NewSet(wire.Struct(new(dao.Role), "*"))
var daoRoleUserSet = wire.NewSet(wire.Struct(new(dao.RoleUser), "*"))
var daoRoleAuthoritySet = wire.NewSet(wire.Struct(new(dao.RoleAuthority), "*"))
var daoOauthClientDetailsSet = wire.Struct(new(dao.OauthClientDetails), "*")

// DaoSet inject
var DaoSet = wire.NewSet(
	daoUserSet,
	daoAuthoritySet,
	daoRoleSet,
	daoRoleUserSet,
	daoRoleAuthoritySet,
	daoOauthClientDetailsSet,
)
