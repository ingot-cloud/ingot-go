package granter

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/utils/maputil"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
	oauth "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/request"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// PasswordTokenGranter 资源所有者密码凭据授予器
type PasswordTokenGranter struct {
	*BaseTokenGranter
	tokenServices         token.AuthorizationServerTokenServices
	authenticationManager authentication.Manager
}

// NewPasswordTokenGranter 实例化
func NewPasswordTokenGranter(tokenServices token.AuthorizationServerTokenServices, manager authentication.Manager) *PasswordTokenGranter {
	return &PasswordTokenGranter{
		BaseTokenGranter:      &BaseTokenGranter{},
		tokenServices:         tokenServices,
		authenticationManager: manager,
	}
}

// Grant 授予
func (g *PasswordTokenGranter) Grant(grantType string, client clientdetails.ClientDetails, tokenRequest *request.TokenRequest) (token.OAuth2AccessToken, error) {
	if grantType != constants.GrantTypePassword {
		return nil, nil
	}

	err := g.ValidateGrantType(grantType, client)
	if err != nil {
		return nil, err
	}

	return g.getAccessToken(client, tokenRequest)
}

func (g *PasswordTokenGranter) getAccessToken(client clientdetails.ClientDetails, tokenRequest *request.TokenRequest) (token.OAuth2AccessToken, error) {
	parameters := maputil.CopyStringStringMap(tokenRequest.GetRequestParameters())
	username := parameters[constants.Username]
	password := parameters[constants.Password]
	delete(parameters, constants.Password)

	userAuth := authentication.NewUnauthenticatedUsernamePasswordAuthToken(username, password)
	userAuth.SetDetails(parameters)

	postUserAuth, err := g.authenticationManager.Authenticate(userAuth)
	if err != nil {
		return nil, err
	}

	storedOAuth2Request := tokenRequest.CreateOAuth2Request(client)

	oauth2Auth := oauth.NewOAuth2Authentication(storedOAuth2Request, postUserAuth)
	return g.tokenServices.CreateAccessToken(oauth2Auth)
}
