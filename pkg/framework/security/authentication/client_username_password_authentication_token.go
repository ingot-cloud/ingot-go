package authentication

// ClientUsernamePasswordAuthenticationToken client
type ClientUsernamePasswordAuthenticationToken struct {
	*UsernamePasswordAuthenticationToken
}

// NewClientUsernamePasswordAuthToken client token
func NewClientUsernamePasswordAuthToken(principal interface{}, credentials string) *ClientUsernamePasswordAuthenticationToken {
	token := NewUnauthenticatedUsernamePasswordAuthToken(principal, credentials)
	return &ClientUsernamePasswordAuthenticationToken{
		UsernamePasswordAuthenticationToken: token,
	}
}
