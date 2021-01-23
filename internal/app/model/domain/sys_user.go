package domain

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"

// SysUser 用户表
type SysUser struct {
	ID        types.ID `gorm:"primary_key;size:20"`
	Version   int64
	TenantID  int
	DeptID    int64
	Username  string
	Password  string
	RealName  string
	Phone     string
	Email     string
	Status    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

// SysUsers 用户列表
type SysUsers []*SysUser
