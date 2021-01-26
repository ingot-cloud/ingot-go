package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// WebSecurity web 配置
type WebSecurity struct {
	securityFilterChainBuilders []HTTPSecurityBuilder
	ignoredRequests             []utils.RequestMatcher
}

// Build 构建Web过滤器
func (w *WebSecurity) Build() (filter.Filter, error) {
	chainSize := len(w.ignoredRequests) + len(w.securityFilterChainBuilders)
	securityFilterChains := make([]filter.SecurityFilterChain, 0, chainSize)

	// 忽略的请求
	for _, ignoredRequest := range w.ignoredRequests {
		securityFilterChains = append(securityFilterChains, &filter.DefaultSecurityFilterChain{
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

	filterChainProxy := &filter.ChainProxy{
		FilterChains: securityFilterChains,
	}
	return filterChainProxy, nil
}

// AddSecurityFilterChainBuilder 添加创建 SecurityFilterChain 的构建器
func (w *WebSecurity) AddSecurityFilterChainBuilder(builder HTTPSecurityBuilder) {
	w.securityFilterChainBuilders = append(w.securityFilterChainBuilders, builder)
}

// AddIgnoreRequestMatcher 添加忽略的请求匹配器
func (w *WebSecurity) AddIgnoreRequestMatcher(matcher utils.RequestMatcher) {
	w.ignoredRequests = append(w.ignoredRequests, matcher)
}
