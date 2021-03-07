package ginwrapper

import (
	"fmt"
	"strings"

	"github.com/ingot-cloud/ingot-go/pkg/framework/core/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/model/enums"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	prefix         = "magician"
	keyContextUser = prefix + ":User"
)

// GetToken 获取Token
func GetToken(ctx *gin.Context, prefix string) string {
	var token string
	auth := ctx.GetHeader(enums.HeaderAuthentication)
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// IsBasicAuth 是否为 basic 认证
func IsBasicAuth(ctx *gin.Context) bool {
	auth := ctx.GetHeader(enums.HeaderAuthentication)
	if auth != "" {
		return strings.HasPrefix(auth, string(enums.BasicToken))
	}
	return false
}

// GetBasicToken 获取 basic token
func GetBasicToken(ctx *gin.Context) string {
	return GetToken(ctx, string(enums.BasicWithSpace))
}

// GetBearerToken 获取 bearer token
func GetBearerToken(ctx *gin.Context) string {
	return GetToken(ctx, string(enums.BearerWithSpace))
}

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.IllegalArgument(fmt.Sprintf("Error parsing request parameters - %s", err.Error()))
	}
	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.IllegalArgument(fmt.Sprintf("Error parsing request parameters - %s", err.Error()))
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.IllegalArgument(fmt.Sprintf("Error parsing request parameters - %s", err.Error()))
	}
	return nil
}
