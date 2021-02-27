package security

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
)

// Build 构建安全配置
func Build() (security.WebSecurityConfigurers, error) {
	// 可以追加自定义 WebSecurityConfigurer
	// todo 追加先后问题
	return nil, nil
}
