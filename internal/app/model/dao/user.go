package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/common/log"
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"

	"gorm.io/gorm"
)

func getUserDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.User))
}

// User Dao
type User struct {
	DB *gorm.DB
}

// List user list
func (u *User) List(ctx context.Context, params dto.UserQueryParams) (*dto.UserQueryResult, error) {

	db := getUserDB(ctx, u.DB)
	if p := params.Username; p != "" {
		db = db.Where("username = ?", p)
	}
	if p := params.Status; p != "" {
		db = db.Where("status = ?", p)
	}

	var list domain.Users
	err := db.Find(&list).Error

	return &dto.UserQueryResult{
		List: list.To(),
	}, err
}

// GetByID 根据ID获取User
func (u *User) GetByID(ctx context.Context, id string) (*domain.User, error) {
	db := getUserDB(ctx, u.DB)

	db = db.Where("id = ?", id)

	var user domain.User
	err := db.First(&user).Error

	return &user, err
}

// One user
func (u *User) One(ctx context.Context, username string) (*domain.User, error) {
	db := getUserDB(ctx, u.DB)
	db = db.Where("username = ?", username)

	var user domain.User
	err := db.First(&user).Error

	log.Debug("user = %v", user)

	return &user, err
}
