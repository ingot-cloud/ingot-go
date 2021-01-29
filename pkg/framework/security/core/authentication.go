package core

// Authentication 身份信息
type Authentication interface {
	// 授予 principal 的权限
	GetAuthorities() []GrantedAuthority
	// 凭证信息
	GetCredentials() string
	// 额外的身份验证请求信息
	GetDetails() interface{}
	// 身份验证的主体
	GetPrincipal() interface{}
	// 是否已经通过身份验证
	IsAuthenticated() bool
	// 主动设置是否通过身份验证
	SetAuthenticated(bool)
	// 获取当前主题的名称
	GetName(Authentication) string
}
