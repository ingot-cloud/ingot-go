package authentication

// ClientUsernamePasswordAuthenticationToken client
type ClientUsernamePasswordAuthenticationToken struct {
	*UsernamePasswordAuthenticationToken
}

// NewClientUsernamePasswordAuthToken client token
func NewClientUsernamePasswordAuthToken(principal interface{}, credentials string) *ClientUsernamePasswordAuthenticationToken {
	token := &UsernamePasswordAuthenticationToken{
		Principal:   principal,
		Credentials: credentials,
	}
	token.SetAuthenticated(false)
	return &ClientUsernamePasswordAuthenticationToken{
		UsernamePasswordAuthenticationToken: token,
	}
}
