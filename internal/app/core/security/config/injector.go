package config

import (
	securityService "github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
	appToken "github.com/ingot-cloud/ingot-go/internal/app/core/security/token"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/internal/app/service"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/utils"
)

// IngotContainerInjector 安全注入
type IngotContainerInjector struct {
	*container.DefaultContainerInjector

	// 此处注入的实例可以通过GetValue方法获取
	OauthClientDetailsDao *dao.OauthClientDetails
	UserDetailService     service.UserDetail
	Ignore                utils.RequestMatcher

	// inject 代表自定义替换默认容器中的实例
	ClientDetailsService             *securityService.ClientDetails             `inject:"true"`
	UserDetailsService               *securityService.UserDetails               `inject:"true"`
	ResourceServerAdapter            *ResourceServerAdapter                     `inject:"true"`
	IngotEnhancerChain               *appToken.IngotEnhancerChain               `inject:"true"`
	IngotUserAuthenticationConverter *appToken.IngotUserAuthenticationConverter `inject:"true"`
}
