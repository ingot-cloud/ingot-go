package token

import (
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/user"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token/store"
)

type IngotEnhancerChain struct {
	*token.EnhancerChain
}

func NewIngotEnhancerChain(jwt *store.JwtAccessTokenConverter) *IngotEnhancerChain {
	chain := token.NewEnhancerChain()
	var enhancers []token.Enhancer
	enhancers = append(enhancers, &IngotEnhancer{})
	// 默认追加 jwt enhancer
	enhancers = append(enhancers, jwt)
	chain.SetTokenEnhancers(enhancers)

	return &IngotEnhancerChain{
		EnhancerChain: chain,
	}
}

// IngotEnhancer 自定义增强
type IngotEnhancer struct {
}

// Enhance 自定义增强
func (e *IngotEnhancer) Enhance(accessToken token.OAuth2AccessToken, authentication *authentication.OAuth2Authentication) (token.OAuth2AccessToken, error) {
	// client 授权模式直接跳过
	if authentication.GetOAuth2Request().GetGrantType() == constants.GrantTypeClient {
		return accessToken, nil
	}

	ingotUser, ok := authentication.GetPrincipal().(*user.IngotUser)
	if ok {
		additionalInfo := accessToken.GetAdditionalInformation()
		additionalInfo[EnhancerUserID] = ingotUser.ID.String()
		additionalInfo[EnhancerDeptID] = ingotUser.DeptID.String()
		additionalInfo[EnhancerTenantID] = ingotUser.TenantID.String()
		additionalInfo[EnhancerAuthType] = ingotUser.AuthType
	}

	return accessToken, nil
}
