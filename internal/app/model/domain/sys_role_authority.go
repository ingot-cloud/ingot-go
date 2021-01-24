package domain

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"

// SysRoleAuthority 关联
type SysRoleAuthority struct {
	RoleID      types.ID `gorm:"primary_key;size:20"`
	AuthorityID types.ID `gorm:"primary_key;size:20"`
}

// TableName 表名
func (*SysRoleAuthority) TableName() string {
	return "sys_role_authority"
}
