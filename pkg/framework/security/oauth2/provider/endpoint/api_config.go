package endpoint

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
)

// OAuth2ApiConfig http 配置
type OAuth2ApiConfig struct {
	OAuth2Api *OAuth2Api
}

// NewOAuth2ApiConfig 实例化
func NewOAuth2ApiConfig(token *TokenEndpoint) *OAuth2ApiConfig {
	return &OAuth2ApiConfig{
		OAuth2Api: &OAuth2Api{
			TokenEndpoint: token,
		},
	}
}

// Configure 应用配置
func (c *OAuth2ApiConfig) Configure(app *ingot.Router) {
}

// GetAPI 获取API
func (c *OAuth2ApiConfig) GetAPI() api.Configurers {
	return api.Configurers{c.OAuth2Api}
}
