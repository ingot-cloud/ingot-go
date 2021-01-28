package authentication

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/userdetails"
)

// AbstractAuthenticationToken 身份验证对象基本实现
type AbstractAuthenticationToken struct {
	Authorities   []core.GrantedAuthority
	details       interface{}
	authenticated bool
}

// GetAuthorities 授予 principal 的权限
func (token *AbstractAuthenticationToken) GetAuthorities() []core.GrantedAuthority {
	return token.Authorities
}

// GetCredentials 凭证信息
func (token *AbstractAuthenticationToken) GetCredentials() string {
	return ""
}

// GetDetails 额外的身份验证请求信息
func (token *AbstractAuthenticationToken) GetDetails() interface{} {
	return token.details
}

// GetPrincipal 身份验证的主体
func (token *AbstractAuthenticationToken) GetPrincipal() interface{} {
	return nil
}

// IsAuthenticated 是否已经通过身份验证
func (token *AbstractAuthenticationToken) IsAuthenticated() bool {
	return token.authenticated
}

// SetAuthenticated 主动设置是否通过身份验证
func (token *AbstractAuthenticationToken) SetAuthenticated(isAuthenticated bool) {
	token.authenticated = isAuthenticated
}

// GetName 获取当前主体名字
func (token *AbstractAuthenticationToken) GetName(target core.Authentication) string {
	object := target.GetPrincipal()
	if object == nil {
		return ""
	}
	switch value := object.(type) {
	case userdetails.UserDetails:
		return value.GetUsername()
	default:
		return ""
	}
}

// EraseCredentials 擦除敏感数据
func (token *AbstractAuthenticationToken) EraseCredentials() {
	token.EraseSecret(token.GetCredentials())
	token.EraseSecret(token.GetPrincipal())
	token.EraseSecret(token.details)
}

// SetDetails 设置额外的信息
func (token *AbstractAuthenticationToken) SetDetails(details interface{}) {
	token.details = details
}

// EraseSecret 擦除敏感信息
func (token *AbstractAuthenticationToken) EraseSecret(object interface{}) {
	if object == nil {
		return
	}
	value, ok := object.(core.CredentialsContainer)
	if ok {
		value.EraseCredentials()
	}
}
