package userdetails

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// User UserDetails 的基本实现
type User struct {
	Authorities           []core.GrantedAuthority
	Username              string
	Password              string
	accountNonExpired     bool
	accountNonLocked      bool
	CredentialsNonExpired bool
	Enabled               bool
}

// GetAuthorities 获取授予用户的权限
func (u *User) GetAuthorities() *[]core.GrantedAuthority {
	return &u.Authorities
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
	return u.accountNonExpired
}

// IsAccountNonLocked 账户是否未锁定
func (u *User) IsAccountNonLocked() bool {
	return u.accountNonLocked
}

// IsCredentialsNonExpired 用户凭证（密码）是否未过期
func (u *User) IsCredentialsNonExpired() bool {
	return u.CredentialsNonExpired
}

// IsEnabled 用户是否可用
func (u *User) IsEnabled() bool {
	return u.Enabled
}
