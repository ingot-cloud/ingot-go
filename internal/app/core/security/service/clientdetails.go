package service

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// ClientDetails 服务
type ClientDetails struct {
	OauthClientDetailsDao *dao.OauthClientDetails
}

// LoadClientByClientID 根据 clientID 获取客户端详细信息
func (c *ClientDetails) LoadClientByClientID(clientID string) (clientdetails.ClientDetails, error) {
	return c.OauthClientDetailsDao.GetByID(context.TODO(), clientID)
}
