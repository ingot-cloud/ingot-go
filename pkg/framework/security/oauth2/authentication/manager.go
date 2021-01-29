package authentication

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// OAuth2AuthenticationManager OAuth2 身份验证管理器
type OAuth2AuthenticationManager struct {
}

// Authenticate 对 Authentication 进行身份验证，验证成功后返回完全填充的Authentication
func (manager *OAuth2AuthenticationManager) Authenticate(auth core.Authentication) (core.Authentication, error) {
	return auth, nil
}
