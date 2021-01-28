package core

// CredentialsContainer 将实现结构中的敏感数据擦除
type CredentialsContainer interface {
	// 擦除敏感数据
	EraseCredentials()
}
