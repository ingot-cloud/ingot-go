package basic

import (
	"encoding/base64"
	"strings"

	ginwrapper "github.com/ingot-cloud/ingot-go/pkg/framework/core/wrapper/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/errors"
)

// AuthenticationConverter 认证转换器
type AuthenticationConverter struct {
}

// NewAuthenticationConverter 实例化
func NewAuthenticationConverter() *AuthenticationConverter {
	return &AuthenticationConverter{}
}

// Converter 转换
func (c *AuthenticationConverter) Converter(ctx *ingot.Context) (*authentication.UsernamePasswordAuthenticationToken, error) {
	if !ginwrapper.IsBasicAuth(ctx.Context) {
		return nil, nil
	}
	base64Token := ginwrapper.GetBasicToken(ctx.Context)
	if base64Token == "" {
		return nil, errors.BadCredentials("Empty basic authentication token")
	}

	raw, err := base64.StdEncoding.DecodeString(base64Token)
	if err != nil {
		return nil, errors.BadCredentials("Failed to decode basic authentication token")
	}
	rawString := string(raw)
	if strings.Index(rawString, ":") == -1 {
		return nil, errors.BadCredentials("Invalid basic authentication token")
	}
	token := strings.Split(rawString, ":")

	result := authentication.NewUnauthenticatedUsernamePasswordAuthToken(token[0], token[1])
	return result, nil
}
