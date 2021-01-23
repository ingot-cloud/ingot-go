package domain

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// SysRole 角色
type SysRole struct {
	ID        types.ID `gorm:"primary_key;size:20"`
	Version   int64
	TenantID  int
	Name      string
	Code      string
	Type      string
	Status    string
	Remark    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// SysRoles 角色列表
type SysRoles []*SysRole
