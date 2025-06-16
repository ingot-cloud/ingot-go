package factory

import "github.com/ingot-cloud/ingot-go/internal/app/model/domain"

// 如果需要自动创建domain中相应的表，需要再次增加domain
func getDomain() []any {
	return []any{
		new(domain.SysUser),
	}
}
