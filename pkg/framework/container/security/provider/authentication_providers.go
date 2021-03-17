package provider

import (
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/container/security"
	coreAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/basic"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
)

// ProvidersImpl 接口实现
type ProvidersImpl struct {
	providers []coreAuth.Provider

	Basic *basic.AuthenticationProvider
	Dao   *dao.AuthenticationProvider
}

// Add 追加provider
func (p *ProvidersImpl) Add(item coreAuth.Provider) {
	p.providers = append(p.providers, item)
}

// Get 获取所有Provider
func (p *ProvidersImpl) Get() []coreAuth.Provider {
	p.providers = append(p.providers, p.Basic)
	p.providers = append(p.providers, p.Dao)
	return p.providers
}

// DaoAuthenticationProvider UsernamePasswordAuthenticationToken 认证提供者
func DaoAuthenticationProvider(common *securityContainer.CommonContainer) *dao.AuthenticationProvider {
	return dao.NewProvider(common.PasswordEncoder, common.UserDetailsService, common.UserCache, common.PreChecker, common.PostChecker)
}

// BasicAuthenticationProvider 认证提供者，其中注入了 ClientDetailsUserDetailsService
func BasicAuthenticationProvider(common *securityContainer.CommonContainer) *basic.AuthenticationProvider {
	return basic.NewProvider(common.PasswordEncoder, common.ClientDetailsService, common.UserCache, common.PreChecker, common.PostChecker)
}
