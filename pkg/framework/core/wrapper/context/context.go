package contextwrapper

import (
	"context"
)

type (
	transCtx     struct{}
	noTransCtx   struct{}
	transLockCtx struct{}
)

// WithTrans 创建事务的上下文
func WithTrans(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

// GetTrans 从上下文中获取事务
func GetTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

// WithNoTrans 创建不使用事务的上下文
func WithNoTrans(ctx context.Context) context.Context {
	return context.WithValue(ctx, noTransCtx{}, true)
}

// GetNoTrans 从上下文中获取不使用事务标识
func GetNoTrans(ctx context.Context) bool {
	v := ctx.Value(noTransCtx{})
	return v != nil && v.(bool)
}

// WithTransLock 创建事务锁的上下文
func WithTransLock(ctx context.Context) context.Context {
	return context.WithValue(ctx, transLockCtx{}, true)
}

// GetTransLock 从上下文中获取事务锁
func GetTransLock(ctx context.Context) bool {
	v := ctx.Value(transLockCtx{})
	return v != nil && v.(bool)
}
