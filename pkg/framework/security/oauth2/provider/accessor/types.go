package accessor

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// SecurityContextAccessor 访问指定安全上下文有关有用信息的策略。
type SecurityContextAccessor interface {
	IsUser(*ingot.Context) bool
	GetAuthorities(*ingot.Context) []core.GrantedAuthority
}
