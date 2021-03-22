package authentication

import (
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

// NewOAuth2AuthenticationManager 实例化
func NewOAuth2AuthenticationManager(tokenService token.ResourceServerTokenServices) *OAuth2AuthenticationManager {
	return &OAuth2AuthenticationManager{
		ResourceServerTokenServices: tokenService,
	}
}

func (*OAuth2AuthenticationManager) Resource() {}

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
			return nil, errors.OAuth2AccessDenied("Invalid token does not contain resource id (", manager.ResourceID, ")")
		}
	}

	err = manager.checkClientDetails(oauth2Auth)
	if err != nil {
		return nil, err
	}

	oauth2Auth.SetDetails(auth.GetDetails())
	oauth2Auth.SetAuthenticated(true)
	return oauth2Auth, nil
}

func (manager *OAuth2AuthenticationManager) checkClientDetails(auth *authentication.OAuth2Authentication) error {
	if manager.ClientDetailsService != nil {
		client, err := manager.ClientDetailsService.LoadClientByClientID(auth.GetOAuth2Request().ClientID)
		if err != nil {
			return err
		}
		allowed := client.GetScope()
		requestScope := auth.GetOAuth2Request().GetScope()
		for _, scope := range requestScope {
			for _, allow := range allowed {
				if allow != scope {
					return errors.OAuth2AccessDenied("Invalid token contains disallowed scope (", scope, ") for this client")
				}
			}
		}
	}
	return nil
}
