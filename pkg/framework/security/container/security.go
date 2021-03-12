package container

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

// SecurityContainerProxyImpl 默认实现
type SecurityContainerProxyImpl struct {
	SecurityContainer SecurityContainer
	SecurityInjector  SecurityInjector
}

// GetSecurityContainer 获取安全容器
func (p *SecurityContainerProxyImpl) GetSecurityContainer() SecurityContainer {
	return p.SecurityContainer
}

// GetSecurityInjector 获取注入器
func (p *SecurityContainerProxyImpl) GetSecurityInjector() SecurityInjector {
	return p.SecurityInjector
}
