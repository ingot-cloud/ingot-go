package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"

	"gorm.io/gorm"
)

func getUserDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.SysUser))
}

// User Dao
type User struct {
	DB *gorm.DB
}

// List user list
func (u *User) List(ctx context.Context, condition dto.QueryCondition) (*[]*domain.SysUser, error) {
	db := getUserDB(ctx, u.DB)

	if status := condition.Status; status != "" {
		db = db.Where("status = ?", status)
	}

	var list []*domain.SysUser
	err := db.Scan(&list).Error

	return &list, err
}

// GetByID 根据ID获取User
func (u *User) GetByID(ctx context.Context, id string) (*domain.SysUser, error) {
	db := getUserDB(ctx, u.DB)

	db = db.Where("id = ?", id)

	var user domain.SysUser
	err := db.First(&user).Error

	return &user, err
}

// One user
func (u *User) One(ctx context.Context, username string) (*domain.SysUser, error) {
	db := getUserDB(ctx, u.DB)
	db = db.Where("username = ?", username)

	var user domain.SysUser
	err := db.First(&user).Error

	log.Debug("user = %v", user)

	return &user, err
}
