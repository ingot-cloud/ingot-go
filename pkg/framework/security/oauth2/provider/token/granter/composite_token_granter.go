package granter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/request"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// CompositeTokenGranter 组合 granter
type CompositeTokenGranter struct {
	tokenGranters []token.Granter
}

// NewCompositeTokenGranter 创建组合 granter
func NewCompositeTokenGranter() *CompositeTokenGranter {
	return &CompositeTokenGranter{}
}

// Grant 授权token
func (g *CompositeTokenGranter) Grant(grantType string, tokenRequest *request.TokenRequest) (token.OAuth2AccessToken, error) {
	for _, granter := range g.tokenGranters {
		token, err := granter.Grant(grantType, tokenRequest)
		if err != nil {
			return nil, err
		}
		if token != nil {
			return token, nil
		}
	}
	return nil, nil
}

// AddTokenGranter 添加 TokenGranter
func (g *CompositeTokenGranter) AddTokenGranter(granter token.Granter) {
	g.tokenGranters = append(g.tokenGranters, granter)
}
