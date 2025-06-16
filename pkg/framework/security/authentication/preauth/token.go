package preauth

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
)

// AuthenticationToken 预身份验证令牌
type AuthenticationToken struct {
	*authentication.AbstractAuthenticationToken
	Principal   any
	Credentials string
}

// NewAuthenticationToken 创建预验证令牌
func NewAuthenticationToken(principal any, credentials string, authorities []core.GrantedAuthority) *AuthenticationToken {
	auth := &AuthenticationToken{
		Principal:   principal,
		Credentials: credentials,
	}

	if authorities != nil {
		auth.SetAuthenticated(true)
	}
	auth.AbstractAuthenticationToken = authentication.NewAbstractAuthenticationToken(authorities)

	return auth
}

// GetCredentials 凭证信息
func (token *AuthenticationToken) GetCredentials() string {
	return token.Credentials
}

// GetPrincipal 身份验证的主体
func (token *AuthenticationToken) GetPrincipal() any {
	return token.Principal
}

// EraseCredentials 擦除敏感数据
func (token *AuthenticationToken) EraseCredentials() {
	token.EraseSecret(token.GetPrincipal())
	token.EraseSecret(token.GetDetails())
	token.Credentials = ""
}
