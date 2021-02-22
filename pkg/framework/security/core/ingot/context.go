package ingot

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
)

// Context 自定义 context
type Context struct {
	*gin.Context
}

// NewContext 实例化
func NewContext(gin *gin.Context) *Context {
	return &Context{
		Context: gin,
	}
}

// SetAuthentication 设置当前身份验证信息
func (c *Context) SetAuthentication(auth core.Authentication) {
	SetAuthentication(c.Context, auth)
}

// GetAuthentication 获取身份验证信息
func (c *Context) GetAuthentication() core.Authentication {
	return GetAuthentication(c.Context)
}
