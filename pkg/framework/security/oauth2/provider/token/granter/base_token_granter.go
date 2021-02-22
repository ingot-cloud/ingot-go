package granter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/request"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// BaseTokenGranter 基础功能
type BaseTokenGranter struct {
	tokenServices token.AuthorizationServerTokenServices
}

// GetAccessToken 获取 AccessToken
func (g *BaseTokenGranter) GetAccessToken(client clientdetails.ClientDetails, tokenRequest request.TokenRequest) (token.OAuth2AccessToken, error) {
	storedOAuth2Request := tokenRequest.CreateOAuth2Request(client)

	auth := authentication.NewOAuth2Authentication(storedOAuth2Request, nil)
	return g.tokenServices.CreateAccessToken(auth)
}

// ValidateGrantType 验证 grant type
func (g *BaseTokenGranter) ValidateGrantType(grantType string, client clientdetails.ClientDetails) error {
	authorizedGrantTypes := client.GetAuthorizedGrantTypes()
	if len(authorizedGrantTypes) == 0 {
		return errors.InvalidClient("Unauthorized grant type: ", grantType)
	}
	var contains bool
	for _, gt := range authorizedGrantTypes {
		if gt == grantType {
			contains = true
		}
	}
	if !contains {
		return errors.InvalidClient("Unauthorized grant type: ", grantType)
	}
	return nil
}
