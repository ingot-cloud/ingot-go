package authentication

import (
	"fmt"

	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/errors"
)

// ProviderManager Provider管理器
type ProviderManager struct {
	providers Providers
}

// NewProviderManager 实例化
func NewProviderManager(providers Providers) *ProviderManager {
	return &ProviderManager{
		providers: providers,
	}
}

func (*ProviderManager) Authorization() {}

// Authenticate 对 Authentication 进行身份验证，验证成功后返回完全填充的Authentication
func (m *ProviderManager) Authenticate(auth core.Authentication) (core.Authentication, error) {

	var result core.Authentication
	var err error = nil
	for _, provider := range m.providers.Get() {
		if !provider.Supports(auth) {
			continue
		}

		result, err = provider.Authenticate(auth)
		if err != nil {
			return nil, err
		}
	}

	if result == nil {
		return nil, errors.ProviderNotFound("No AuthenticationProvider found for ", fmt.Sprintf("%v", auth))
	}

	return result, nil
}
