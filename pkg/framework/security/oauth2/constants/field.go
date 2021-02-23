package constants

// OAuth2 请求映射参数常量
const (
	GrantType         = "grant_type"
	Scope             = "scope"
	ClientID          = "client_id"
	ClientSecret      = "client_secret"
	State             = "state"
	RedirectURI       = "redirect_uri"
	ResponseType      = "response_type"
	Code              = "code"
	Username          = "username"
	Password          = "password"
	RefreshToken      = "refresh_token"
	UserOAuthApproval = "user_oauth_approval"
	ScopePrefix       = "scope."
)

// TokenPayloadKey Token载体key
type TokenPayloadKey string

// Token载体key
const (
	TokenUsername    TokenPayloadKey = "username"
	TokenUser        TokenPayloadKey = "user"
	TokenAud         TokenPayloadKey = "aud"
	TokenClientID    TokenPayloadKey = "client_id"
	TokenExp         TokenPayloadKey = "exp"
	TokenJti         TokenPayloadKey = "jti"
	TokenGrantType   TokenPayloadKey = "grant_type"
	TokenAti         TokenPayloadKey = "ati"
	TokenScope       TokenPayloadKey = "scope"
	TokenAuthorities TokenPayloadKey = "authorities"
)
