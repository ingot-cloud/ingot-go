package domain

import "time"

// SysTenant 租户
type SysTenant struct {
	ID        int `gorm:"primary_key;size:11"`
	Version   int64
	Name      string
	Code      string
	StartAt   time.Time
	EndAt     time.Time
	Status    string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt *time.Time
}

// TableName 表名
func (*SysTenant) TableName() string {
	return "sys_tenant"
}
