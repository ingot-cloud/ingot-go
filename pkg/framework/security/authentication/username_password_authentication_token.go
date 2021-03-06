package authentication

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// NewUnauthenticatedUsernamePasswordAuthToken 获取未验证的token
func NewUnauthenticatedUsernamePasswordAuthToken(principal interface{}, credentials string) *UsernamePasswordAuthenticationToken {
	token := &UsernamePasswordAuthenticationToken{
		Principal:                   principal,
		Credentials:                 credentials,
		AbstractAuthenticationToken: &AbstractAuthenticationToken{},
	}
	token.SetAuthenticated(false)
	return token
}

// NewAuthenticatedUsernamePasswordAuthToken 获取验证的token
func NewAuthenticatedUsernamePasswordAuthToken(principal interface{}, credentials string, authorities []core.GrantedAuthority) *UsernamePasswordAuthenticationToken {
	token := &UsernamePasswordAuthenticationToken{
		Principal:   principal,
		Credentials: credentials,
		AbstractAuthenticationToken: &AbstractAuthenticationToken{
			Authorities: authorities,
		},
	}
	token.SetAuthenticated(true)
	return token
}

// UsernamePasswordAuthenticationToken 用户密码身份验证令牌
type UsernamePasswordAuthenticationToken struct {
	Principal   interface{}
	Credentials string
	*AbstractAuthenticationToken
}

// GetCredentials 凭证信息
func (token *UsernamePasswordAuthenticationToken) GetCredentials() string {
	return token.Credentials
}

// GetPrincipal 身份验证的主体
func (token *UsernamePasswordAuthenticationToken) GetPrincipal() interface{} {
	return token.Principal
}

// EraseCredentials 擦除敏感数据
func (token *UsernamePasswordAuthenticationToken) EraseCredentials() {
	token.AbstractAuthenticationToken.EraseCredentials()
	token.Credentials = ""
}
