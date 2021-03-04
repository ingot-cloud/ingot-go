package dao

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"gorm.io/gorm"
)

func getOauthClientDetailsDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, db, new(domain.SysOauthClientDetails))
}

// OauthClientDetails Dao
type OauthClientDetails struct {
	DB *gorm.DB
}

// GetByID 根据ID获取Client信息
func (c *OauthClientDetails) GetByID(context context.Context, clientID string) (*domain.SysOauthClientDetails, error) {
	db := getOauthClientDetailsDB(context, c.DB).Where("client_id = ?", clientID)

	client := new(domain.SysOauthClientDetails)
	err := db.First(client).Error

	return client, err
}
