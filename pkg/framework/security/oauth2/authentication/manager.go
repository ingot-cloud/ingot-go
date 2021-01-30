package authentication

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// OAuth2AuthenticationManager OAuth2 身份验证管理器
type OAuth2AuthenticationManager struct {
	ResourceServerTokenServices token.ResourceServerTokenServices
	ClientDetailsService        clientdetails.Service
	ResourceID                  string
}

// Authenticate 对 Authentication 进行身份验证，验证成功后返回完全填充的Authentication
func (manager *OAuth2AuthenticationManager) Authenticate(auth core.Authentication) (core.Authentication, error) {
	if auth == nil {
		return nil, errors.InvalidToken("Invalid token (token not found)")
	}

	token, ok := auth.GetPrincipal().(string)
	if !ok {
		return nil, errors.ErrInvalidToken
	}

	oauth2Auth, err := manager.ResourceServerTokenServices.LoadAuthentication(token)
	if err != nil {
		return nil, err
	}

	resourceIDs := oauth2Auth.GetOAuth2Request().GetResourceIDs()
	if manager.ResourceID != "" && len(resourceIDs) != 0 {
		var contains bool
		for _, rid := range resourceIDs {
			if rid == manager.ResourceID {
				contains = true
				break
			}
		}
		if !contains {
			msg := utils.StringCombine("Invalid token does not contain resource id (", manager.ResourceID, ")")
			return nil, errors.Forbidden(msg)
		}
	}

	err = manager.checkClientDetails(oauth2Auth)
	if err != nil {
		return nil, err
	}

	auth.SetAuthenticated(true)
	return auth, nil
}

func (manager *OAuth2AuthenticationManager) checkClientDetails(auth *authentication.OAuth2Authentication) error {
	if manager.ClientDetailsService != nil {
		client, err := manager.ClientDetailsService.LoadClientByClientId(auth.GetOAuth2Request().ClientID)
		if err != nil {
			return err
		}
		allowed := client.GetScope()
		requestScope := auth.GetOAuth2Request().GetScope()
		for _, scope := range requestScope {
			for _, allow := range allowed {
				if allow != scope {
					msg := utils.StringCombine("Invalid token contains disallowed scope (", scope, ") for this client")
					return errors.Forbidden(msg)
				}
			}
		}
	}
	return nil
}
