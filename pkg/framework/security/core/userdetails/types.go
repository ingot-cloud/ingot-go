package userdetails

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// UserDetails 用户信息
type UserDetails interface {
	// 获取授予用户的权限
	GetAuthorities() []core.GrantedAuthority
	// 获取用户名
	GetUsername() string
	// 获取密码
	GetPassword() string
	// 账户是否未过期
	IsAccountNonExpired() bool
	// 账户是否未锁定
	IsAccountNonLocked() bool
	// 用户凭证（密码）是否未过期
	IsCredentialsNonExpired() bool
	// 用户是否可用
	IsEnabled() bool
}

// Service 用于加载 UserDetails
type Service interface {
	// 加载指定 username 的用户
	LoadUserByUsername(username string) (UserDetails, error)
}

// Checker 检查加载 UserDetails 的状态
type Checker interface {
	// 检测用户状态
	Check(user UserDetails) error
}

// PreChecker 前置检查器
type PreChecker interface {
	Checker
}

// PostChecker 后置检查器
type PostChecker interface {
	Checker
}

// UserCache 用户缓存接口
type UserCache interface {
	// 从缓存中获取用户信息
	GetUserFromCache(string) (UserDetails, error)
	// 增加缓存
	PutUserInCache(UserDetails) error
	// 移除缓存
	RemoveUserFromCache(string) error
}
