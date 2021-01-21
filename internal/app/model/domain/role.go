package domain

import (
	"time"
)

// Role 角色
type Role struct {
	ID        string     `gorm:"primary_key;size:20"`
	Name      string     `gorm:"size:20;not null"`
	Remark    string     `gorm:"size:255;"`
	Status    string     `gorm:"size:10;default:'enable';not null;"`
	Creator   string     `gorm:"size:36;"`
	CreatedAt time.Time  `gorm:"column:created_time"`
	UpdatedAt time.Time  `gorm:"column:updated_time"`
	DeletedAt *time.Time `gorm:"index"`
}

// TableName for role
func (*Role) TableName() string {
	return "gm_role"
}
