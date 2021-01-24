package domain

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"

// SysRoleMenu 关联
type SysRoleMenu struct {
	RoleID types.ID `gorm:"primary_key;size:20"`
	MenuID types.ID `gorm:"primary_key;size:20"`
}

// TableName 表名
func (*SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
