package authentication

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// OAuth2ProcessingFilter OAuth2处理
type OAuth2ProcessingFilter struct {
	TokenExtractor        TokenExtractor
	AuthenticationManager authentication.Manager
}

// Order 过滤器排序
func (filter *OAuth2ProcessingFilter) Order() int {
	return constants.OrderFilterOAuth2
}

// DoFilter 执行过滤器
func (filter *OAuth2ProcessingFilter) DoFilter(context *ingot.Context, chain filter.Chain) error {

	auth := filter.TokenExtractor.Extract(context)
	if auth != nil {
		authResult, err := filter.AuthenticationManager.Authenticate(auth)
		if err != nil {
			return err
		}
		context.SetAuthentication(authResult)
	}

	return chain.DoFilter(context)
}
