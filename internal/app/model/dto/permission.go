package dto

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// RolePolicy 角色策略
type RolePolicy struct {
	RoleID        types.ID
	TenantID      int
	AuthorityList domain.SysAuthoritys
}

// RolePolicys 角色策略列表
type RolePolicys []*RolePolicy

// UserPolicy 用户策略
type UserPolicy struct {
	UserID   types.ID
	TenantID int
	RoleList []types.ID
}

// UserPolicys 用户策略列表
type UserPolicys []*UserPolicy
