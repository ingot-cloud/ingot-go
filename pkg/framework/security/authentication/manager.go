package authentication

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// Manager 身份验证管理器
type Manager interface {
	// 对 Authentication 进行身份验证，验证成功后返回完全填充的Authentication
	Authenticate(core.Authentication) (core.Authentication, error)
}
