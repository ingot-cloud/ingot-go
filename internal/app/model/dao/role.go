package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dto"
	"gorm.io/gorm"
)

func getRoleDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.SysRole))
}

// Role DAO
type Role struct {
	DB *gorm.DB
}

// GetRoleByID 根据ID获取角色
func (r *Role) GetRoleByID(ctx context.Context, condition dto.QueryCondition) (*domain.SysRole, error) {
	db := getRoleDB(ctx, r.DB)

	if id := condition.ID; id != 0 {
		db = db.Where("id = ?", id)
	}
	if status := condition.Status; status != "" {
		db = db.Where("status = ?", status)
	}

	var role domain.SysRole
	err := db.Find(&role).Error
	return &role, err
}

// List 获取角色列表
func (r *Role) List(ctx context.Context, condition dto.QueryCondition) (*[]*domain.SysRole, error) {
	db := getRoleDB(ctx, r.DB)

	if status := condition.Status; status != "" {
		db = db.Where("status = ?", status)
	}
	if ids := condition.IDs; len(ids) != 0 {
		db = db.Where("id IN ?", ids)
	}

	var roles []*domain.SysRole
	err := db.Scan(&roles).Error

	return &roles, err
}
