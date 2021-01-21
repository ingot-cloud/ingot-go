package dto

// LoginParams for request
type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	AppID    string `json:"appID" binding:"required"`
}

// LoginResult for response
type LoginResult struct {
	Username    string   `json:"username"`
	Role        []string `json:"role"`
	AccessToken string   `json:"accessToken"`
	TokenType   string   `json:"tokenType"`
	Expiration  int64    `json:"expiration"`
}
