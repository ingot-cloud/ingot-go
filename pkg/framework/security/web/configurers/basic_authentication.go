package configurers

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// BasicAuthentication basic 验证
type BasicAuthentication struct {
}

// Configure 配置
func (b *BasicAuthentication) Configure(http security.HTTPSecurityBuilder) error {

	return nil
}

// BasicFilter basic token 验证
type BasicFilter struct {
}

// Name 名字
func (b *BasicFilter) Name() string {
	return "BasicFilter"
}

// Order 过滤器排序
func (b *BasicFilter) Order() int {
	return constants.OrderFilterBasic
}

// DoFilter 执行过滤器
func (b *BasicFilter) DoFilter(context *ingot.Context, chain filter.Chain) error {

	return chain.DoFilter(context)
}
