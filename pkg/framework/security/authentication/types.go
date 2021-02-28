package authentication

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/core"

// Manager 身份验证管理器
type Manager interface {
	// 对 Authentication 进行身份验证，验证成功后返回完全填充的Authentication
	Authenticate(core.Authentication) (core.Authentication, error)
}

// ResourceManager 资源认证管理器
type ResourceManager interface {
	Manager
}

// AuthorizationManager 授权认证管理器
type AuthorizationManager interface {
	Manager
}

// Provider 身份验证提供者
type Provider interface {
	// 身份验证
	Authenticate(core.Authentication) (core.Authentication, error)
	// 该身份验证提供者是否支持指定的认证信息
	Supports(interface{}) bool
}

// Providers 提供者列表
type Providers []Provider
