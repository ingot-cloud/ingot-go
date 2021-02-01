package factory

import (
	"github.com/ingot-cloud/ingot-go/pkg/component/id"
	"github.com/ingot-cloud/ingot-go/pkg/component/id/snowflake"
)

// NewIDGenerator 注入 id 生成器
func NewIDGenerator() id.Generator {
	return &snowflake.Generator{
		WorkID: 1,
	}
}
