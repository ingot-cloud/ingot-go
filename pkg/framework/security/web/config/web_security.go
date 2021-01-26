package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// WebSecurity web
type WebSecurity struct {
	securityFilterChainBuilders []HTTPSecurityBuilder
	ignoredRequests             []utils.RequestMatcher
}

// Build 构建Web过滤器
func (w *WebSecurity) Build() filter.Filter {
	chainSize := len(w.ignoredRequests) + len(w.securityFilterChainBuilders)
	securityFilterChains := make([]filter.SecurityFilterChain, 0, chainSize)

	for _, ignoredRequest := range w.ignoredRequests {
		securityFilterChains = append(securityFilterChains, &filter.DefaultSecurityFilterChain{
			RequestMatcher: ignoredRequest,
		})
	}
	for _, builder := range w.securityFilterChainBuilders {
		securityFilterChains = append(securityFilterChains, builder.Build())
	}

	filterChainProxy := &filter.ChainProxy{
		FilterChains: securityFilterChains,
	}
	return filterChainProxy
}

// AddSecurityFilterChainBuilder 添加创建 SecurityFilterChain 的构建器
func (w *WebSecurity) AddSecurityFilterChainBuilder(builder HTTPSecurityBuilder) {
	w.securityFilterChainBuilders = append(w.securityFilterChainBuilders, builder)
}

// AddIgnoreRequestMatcher 添加忽略的请求匹配器
func (w *WebSecurity) AddIgnoreRequestMatcher(matcher utils.RequestMatcher) {
	w.ignoredRequests = append(w.ignoredRequests, matcher)
}
