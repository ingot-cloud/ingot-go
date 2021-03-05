package token

import "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"

// EnhancerChain token 增强链
type EnhancerChain struct {
	delegates []Enhancer
}

// NewEnhancerChain 实例化
func NewEnhancerChain() *EnhancerChain {
	return &EnhancerChain{}
}

// SetTokenEnhancers 设置token增强列表
func (c *EnhancerChain) SetTokenEnhancers(delegates []Enhancer) {
	c.delegates = delegates
}

// Enhance 增强
func (c *EnhancerChain) Enhance(accessToken OAuth2AccessToken, auth *authentication.OAuth2Authentication) (OAuth2AccessToken, error) {
	result := accessToken
	var err error = nil
	for _, enhancer := range c.delegates {
		result, err = enhancer.Enhance(result, auth)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
