package container

// SecurityContainerImpl 接口实现
type SecurityContainerImpl struct {
	CommonContainer              *CommonContainer              `container:"true"`
	OAuth2Container              *OAuth2Container              `container:"true"`
	ResourceServerContainer      *ResourceServerContainer      `container:"true"`
	AuthorizationServerContainer *AuthorizationServerContainer `container:"true"`
	AuthProvidersContainer       *AuthProvidersContainer       `container:"true"`
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
