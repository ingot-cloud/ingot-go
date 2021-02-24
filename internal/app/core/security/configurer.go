package security

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/web/container"
)

// Build 构建安全配置
func Build() (security.WebSecurityConfigurers, error) {
	configurers := container.WebSecurityConfigurers()
	// 可以追加自定义 WebSecurityConfigurer
	// todo 追加先后问题
	return configurers, nil
}
