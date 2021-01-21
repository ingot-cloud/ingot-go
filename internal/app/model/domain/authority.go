package domain

import "time"

// Authority 权限
type Authority struct {
	ID        string     `gorm:"primary_key;size:36"`
	Name      string     `gorm:"size:20;not null"`
	Path      string     `gorm:"size:255;not null"`
	Status    string     `gorm:"size:10;default:'enable';not null"`
	Creator   string     `gorm:"size:36;"`
	CreatedAt time.Time  `gorm:"column:created_time"`
	UpdatedAt time.Time  `gorm:"column:updated_time"`
	DeletedAt *time.Time `gorm:"index"`
}

// TableName for authority
func (*Authority) TableName() string {
	return "gm_authority"
}
