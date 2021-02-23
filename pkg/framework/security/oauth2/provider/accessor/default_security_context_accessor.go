package accessor

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
)

// DefaultSecurityContextAccessor 默认实现
type DefaultSecurityContextAccessor struct {
}

// IsUser 当前身份验证信息是否为用户身份验证信息
func (a *DefaultSecurityContextAccessor) IsUser(ctx *ingot.Context) bool {
	auth := ctx.GetAuthentication()
	if auth == nil {
		return false
	}
	if oauth, ok := auth.(*authentication.OAuth2Authentication); ok {
		return oauth.UserAuthentication != nil
	}
	return true
}

// GetAuthorities 获取权限
func (a *DefaultSecurityContextAccessor) GetAuthorities(ctx *ingot.Context) []core.GrantedAuthority {
	auth := ctx.GetAuthentication()
	if auth == nil {
		return nil
	}

	return auth.GetAuthorities()
}
