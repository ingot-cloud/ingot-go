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
