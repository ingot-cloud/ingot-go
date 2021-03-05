package container

// SecurityContainer 容器汇总
type SecurityContainer interface {
	GetCommonContainer() *CommonContainer
	GetOAuth2Container() *OAuth2Container
	GetResourceServerContainer() *ResourceServerContainer
	GetAuthorizationServerContainer() *AuthorizationServerContainer
	GetAuthProvidersContainer() *AuthProvidersContainer
}

// SecurityContainerImpl 接口实现
type SecurityContainerImpl struct {
	CommonContainer              *CommonContainer
	OAuth2Container              *OAuth2Container
	ResourceServerContainer      *ResourceServerContainer
	AuthorizationServerContainer *AuthorizationServerContainer
	AuthProvidersContainer       *AuthProvidersContainer
}

// GetCommonContainer 获取容器
func (s *SecurityContainerImpl) GetCommonContainer() *CommonContainer {
	return s.CommonContainer
}

// GetOAuth2Container 获取容器
func (s *SecurityContainerImpl) GetOAuth2Container() *OAuth2Container {
	return s.OAuth2Container
}

// GetResourceServerContainer 获取容器
func (s *SecurityContainerImpl) GetResourceServerContainer() *ResourceServerContainer {
	return s.ResourceServerContainer
}

// GetAuthorizationServerContainer 获取容器
func (s *SecurityContainerImpl) GetAuthorizationServerContainer() *AuthorizationServerContainer {
	return s.AuthorizationServerContainer
}

// GetAuthProvidersContainer 获取容器
func (s *SecurityContainerImpl) GetAuthProvidersContainer() *AuthProvidersContainer {
	return s.AuthProvidersContainer
}
