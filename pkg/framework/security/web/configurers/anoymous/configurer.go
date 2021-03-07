package anonymous

import "github.com/ingot-cloud/ingot-go/pkg/framework/security"

// SecurityConfigurer basic 验证
type SecurityConfigurer struct {
}

// NewSecurityConfigurer 配置
func NewSecurityConfigurer() *SecurityConfigurer {
	return &SecurityConfigurer{}
}

// HTTPConfigure 配置
func (b *SecurityConfigurer) HTTPConfigure(http security.HTTPSecurityBuilder) error {
	http.AddFilter(NewFilter())
	return nil
}
