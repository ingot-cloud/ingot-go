package user

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
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
