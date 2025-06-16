package request

import (
	"strings"

	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
)

// OAuth2Request 保存OAuth2请求的相关信息
type OAuth2Request struct {
	*BaseRequestField
	ResourceIDs   []string
	Authorities   []core.GrantedAuthority
	Approved      bool
	Refresh       *TokenRequest
	RedirectURI   string
	ResponseTypes []string
	Extensions    map[string]any
}

// NewOAuth2Request 创建OAuth2Request
func NewOAuth2Request(params map[string]string, clientID string, scope []string) *OAuth2Request {
	return &OAuth2Request{
		BaseRequestField: &BaseRequestField{
			ClientID:          clientID,
			Scope:             scope,
			RequestParameters: params,
		},
		Extensions: make(map[string]any),
	}
}

// GetRedirectURI 获取重定向uri
func (r *OAuth2Request) GetRedirectURI() string {
	return r.RedirectURI
}

// GetResponseTypes 获取响应类型
func (r *OAuth2Request) GetResponseTypes() []string {
	return r.ResponseTypes
}

// GetAuthorities 获取权限
func (r *OAuth2Request) GetAuthorities() []core.GrantedAuthority {
	return r.Authorities
}

// IsApproved 请求是否被批准
func (r *OAuth2Request) IsApproved() bool {
	return r.Approved
}

// GetResourceIDs 获取资源ID列表
func (r *OAuth2Request) GetResourceIDs() []string {
	return r.ResourceIDs
}

// GetExtensions 扩展信息
func (r *OAuth2Request) GetExtensions() map[string]any {
	return r.Extensions
}

// IsRefresh 是否请求刷新令牌
func (r *OAuth2Request) IsRefresh() bool {
	return r.Refresh != nil
}

// GetRefreshTokenRequest 如果请求刷新令牌，那么不能为空，原始授权类型可以通过 GetGrantType 获取
func (r *OAuth2Request) GetRefreshTokenRequest() *TokenRequest {
	return r.Refresh
}

// GetGrantType 获取授权类型
func (r *OAuth2Request) GetGrantType() string {
	value, ok := r.GetRequestParameters()[constants.GrantType]
	if ok {
		return value
	}
	value, ok = r.GetRequestParameters()[constants.ResponseType]
	if ok && strings.Index(value, "token") != -1 {
		return "implicit"
	}
	return ""
}

// CreateOAuth2Request 创建OAuth2Request并且更新RequestParameters
func (r *OAuth2Request) CreateOAuth2Request(params map[string]string) *OAuth2Request {
	request := NewOAuth2Request(params, r.GetClientID(), r.GetScope())
	request.Authorities = r.Authorities
	request.Approved = r.Approved
	request.ResourceIDs = r.ResourceIDs
	request.RedirectURI = r.RedirectURI
	request.ResponseTypes = r.ResponseTypes
	request.Extensions = r.Extensions
	return request
}

// NarrowScope 创建OAuth2Request并且更新Scope
func (r *OAuth2Request) NarrowScope(scope []string) *OAuth2Request {
	request := NewOAuth2Request(r.GetRequestParameters(), r.GetClientID(), scope)
	request.Authorities = r.Authorities
	request.Approved = r.Approved
	request.ResourceIDs = r.ResourceIDs
	request.RedirectURI = r.RedirectURI
	request.ResponseTypes = r.ResponseTypes
	request.Extensions = r.Extensions

	request.Refresh = r.Refresh
	return request
}

// UpdateRefresh 创建OAuth2Request并且更新Refresh
func (r *OAuth2Request) UpdateRefresh(tokenRequest *TokenRequest) *OAuth2Request {
	request := NewOAuth2Request(r.GetRequestParameters(), r.GetClientID(), r.GetScope())
	request.Authorities = r.Authorities
	request.Approved = r.Approved
	request.ResourceIDs = r.ResourceIDs
	request.RedirectURI = r.RedirectURI
	request.ResponseTypes = r.ResponseTypes
	request.Extensions = r.Extensions

	request.Refresh = tokenRequest
	return request
}
