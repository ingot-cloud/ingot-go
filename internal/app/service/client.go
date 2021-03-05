package service

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// OAuth2Client 服务
type OAuth2Client interface {
	// 获取指定角色绑定的所有client
	GetClientsByRoles(ids types.ID) []*domain.SysOauthClientDetails
}
