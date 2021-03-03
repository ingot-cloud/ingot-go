package basic

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/filter"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
)

// Filter basic token 验证
type Filter struct {
	BasicAuthenticationConverter *AuthenticationConverter
	AuthenticationManager        authentication.Manager
}

// NewFilter 实例化
func NewFilter(manager authentication.Manager) *Filter {
	return &Filter{
		BasicAuthenticationConverter: NewAuthenticationConverter(),
		AuthenticationManager:        manager,
	}
}

// Name 名字
func (b *Filter) Name() string {
	return "BasicAuthenticationFilter"
}

// Order 过滤器排序
func (b *Filter) Order() int {
	return constants.OrderFilterBasic
}

// DoFilter 执行过滤器
func (b *Filter) DoFilter(context *ingot.Context, chain filter.Chain) error {
	auth, err := b.BasicAuthenticationConverter.Converter(context)
	if err != nil {
		return err
	}
	if auth == nil {
		return chain.DoFilter(context)
	}

	username := auth.GetName(auth)
	if b.authenticationIsRequired(context, username) {
		authResult, err := b.AuthenticationManager.Authenticate(auth)
		if err != nil {
			return err
		}
		context.SetAuthentication(authResult)
	}

	return chain.DoFilter(context)
}

func (b *Filter) authenticationIsRequired(ctx *ingot.Context, username string) bool {
	existingAuth := ctx.GetAuthentication()
	if existingAuth == nil || !existingAuth.IsAuthenticated() {
		return true
	}

	if temp, ok := existingAuth.(*authentication.UsernamePasswordAuthenticationToken); ok {
		if temp.GetName(temp) != username {
			return true
		}
	}

	_, ok := existingAuth.(*authentication.AnonymousAuthenticationToken)
	return ok
}
