package domain

import (
	"time"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// SysSocialDetails 社交信息
type SysSocialDetails struct {
	ID          types.ID `gorm:"primary_key;size:20"`
	TenantID    int64
	AppID       string
	AppSecret   string
	RedirectURL string
	Name        string
	Type        string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// TableName 表名
func (*SysSocialDetails) TableName() string {
	return "sys_social_details"
}
