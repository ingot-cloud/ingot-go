package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"

	"gorm.io/gorm"
)

func getRoleUserDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.RoleUser))
}

// RoleUser DAO
type RoleUser struct {
	DB *gorm.DB
}

// List for role user
func (ru *RoleUser) List(ctx context.Context, params dto.RoleUserQueryParams) (*dto.RoleUserQueryResult, error) {
	db := getRoleUserDB(ctx, ru.DB)

	var list domain.RoleUsers
	db.Find(&list)

	return &dto.RoleUserQueryResult{
		List: list.To(),
	}, nil
}

// GetUserRole 获取用户的角色
func (ru *RoleUser) GetUserRole(ctx context.Context, userID string) (*dto.RoleUserQueryResult, error) {
	db := getRoleUserDB(ctx, ru.DB).Where("user_id = ?", userID)

	var list domain.RoleUsers
	err := db.Find(&list).Error

	return &dto.RoleUserQueryResult{
		List: list.To(),
	}, err
}

// BindRoleToUser 绑定角色到用户
func (ru *RoleUser) BindRoleToUser(ctx context.Context, params dto.RoleUser) error {
	db := getRoleUserDB(ctx, ru.DB)

	roleUser := new(domain.RoleUser)
	utils.Copy(params, roleUser)

	result := db.Create(roleUser)

	return result.Error
}
