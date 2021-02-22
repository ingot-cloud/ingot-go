package ingot

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
)

type (
	authKey struct{}
)

// SetAuthentication 设置当前身份验证信息
func SetAuthentication(c *gin.Context, auth core.Authentication) {
	ctx := context.WithValue(c.Request.Context(), authKey{}, auth)
	c.Request = c.Request.WithContext(ctx)
}

// GetAuthentication 获取身份验证信息
func GetAuthentication(c *gin.Context) core.Authentication {
	v := c.Request.Context().Value(authKey{})
	if v != nil {
		if auth, ok := v.(core.Authentication); ok {
			return auth
		}
	}
	return nil
}
