package configurers

import "github.com/ingot-cloud/ingot-go/pkg/framework/security"

// BasicAuthentication basic 验证
type BasicAuthentication struct {
}

// Configure 配置
func (b *BasicAuthentication) Configure(http security.HTTPSecurityBuilder) error {

	return nil
}
