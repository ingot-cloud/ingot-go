package authority

// SimpleGrantedAuthority 授权权限简单实现
type SimpleGrantedAuthority struct {
	Role string
}

// GetAuthority 获取权限
func (a *SimpleGrantedAuthority) GetAuthority() string {
	return a.Role
}
