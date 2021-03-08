package filter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// Tenant 过滤器
type Tenant struct {
}

// NewTenantFilter 实例化
func NewTenantFilter() *Tenant {
	return &Tenant{}
}

// Name 名字
func (b *Tenant) Name() string {
	return "TenantFilter"
}

// Order 过滤器排序
func (b *Tenant) Order() int {
	return constants.OrderFilterBasic - 10
}

// DoFilter 执行过滤器
func (b *Tenant) DoFilter(context *ingot.Context, chain filter.Chain) error {

	// todo
	log.Infof("测试 TenantFilter")

	return chain.DoFilter(context)
}
