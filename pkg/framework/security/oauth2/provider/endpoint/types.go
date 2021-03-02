package endpoint

import "github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"

// OAuth2HTTPConfigurer oauth配置
type OAuth2HTTPConfigurer interface {
	api.HTTPConfigurer
}
