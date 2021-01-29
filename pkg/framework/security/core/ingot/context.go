package ingot

import (
	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
)

// Context 自定义 context
type Context struct {
	*gin.Context
	auth core.Authentication
}

// SetAuthentication 设置当前身份验证信息
func (ctx *Context) SetAuthentication(auth core.Authentication) {
	ctx.auth = auth
}
