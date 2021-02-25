package cache

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"

// NilUserCache 空实现
type NilUserCache struct {
}

// NewNilUserCache 实例化
func NewNilUserCache() *NilUserCache {
	return &NilUserCache{}
}

// GetUserFromCache 从缓存中获取用户信息
func (*NilUserCache) GetUserFromCache(username string) (userdetails.UserDetails, error) {
	return nil, nil
}

// PutUserInCache 增加缓存
func (*NilUserCache) PutUserInCache(user userdetails.UserDetails) error {
	return nil
}

// RemoveUserFromCache 移除缓存
func (*NilUserCache) RemoveUserFromCache(username string) error {
	return nil
}
