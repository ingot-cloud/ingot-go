package domain

import (
	"github.com/ingot-cloud/ingot-go/internal/app/common/utils"
	"github.com/ingot-cloud/ingot-go/internal/app/common/uuid"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"time"

	"gorm.io/gorm"
)

// Users for domain
type Users []*User

// To domain to dto
func (users Users) To() dto.Users {
	target := make(dto.Users, len(users))
	for i, item := range users {
		target[i] = item.To()
	}
	return target
}

// User 用户实体
type User struct {
	ID            string         `gorm:"primary_key;size:36;"`
	Username      string         `gorm:"size:64;uniqueIndex;default:'';not null;"`
	Password      string         `gorm:"size:256;default:'';not null;"`
	RealName      string         `gorm:"size:64;default:'';not null;"`
	Phone         string         `gorm:"size:20;uniqueIndex;'';not null;"`
	Status        string         `gorm:"size:10;default:'enable';not null;"`
	Remark        string         `gorm:"size:255;default:'';"`
	LastLoginTime *time.Time     `gorm:"column:last_login_time"`
	Creator       string         `gorm:"size:36;"`
	CreatedAt     time.Time      `gorm:"column:created_time"`
	UpdatedAt     time.Time      `gorm:"column:updated_time"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// TableName for user
func (*User) TableName() string {
	return "gm_user"
}

// BeforeCreate hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.MustString()
	return
}

// To domain to dto
func (u User) To() *dto.User {
	target := new(dto.User)
	utils.Copy(u, target)
	return target
}
