package impl

import (
	"github.com/ingot-cloud/ingot-go/internal/app/model/domain"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/types"
)

// OAuth2Client 实现
type OAuth2Client struct {
}

// GetClientsByRoles 获取指定角色绑定的所有client
func (c *OAuth2Client) GetClientsByRoles(ids types.ID) []*domain.SysOauthClientDetails {
	return nil
}
