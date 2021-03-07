package authresult

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/errors"
)

// Filter 过滤器
type Filter struct {
}

// NewFilter 实例化
func NewFilter() *Filter {
	return &Filter{}
}

// Name 名字
func (b *Filter) Name() string {
	return "AuthenticationResultFilter"
}

// Order 过滤器排序
func (b *Filter) Order() int {
	return constants.OrderFilterAuthenticationResult
}

// DoFilter 执行过滤器
func (b *Filter) DoFilter(context *ingot.Context, chain filter.Chain) error {

	currentAuth := context.GetAuthentication()
	_, isAnonymous := currentAuth.(*authentication.AnonymousAuthenticationToken)
	if currentAuth == nil || isAnonymous {
		log.Errorf("身份验证不充分(Full authentication is required to access this resource), url=%s", context.Request.RequestURI)
		return errors.InsufficientAuthentication("Full authentication is required to access this resource")
	}

	return chain.DoFilter(context)
}
