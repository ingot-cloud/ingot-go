package provider

import (
	"github.com/google/wire"
	appConfig "github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils/pathmatcher"
	securityAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// CustomSecurityAll 自定义
var CustomSecurityAll = wire.NewSet(
	SecurityClientDetailsService,
	SecurityUserDetailsService,
	ResourceServerAdapter,
	PermitURLMatcher,
	IngotTokenEnhancer,
)

// SecurityClientDetailsService 服务实现
var SecurityClientDetailsService = wire.Struct(new(service.ClientDetails), "*")

// SecurityUserDetailsService 服务实现
var SecurityUserDetailsService = wire.Struct(new(service.UserDetails), "*")

// IngotTokenEnhancer token增强
var IngotTokenEnhancer = wire.Struct(new(token.IngotEnhancer), "*")

// ResourceServerAdapter 自定义适配器
func ResourceServerAdapter(tokenExtractor authentication.TokenExtractor, resourceManager securityAuth.ResourceManager, ignore utils.RequestMatcher) *config.ResourceServerAdapter {
	parent := configurer.NewResourceServerConfigurer(tokenExtractor, resourceManager)
	return config.NewResourceServerAdapter(parent, ignore)
}

// PermitURLMatcher 忽略请求匹配器
func PermitURLMatcher(securityConfig appConfig.Security) utils.RequestMatcher {
	pathMatcher := pathmatcher.NewAntPathMatcher()
	return func(ctx *ingot.Context) bool {
		permitURLs := securityConfig.PermitURLs
		requestURL := ctx.Request.RequestURI

		for _, pattern := range permitURLs {
			if !pathMatcher.Match(pattern, requestURL) {
				return false
			}
		}

		return true
	}
}
