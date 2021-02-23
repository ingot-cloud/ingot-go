package constants

// OAuth2 请求映射参数常量
const (
	ClientID          = "client_id"
	State             = "state"
	Scope             = "scope"
	RedirectURI       = "redirect_uri"
	ResponseType      = "response_type"
	UserOAuthApproval = "user_oauth_approval"
	ScopePrefix       = "scope."
	GrantType         = "grant_type"
	Password          = "password"
	ClientSecret      = "client_secret"
	Code              = "code"
	RefreshToken      = "refresh_token"
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
