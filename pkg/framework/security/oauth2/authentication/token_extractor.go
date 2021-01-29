package authentication

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// TokenExtractor token提取器
type TokenExtractor interface {
	// 从传入的请求中提取AccessToken并且不进行身份验证
	// 返回身份验证令牌
	Extract(*ingot.Context) core.Authentication
}
