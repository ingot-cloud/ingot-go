package domain

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// SysAuthority 权限
type SysAuthority struct {
	ID        types.ID `gorm:"primary_key;size:20"`
	Version   int64
	TenantID  int
	PID       int64
	Name      string
	Code      string
	Path      string
	Method    string
	Status    string
	remark    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// SysAuthoritys 权限列表
type SysAuthoritys []*SysAuthority
