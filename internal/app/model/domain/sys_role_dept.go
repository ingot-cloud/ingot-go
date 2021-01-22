package domain

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"

// SysRoleDept 关联
type SysRoleDept struct {
	RoleID types.ID `gorm:"primary_key;size:20"`
	DeptID types.ID `gorm:"primary_key;size:20"`
}
