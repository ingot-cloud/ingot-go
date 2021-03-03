package authentication

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// AnonymousAuthenticationToken 匿名认证token
type AnonymousAuthenticationToken struct {
	Principal interface{}
	*AbstractAuthenticationToken
}

// NewAnonymousAuthenticationToken 实例化
func NewAnonymousAuthenticationToken(principal interface{}, authorities []core.GrantedAuthority) *AnonymousAuthenticationToken {
	token := &AnonymousAuthenticationToken{
		Principal: principal,
		AbstractAuthenticationToken: &AbstractAuthenticationToken{
			Authorities: authorities,
		},
	}
	token.SetAuthenticated(true)
	return token
}

// GetPrincipal 身份验证的主体
func (token *AnonymousAuthenticationToken) GetPrincipal() interface{} {
	return token.Principal
}
