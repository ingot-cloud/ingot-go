package container

// SecurityContainer 安全容器实例
type SecurityContainer interface {
	GetCommonContainer() *CommonContainer
	GetOAuth2Container() *OAuth2Container
	GetResourceServerContainer() *ResourceServerContainer
	GetAuthorizationServerContainer() *AuthorizationServerContainer
	GetAuthProvidersContainer() *AuthProvidersContainer
}
