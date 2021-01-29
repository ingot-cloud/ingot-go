package authentication

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
)

// OAuth2AuthenticationManager OAuth2 身份验证管理器
type OAuth2AuthenticationManager struct {
}

// Authenticate 对 Authentication 进行身份验证，验证成功后返回完全填充的Authentication
func (manager *OAuth2AuthenticationManager) Authenticate(auth core.Authentication) (core.Authentication, error) {
	if auth == nil {
		return nil, errors.InvalidToken("Invalid token (token not found)")
	}

	token, ok := auth.GetPrincipal().(string)
	if !ok {
		return nil, errors.ErrInvalidToken
	}

	// token service

	return auth, nil
}
