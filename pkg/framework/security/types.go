package security

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	securityFilter "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// WebSecurityConfigurer Web security 配置
type WebSecurityConfigurer interface {
	WebConfigure(WebSecurityBuilder) error
}

// HTTPSecurityConfigurer HTTP security 配置
type HTTPSecurityConfigurer interface {
	HTTPConfigure(HTTPSecurityBuilder) error
}

// WebSecurityBuilder 构造器
type WebSecurityBuilder interface {
	Build() (filter.Filter, error)
	AddSecurityFilterChainBuilder(HTTPSecurityBuilder)
	AddIgnoreRequestMatcher(utils.RequestMatcher)
	Apply(WebSecurityConfigurer)
}

// HTTPSecurityBuilder 构造器
type HTTPSecurityBuilder interface {
	Build() (securityFilter.SecurityFilterChain, error)
	RequestMatcher(utils.RequestMatcher)
	AddFilter(filter.Filter)
	Apply(HTTPSecurityConfigurer)
}

// WebSecurityConfigurers 定义 Web Security 配置列表接口
type WebSecurityConfigurers interface {
	Add(WebSecurityConfigurer)
	// 获取所有 WebSecurityConfigurer
	Get() []WebSecurityConfigurer
}

// ResourceServerConfigurer 资源服务器配置
type ResourceServerConfigurer interface {
	WebSecurityConfigurer
}

// AuthorizationServerConfigurer 授权服务器配置
type AuthorizationServerConfigurer interface {
	WebSecurityConfigurer
}
