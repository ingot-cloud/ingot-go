package anonymous

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/authority"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// Filter basic token 验证
type Filter struct {
}

// NewFilter 实例化
func NewFilter() *Filter {
	return &Filter{}
}

// Name 名字
func (b *Filter) Name() string {
	return "AnonymousAuthenticationFilter"
}

// Order 过滤器排序
func (b *Filter) Order() int {
	return constants.OrderFilterAnonymous
}

// DoFilter 执行过滤器
func (b *Filter) DoFilter(context *ingot.Context, chain filter.Chain) error {

	currentAuth := context.GetAuthentication()
	if currentAuth == nil {
		context.SetAuthentication(b.newAuth())
	}

	return chain.DoFilter(context)
}

func (b *Filter) newAuth() core.Authentication {
	return authentication.NewAnonymousAuthenticationToken("anonymousUser", authority.CreateAuthorityList("ROLE_ANONYMOUS"))
}
