package container

// SecurityContainer 容器汇总
type SecurityContainer struct {
	CommonContainer              *CommonContainer
	OAuth2Container              *OAuth2Container
	ResourceServerContainer      *ResourceServerContainer
	AuthorizationServerContainer *AuthorizationServerContainer
	AuthProvidersContainer       *AuthProvidersContainer
}
