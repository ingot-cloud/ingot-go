package model

import (
	"regexp"
	"strings"

	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
)

// RequestParameters 请求参数
type RequestParameters struct {
	ClientID     string `form:"client_id"`
	State        string `form:"state"`
	Scope        string `form:"scope"`
	RedirectURI  string `form:"redirect_uri"`
	ResponseType string `form:"response_type"`
	GrantType    string `form:"grant_type"`
	Code         string `form:"code"`
	RefreshToken string `form:"refresh_token"`
}

// ToMap 转为 map
func (r RequestParameters) ToMap() map[string]string {
	result := make(map[string]string)

	result[constants.ClientID] = r.ClientID
	result[constants.State] = r.State
	result[constants.Scope] = r.Scope
	result[constants.RedirectURI] = r.RedirectURI
	result[constants.ResponseType] = r.ResponseType
	result[constants.GrantType] = r.GrantType
	result[constants.Code] = r.Code
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
