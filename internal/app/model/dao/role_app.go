package dao

import (
	"context"
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"

	"gorm.io/gorm"
)

// RoleApp Dao
type RoleApp struct {
	DB *gorm.DB
}

func getRoleAppDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.RoleApp))
}

// GetAppRole 获取App的角色
func (ra *RoleApp) GetAppRole(context context.Context, appID string) ([]domain.RoleApp, error) {
	db := getRoleAppDB(context, ra.DB).Where("app_id = ?", appID)

	var list []domain.RoleApp
	err := db.Find(&list).Error

	return list, err
}
