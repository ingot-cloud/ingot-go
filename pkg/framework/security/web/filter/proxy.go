package filter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// ChainProxy 过滤器链代理
type ChainProxy struct {
	FilterChains []SecurityFilterChain
}

// Name 名字
func (p *ChainProxy) Name() string {
	return "ChainProxy"
}

// Order 排序
func (p *ChainProxy) Order() int {
	return 0
}

// DoFilter 执行过滤器
func (p *ChainProxy) DoFilter(context *ingot.Context, chain filter.Chain) error {
	filters := p.GetFilters(context)
	if len(filters) == 0 {
		log.Infof("No security filter for %s", context.Request.RequestURI)
		return chain.DoFilter(context)
	}

	virtualFilterChain := &virtualFilterChain{
		context:           context,
		originalChain:     chain,
		additionalFilters: filters,
		currentPosition:   0,
		size:              len(filters),
	}
	return virtualFilterChain.DoFilter(context)
}

// GetFilters 获取指定请求需要执行的过滤器列表
func (p *ChainProxy) GetFilters(context *ingot.Context) filter.Filters {
	for _, chain := range p.FilterChains {
		if chain.Matches(context) {
			return chain.GetFilters()
		}
	}
	return nil
}

// 内部虚拟过滤器链
type virtualFilterChain struct {
	context           *ingot.Context
	originalChain     filter.Chain
	additionalFilters filter.Filters
	currentPosition   int
	size              int
}

func (c *virtualFilterChain) DoFilter(context *ingot.Context) error {
	if c.currentPosition == c.size {
		return c.originalChain.DoFilter(context)
	}
	nextFilter := c.additionalFilters[c.currentPosition]
	c.currentPosition++
	log.Infof("Do filter [%s] for url: %s", nextFilter.Name(), context.Request.RequestURI)

	return nextFilter.DoFilter(context, c)
}
