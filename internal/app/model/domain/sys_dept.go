package domain

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// SysDept 部门
type SysDept struct {
	ID        types.ID `gorm:"primary_key;size:20"`
	Version   int64
	TenantID  int
	PID       types.ID
	Name      string
	Scope     string
	Sort      int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TableName 表名
func (*SysDept) TableName() string {
	return "sys_dept"
}
