package contextwrapper

import (
	"context"
	"ingot/core/security"
)

type (
	userCtx struct{}
)

// WithUser 创建用户的上下文
func WithUser(ctx context.Context, user *security.User) context.Context {
	return context.WithValue(ctx, userCtx{}, user)
}

// GetUser 从上下文中获取用户
func GetUser(ctx context.Context) (*security.User, bool) {
	v := ctx.Value(userCtx{})
	if v != nil {
		if u, ok := v.(*security.User); ok {
			return u, ok
		}
	}
	return nil, false
}
