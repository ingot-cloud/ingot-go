package model

import (
	"regexp"
	"strings"

	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
)

// RequestParameters 请求参数
type RequestParameters struct {
	GrantType string `form:"grant_type"`
	Scope     string `form:"scope"`

	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`

	State        string `form:"state"`
	RedirectURI  string `form:"redirect_uri"`
	ResponseType string `form:"response_type"`
	Code         string `form:"code"`

	Username string `form:"username"`
	Password string `form:"password"`

	RefreshToken string `form:"refresh_token"`
}

// ToMap 转为 map
func (r RequestParameters) ToMap() map[string]string {
	result := make(map[string]string)

	result[constants.GrantType] = r.GrantType
	result[constants.Scope] = r.Scope

	result[constants.ClientID] = r.ClientID
	result[constants.ClientSecret] = r.ClientSecret

	result[constants.State] = r.State
	result[constants.RedirectURI] = r.RedirectURI
	result[constants.ResponseType] = r.ResponseType
	result[constants.Code] = r.Code

	result[constants.Username] = r.Username
	result[constants.Password] = r.Password

	result[constants.RefreshToken] = r.RefreshToken

	return result
}

// Scopes 解析 scope
func (r RequestParameters) Scopes() []string {
	if r.Scope != "" {
		scope := strings.TrimSpace(r.Scope)
		rule := regexp.MustCompile("[\\s+]")
		return rule.Split(scope, -1)
	}
	return nil
}
