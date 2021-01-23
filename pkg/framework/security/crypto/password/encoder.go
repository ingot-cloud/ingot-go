package password

// Encoder Service interface for encoding passwords.
type Encoder interface {
	// 对密码进行编码
	Encode(raw string) (string, error)
	// 验证原始密码和编码后的密码是否相等
	Matches(raw string, encodedPassword string) (bool, error)
}
