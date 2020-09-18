package domain

import (
	"time"
)

// User 用户实体
type User struct {
	ID        string     `gorm:"primary_key;size:36;"`
	UserName  string     `gorm:"size:64;index;default:'';not null;"`
	RealName  string     `gorm:"size:64;index;default:'';not null;"`
	Password  string     `gorm:"size:256;default:'';not null;"`
	Status    string     `gorm:"size:10;default:'';not null;"`
	Creator   string     `gorm:"size:36;"`
	CreatedAt time.Time  `gorm:"column:created_time"`
	UpdatedAt time.Time  `gorm:"column:updated_time"`
	DeletedAt *time.Time `gorm:"index"`
}

// TableName for user
func (*User) TableName() string {
	return "gm_user"
}
