package request

import (
	"strings"

	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
)

// OAuth2Request 保存OAuth2请求的相关信息
type OAuth2Request struct {
	*BaseRequestField
	ResourceIds   []string
	Authorities   []core.GrantedAuthority
	Approved      bool
	Refresh       *TokenRequest
	RedirectURI   string
	ResponseTypes []string
	Extensions    map[string]interface{}
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

// GetResourceIds 获取资源ID列表
func (r *OAuth2Request) GetResourceIds() []string {
	return r.ResourceIds
}

// GetExtensions 扩展信息
func (r *OAuth2Request) GetExtensions() map[string]interface{} {
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
