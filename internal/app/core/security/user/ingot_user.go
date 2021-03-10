package user

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
)

// IngotUser 自定义User
type IngotUser struct {
	*userdetails.User

	ID       types.ID
	DeptID   types.ID
	TenantID types.ID
	AuthType string
}

// NewIngotUser 实例化
func NewIngotUser(id, deptID, tenantID types.ID, authType string, user *userdetails.User) *IngotUser {
	return &IngotUser{
		User:     user,
		ID:       id,
		DeptID:   deptID,
		TenantID: tenantID,
		AuthType: authType,
	}
}

// NewPermitIngotUser 实例化，并且userdetails.User中所有检查信息都为true
func NewPermitIngotUser(id, deptID, tenantID types.ID, authType string, username, password string, authorities []core.GrantedAuthority) *IngotUser {
	coreUser := userdetails.NewUserAllParams(username, password, authorities, true, true, true, true)
	return NewIngotUser(id, deptID, tenantID, authType, coreUser)
}
