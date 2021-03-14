package provider

import (
	"github.com/google/wire"
	appConfig "github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils/pathmatcher"
	securityAuth "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/configurer"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// SecurityInjector 注入器
var SecurityInjector = wire.NewSet(
	wire.Struct(new(config.IngotSecurityInjector), "*"),
	wire.Bind(new(securityContainer.SecurityInjector), new(*config.IngotSecurityInjector)),

	SecurityClientDetailsService,
	SecurityUserDetailsService,
	ResourceServerAdapter,
	PermitURLMatcher,
	IngotEnhancerChain,
	IngotUserAuthenticationConverter,
)

// SecurityClientDetailsService 服务实现
var SecurityClientDetailsService = wire.Struct(new(service.ClientDetails), "*")

// SecurityUserDetailsService 服务实现
var SecurityUserDetailsService = wire.Struct(new(service.UserDetails), "*")

// IngotUserAuthenticationConverter 自定义
var IngotUserAuthenticationConverter = wire.Struct(new(token.IngotUserAuthenticationConverter), "*")

// IngotEnhancerChain token 增强
func IngotEnhancerChain(jwt *store.JwtAccessTokenConverter) *token.IngotEnhancerChain {
	return token.NewIngotEnhancerChain(jwt)
}

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
