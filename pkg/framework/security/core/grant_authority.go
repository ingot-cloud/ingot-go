package core

// GrantedAuthority 授予身份验证对象的权限
type GrantedAuthority interface {
	// 获取授予的权限
	GetAuthority() string
}
