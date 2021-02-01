package authentication

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/request"
)

// OAuth2Authentication oauth2 身份验证信息
type OAuth2Authentication struct {
	StoredRequest      *request.OAuth2Request
	UserAuthentication core.Authentication
	*authentication.AbstractAuthenticationToken
}

// NewOAuth2Authentication 创建OAuth2身份验证信息
// storedRequest 不能为空
func NewOAuth2Authentication(storedRequest *request.OAuth2Request, userAuthentication core.Authentication) *OAuth2Authentication {
	auth := &OAuth2Authentication{
		StoredRequest:      storedRequest,
		UserAuthentication: userAuthentication,
	}

	var authorities []core.GrantedAuthority
	if userAuthentication == nil {
		authorities = storedRequest.GetAuthorities()
	} else {
		authorities = userAuthentication.GetAuthorities()
	}
	auth.AbstractAuthenticationToken = authentication.NewAbstractAuthenticationToken(authorities)

	return auth
}

// GetCredentials 获取凭证
func (auth *OAuth2Authentication) GetCredentials() string {
	return ""
}

// GetPrincipal 获取身份验证信息主体
func (auth *OAuth2Authentication) GetPrincipal() interface{} {
	if auth.UserAuthentication == nil {
		return auth.StoredRequest.GetClientID()
	}
	return auth.UserAuthentication.GetPrincipal()
}

// IsClientOnly 是否有用户身份验证信息和该令牌关联，如果返回true则表示没有
func (auth *OAuth2Authentication) IsClientOnly() bool {
	return auth.UserAuthentication == nil
}

// GetOAuth2Request 获取OAuth2请求
func (auth *OAuth2Authentication) GetOAuth2Request() *request.OAuth2Request {
	return auth.StoredRequest
}

// IsAuthenticated 是否已经通过身份验证
func (auth *OAuth2Authentication) IsAuthenticated() bool {
	return auth.StoredRequest.IsApproved() && (auth.UserAuthentication == nil || auth.UserAuthentication.IsAuthenticated())
}

// EraseCredentials 擦除敏感数据
func (auth *OAuth2Authentication) EraseCredentials() {
	auth.AbstractAuthenticationToken.EraseCredentials()
	auth.EraseSecret(auth.UserAuthentication)
}
