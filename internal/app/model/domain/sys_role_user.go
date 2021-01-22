package domain

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"

// SysRoleUser 关联
type SysRoleUser struct {
	RoleID types.ID `gorm:"primary_key;size:20"`
	UserID types.ID `gorm:"primary_key;size:20"`
}
