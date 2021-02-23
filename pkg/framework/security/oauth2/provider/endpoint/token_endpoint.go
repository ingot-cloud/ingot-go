package endpoint

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/core/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/constants"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/errors"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/model"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/clientdetails"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/request"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// TokenEndpoint token 端点
type TokenEndpoint struct {
	TokenGranter         token.Granter
	ClientDetailsService clientdetails.Service
}

// NewTokenEndpoint 实例
func NewTokenEndpoint(granter token.Granter, clientDetailsService clientdetails.Service) *TokenEndpoint {
	return &TokenEndpoint{
		TokenGranter:         granter,
		ClientDetailsService: clientDetailsService,
	}
}

// AccessToken /oauth/token
func (t *TokenEndpoint) AccessToken(ctx *gin.Context) (interface{}, error) {
	auth := ingot.GetAuthentication(ctx)
	if auth == nil {
		return nil, errors.InsufficientAuthentication("There is no client authentication. Try adding an appropriate authentication filter.")
	}
	var parameters model.RequestParameters
	if err := ctx.ShouldBindWith(&parameters, binding.Form); err != nil {
		return nil, errors.InsufficientAuthentication("Error parsing request parameters - ", err.Error())
	}

	clientID, err := t.getClientID(auth)
	if err != nil {
		return nil, err
	}

	authenticatedClient, err := t.ClientDetailsService.LoadClientByClientId(clientID)
	if err != nil {
		return nil, err
	}

	tokenRequest, err := t.createTokenRequest(parameters, authenticatedClient)
	if err != nil {
		return nil, err
	}

	if clientID != "" && clientID != tokenRequest.GetClientID() {
		return nil, errors.InvalidClient("Given client ID does not match authenticated client")
	}

	err = t.validateScope(tokenRequest.GetScope(), authenticatedClient.GetScope())
	if err != nil {
		return nil, err
	}

	if tokenRequest.GrantType == "" {
		return nil, errors.InvalidRequest("Missing grant type")
	}

	if tokenRequest.GrantType == constants.GrantTypeImplicit {
		return nil, errors.InvalidGrant("Implicit grant type not supported from token endpoint")
	}

	if t.isAuthCodeRequest(parameters) {
		// 如果是授权码模式，scope在授权的时候已经处理完成，这里不再需要
		if len(tokenRequest.GetScope()) != 0 {
			tokenRequest.Scope = nil
		}
	}

	if t.isRefreshTokenRequest(parameters) {
		// 如果是刷新Token模式，需要使用请求参数中的 scope
		tokenRequest.Scope = parameters.Scopes()
	}

	accessToken, err := t.TokenGranter.Grant(tokenRequest.GetGrantType(), tokenRequest)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (t *TokenEndpoint) getClientID(auth core.Authentication) (string, error) {
	if !auth.IsAuthenticated() {
		return "", errors.InsufficientAuthentication("The client is not authenticated.")
	}
	clientID := auth.GetName(auth)
	if oauth2Auth, ok := auth.(*authentication.OAuth2Authentication); ok {
		clientID = oauth2Auth.GetOAuth2Request().GetClientID()
	}
	return clientID, nil
}

func (t *TokenEndpoint) createTokenRequest(parameters model.RequestParameters, authenticatedClient clientdetails.ClientDetails) (*request.TokenRequest, error) {
	clientID := parameters.ClientID
	if clientID == "" {
		clientID = authenticatedClient.GetClientID()
	} else {
		if clientID != authenticatedClient.GetClientID() {
			return nil, errors.InvalidClient("Given client ID does not match authenticated client")
		}
	}

	grantType := parameters.GrantType
	scopes, err := t.extractScopes(parameters, authenticatedClient)
	if err != nil {
		return nil, err
	}
	tokenRequest := request.NewTokenRequest(parameters.ToMap(), clientID, scopes, grantType)
	return tokenRequest, nil
}

func (t *TokenEndpoint) extractScopes(parameters model.RequestParameters, authenticatedClient clientdetails.ClientDetails) ([]string, error) {
	scopes := parameters.Scopes()
	if len(scopes) == 0 {
		scopes = authenticatedClient.GetScope()
	}

	return scopes, nil
}

func (t *TokenEndpoint) validateScope(requestScopes []string, clientScopes []string) error {
	if len(clientScopes) != 0 {
		var contains bool
		for _, scope := range requestScopes {
			contains = false
			for _, clientScope := range clientScopes {
				contains = clientScope == scope
			}
			if !contains {
				return errors.InvalidScope("Invalid scope: ", scope, "; client scope: ", strings.Join(clientScopes, " "))
			}
		}
	}

	if len(requestScopes) == 0 {
		return errors.InvalidScope("Empty scope (either the client or the user is not allowed the requested scopes)")
	}

	return nil
}

func (t *TokenEndpoint) isRefreshTokenRequest(parameters model.RequestParameters) bool {
	return parameters.GrantType == constants.GrantTypeRefresh && parameters.RefreshToken != ""
}

func (t *TokenEndpoint) isAuthCodeRequest(parameters model.RequestParameters) bool {
	return parameters.GrantType == constants.GrantTypeCode && parameters.Code != ""
}
