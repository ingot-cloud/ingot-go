package preauth

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"

// AuthenticationToken 预身份验证令牌
type AuthenticationToken struct {
	*authentication.AbstractAuthenticationToken
	Principal   interface{}
	Credentials string
}

// GetCredentials 凭证信息
func (token *AuthenticationToken) GetCredentials() string {
	return token.Credentials
}

// GetPrincipal 身份验证的主体
func (token *AuthenticationToken) GetPrincipal() interface{} {
	return token.Principal
}

// EraseCredentials 擦除敏感数据
func (token *AuthenticationToken) EraseCredentials() {
	token.EraseSecret(token.GetPrincipal())
	token.EraseSecret(token.GetDetails())
	token.Credentials = ""
}
