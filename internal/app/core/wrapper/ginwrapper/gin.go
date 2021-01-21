package ginwrapper

import (
	"fmt"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security"
	"github.com/ingot-cloud/ingot-go/internal/app/support/errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	prefix         = "magician"
	keyContextUser = prefix + ":User"
)

// GetToken 获取Token
func GetToken(ctx *gin.Context) string {
	var token string
	auth := ctx.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetUser 获取User
func GetUser(ctx *gin.Context) (*security.User, bool) {
	user, ok := ctx.Get(keyContextUser)
	if !ok {
		return nil, false
	}

	user1, ok1 := (user).(*security.User)
	return user1, ok1
}

// SetUser 设置 user
func SetUser(ctx *gin.Context, user *security.User) {
	ctx.Set(keyContextUser, user)
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
