package ginwrapper

import (
	"ingot/core/security"

	"github.com/gin-gonic/gin"
)

const (
	prefix         = "ingot"
	keyContextUser = prefix + ":User"
)

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
