package request

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils/maputil"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// TokenRequest TokenEndpoint发出的OAuth2令牌请求，
// RequestParameters包含原始OAuth2请求中未经修改的原始参数
type TokenRequest struct {
	*BaseRequestField
	GrantType string
}

// NewTokenRequest 创建TokenRequest
func NewTokenRequest(params map[string]string, clientID string, scope []string, grantType string) *TokenRequest {
	return &TokenRequest{
		BaseRequestField: &BaseRequestField{
			ClientID:          clientID,
			Scope:             scope,
			RequestParameters: params,
		},
		GrantType: grantType,
	}
}

// GetGrantType 获取授权类型
func (r *TokenRequest) GetGrantType() string {
	return r.GrantType
}

// CreateOAuth2Request 创建OAuth2Request
func (r *TokenRequest) CreateOAuth2Request(clientDetails clientdetails.ClientDetails) *OAuth2Request {
	requestParameters := r.GetRequestParameters()
	modifiable := maputil.CopyStringStringMap(requestParameters)

	// Remove password if present to prevent leaks
	delete(modifiable, constants.Password)
	delete(modifiable, constants.ClientSecret)
	// Add grant type so it can be retrieved from OAuth2Request
	modifiable[constants.GrantType] = r.GetGrantType()

	result := NewOAuth2Request(modifiable, clientDetails.GetClientID(), r.GetScope())
	result.Authorities = clientDetails.GetAuthorities()
	result.Approved = true
	result.ResourceIDs = clientDetails.GetResourceIDs()
	return result
}
