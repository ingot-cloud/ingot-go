package userdetails

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// User UserDetails 的基本实现
type User struct {
	Authorities           []core.GrantedAuthority
	Username              string
	Password              string
	AccountNonExpired     bool
	AccountNonLocked      bool
	CredentialsNonExpired bool
	Enabled               bool
}

// NewUser 实例化
func NewUser(username, password string, authorities []core.GrantedAuthority) *User {
	return NewUserAllParams(username, password, authorities, true, true, true, true)
}

// NewUserAllParams 实例化
func NewUserAllParams(username, password string, authorities []core.GrantedAuthority, accountNonExpired, accountNonLocked, credentialsNonExpired, enabled bool) *User {
	return &User{
		Username:              username,
		Password:              password,
		Authorities:           authorities,
		AccountNonExpired:     accountNonExpired,
		AccountNonLocked:      accountNonLocked,
		CredentialsNonExpired: credentialsNonExpired,
		Enabled:               enabled,
	}
}

// GetAuthorities 获取授予用户的权限
func (u *User) GetAuthorities() []core.GrantedAuthority {
	return u.Authorities
}

// GetUsername 获取用户名
func (u *User) GetUsername() string {
	return u.Username
}

// GetPassword 获取密码
func (u *User) GetPassword() string {
	return u.Password
}

// IsAccountNonExpired 账户是否未过期
func (u *User) IsAccountNonExpired() bool {
	return u.AccountNonExpired
}

// IsAccountNonLocked 账户是否未锁定
func (u *User) IsAccountNonLocked() bool {
	return u.AccountNonLocked
}

// IsCredentialsNonExpired 用户凭证（密码）是否未过期
func (u *User) IsCredentialsNonExpired() bool {
	return u.CredentialsNonExpired
}

// IsEnabled 用户是否可用
func (u *User) IsEnabled() bool {
	return u.Enabled
}
