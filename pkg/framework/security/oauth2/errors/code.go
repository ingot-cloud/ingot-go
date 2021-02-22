package errors

// Token Code
const (
	TokenInvalid = "S0001"
	TokenExpired = "S0002"
)

// OAuth2 error code
const (
	InvalidRequestCode          = "invalid_request"
	InvalidClientCode           = "invalid_client"
	InvalidTokenCode            = "invalid_token"
	InvalidScopeCode            = "invalid_scope"
	InvalidGrantCode            = "invalid_grant"
	InsufficientScopeCode       = "insufficient_scope"
	UnauthorizedClientCode      = "unauthorized_client"
	UnauthorizedUserCode        = "unauthorized_user"
	UnsupportedGrantTypeCode    = "unsupported_grant_type"
	UnsupportedResponseTypeCode = "unsupported_response_type"
	AccessDeniedCode            = "access_denied"
)
