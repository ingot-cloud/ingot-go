package domain

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// SysMenu 菜单
type SysMenu struct {
	ID        types.ID `gorm:"primary_key;size:20"`
	TenantID  int
	PID       types.ID
	Name      string
	Path      string
	ViewPath  string
	Icon      string
	Sort      int
	Cache     uint8
	Hidden    uint8
	Params    string
	Status    string
	Remark    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TableName 表名
func (*SysMenu) TableName() string {
	return "sys_menu"
}
