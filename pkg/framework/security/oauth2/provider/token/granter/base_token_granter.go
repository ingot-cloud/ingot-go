package granter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
)

// BaseTokenGranter 基础功能
type BaseTokenGranter struct {
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
