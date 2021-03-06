package builders

import (
	coreUtils "github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	securityFilter "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// WebSecurity web 配置
type WebSecurity struct {
	securityFilterChainBuilders []security.HTTPSecurityBuilder
	ignoredRequests             []utils.RequestMatcher
	webSecurityConfigurers      map[string]security.WebSecurityConfigurer
}

// NewWebSecurity 创建 WebSecurity
func NewWebSecurity() *WebSecurity {
	return &WebSecurity{
		webSecurityConfigurers: make(map[string]security.WebSecurityConfigurer),
	}
}

// Build 构建Web过滤器
func (w *WebSecurity) Build() (filter.Filter, error) {
	err := w.configure()
	if err != nil {
		return nil, err
	}
	return w.performBuild()
}

// AddSecurityFilterChainBuilder 添加创建 SecurityFilterChain 的构建器
func (w *WebSecurity) AddSecurityFilterChainBuilder(builder security.HTTPSecurityBuilder) {
	w.securityFilterChainBuilders = append(w.securityFilterChainBuilders, builder)
}

// AddIgnoreRequestMatcher 添加忽略的请求匹配器
func (w *WebSecurity) AddIgnoreRequestMatcher(matcher utils.RequestMatcher) {
	w.ignoredRequests = append(w.ignoredRequests, matcher)
}

// Apply 应用Web安全配置
func (w *WebSecurity) Apply(configurer security.WebSecurityConfigurer) {
	typeStr := coreUtils.GetType(configurer)
	log.Debugf("web security apply config = %s", typeStr)
	if _, ok := w.webSecurityConfigurers[typeStr]; !ok {
		w.webSecurityConfigurers[typeStr] = configurer
	}
}

func (w *WebSecurity) configure() error {
	// 执行配置
	for _, item := range w.webSecurityConfigurers {
		if err := item.WebConfigure(w); err != nil {
			return err
		}
	}
	return nil
}

func (w *WebSecurity) performBuild() (filter.Filter, error) {
	chainSize := len(w.ignoredRequests) + len(w.securityFilterChainBuilders)
	securityFilterChains := make([]securityFilter.SecurityFilterChain, 0, chainSize)

	// 忽略的请求
	for _, ignoredRequest := range w.ignoredRequests {
		securityFilterChains = append(securityFilterChains, &securityFilter.DefaultSecurityFilterChain{
			RequestMatcher: ignoredRequest,
		})
	}
	// 安全校验链
	for _, builder := range w.securityFilterChainBuilders {
		chain, err := builder.Build()
		if err != nil {
			return nil, err
		}
		securityFilterChains = append(securityFilterChains, chain)
	}

	filterChainProxy := &securityFilter.ChainProxy{
		FilterChains: securityFilterChains,
	}
	return filterChainProxy, nil
}
