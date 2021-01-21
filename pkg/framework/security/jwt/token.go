package jwt

import "time"

// AccessToken impl security.AccessToken
type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiration  int64  `json:"expiration"`
}

// GetValue 获取Token
func (t *AccessToken) GetValue() string {
	return t.AccessToken
}

// GetTokenType 获取Token类型
func (t *AccessToken) GetTokenType() string {
	return t.TokenType
}

// GetExpiration 获取到期时间戳
func (t *AccessToken) GetExpiration() int64 {
	return t.Expiration
}

// IsExpired 是否过期
func (t *AccessToken) IsExpired() bool {
	return time.Unix(t.Expiration, 0).Sub(time.Now()) < 0
}
