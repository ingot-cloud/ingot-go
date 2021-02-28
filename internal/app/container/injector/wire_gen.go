// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/container"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider/factory"
	"github.com/ingot-cloud/ingot-go/internal/app/core/http"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/internal/app/service"
	container2 "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	dao2 "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	container3 "github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	provider2 "github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider"
	config2 "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/provider/token"
)

// Injectors from app.go:

func BuildConfiguration(options *config.Options) (*config.Config, error) {
	configConfig, err := provider.NewConfig(options)
	if err != nil {
		return nil, err
	}
	return configConfig, nil
}

func BuildAppContainer(config2 *config.Config, options *config.Options) (*container.AppContainer, func(), error) {
	httpConfig, err := factory.HTTPConfigSet(config2)
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := factory.NewGorm(config2)
	if err != nil {
		return nil, nil, err
	}
	role := &dao.Role{
		DB: db,
	}
	roleAuthority := &dao.RoleAuthority{
		DB: db,
	}
	authority := &dao.Authority{
		DB: db,
	}
	user := &dao.User{
		DB: db,
	}
	roleUser := &dao.RoleUser{
		DB: db,
	}
	permission := &service.Permission{
		RoleDao:          role,
		RoleAuthorityDao: roleAuthority,
		AuthorityDao:     authority,
		UserDao:          user,
		RoleUserDao:      roleUser,
	}
	casbinAdapterService := &service.CasbinAdapterService{
		PermissionService: permission,
	}
	syncedEnforcer, cleanup2, err := factory.NewCasbin(options, casbinAdapterService)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	security, err := factory.SecurityConfigSet(config2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	auth := &service.Auth{
		UserDao:     user,
		RoleUserDao: roleUser,
		RoleDao:     role,
	}
	apiAuth := &api.Auth{
		AuthService: auth,
	}
	apiConfig := &http.APIConfig{
		CasbinEnforcer: syncedEnforcer,
		HTTPConfig:     httpConfig,
		SecurityConfig: security,
		AuthAPI:        apiAuth,
	}
	appContainer := &container.AppContainer{
		HTTPConfig:     httpConfig,
		HTTPConfigurer: apiConfig,
	}
	return appContainer, func() {
		cleanup2()
		cleanup()
	}, nil
}

func BuildContainer(httpInjector container2.HTTPInjector, securityContainer container3.SecurityAllContainer) (*container2.Container, func(), error) {
	containerContainer := container2.BuildContainer(httpInjector, securityContainer)
	return containerContainer, func() {
	}, nil
}

// Injectors from security.go:

func BuildSecurityContainer(injector container3.SecurityInjector) (*container3.SecurityContainer, error) {
	webSecurityConfigurers := provider2.WebSecurityConfigurers(injector)
	encoder := provider2.PasswordEncoder(injector)
	userdetailsService := provider2.UserDetailsService(injector)
	userCache := provider2.UserCache(injector)
	preChecker := provider2.PreChecker(injector)
	postChecker := provider2.PostChecker(injector)
	authenticationProvider := &dao2.AuthenticationProvider{
		PasswordEncoder:          encoder,
		UserDetailsService:       userdetailsService,
		UserCache:                userCache,
		PreAuthenticationChecks:  preChecker,
		PostAuthenticationChecks: postChecker,
	}
	providers := provider2.Providers(authenticationProvider, injector)
	clientdetailsService := provider2.ClientDetailsService(injector)
	securityContainer := &container3.SecurityContainer{
		WebSecurityConfigurers: webSecurityConfigurers,
		Providers:              providers,
		PasswordEncoder:        encoder,
		UserCache:              userCache,
		PreChecker:             preChecker,
		PostChecker:            postChecker,
		UserDetailsService:     userdetailsService,
		ClientDetailsService:   clientdetailsService,
	}
	return securityContainer, nil
}

func BuildOAuth2Container(oauth2Config config2.OAuth2, injector container3.SecurityInjector) (*container3.OAuth2Container, error) {
	userAuthenticationConverter := provider2.UserAuthenticationConverter(injector)
	accessTokenConverter := provider2.AccessTokenConverter(oauth2Config, userAuthenticationConverter, injector)
	jwtAccessTokenConverter := provider2.JwtAccessTokenConverter(oauth2Config, accessTokenConverter)
	store := provider2.TokenStore(jwtAccessTokenConverter, injector)
	defaultTokenServices := provider2.DefaultTokenServices(oauth2Config, store)
	oAuth2Container := &container3.OAuth2Container{
		Config:                      oauth2Config,
		DefaultTokenServices:        defaultTokenServices,
		TokenStore:                  store,
		JwtAccessTokenConverter:     jwtAccessTokenConverter,
		AccessTokenConverter:        accessTokenConverter,
		UserAuthenticationConverter: userAuthenticationConverter,
	}
	return oAuth2Container, nil
}

func BuildResourceServerContainer(oauth2Container *container3.OAuth2Container, injector container3.SecurityInjector) (*container3.ResourceServerContainer, error) {
	resourceServerTokenServices := provider2.ResourceServerTokenServices(oauth2Container, injector)
	tokenExtractor := provider2.TokenExtractor(injector)
	manager := provider2.ResourceAuthenticationManager(oauth2Container, resourceServerTokenServices, injector)
	oAuth2SecurityConfigurer := provider2.OAuth2SecurityConfigurer(tokenExtractor, manager)
	resourceServerContainer := &container3.ResourceServerContainer{
		ResourceServerTokenServices: resourceServerTokenServices,
		OAuth2SecurityConfigurer:    oAuth2SecurityConfigurer,
		TokenExtractor:              tokenExtractor,
		AuthenticationManager:       manager,
	}
	return resourceServerContainer, nil
}

func BuildAuthorizationServerContainer(oauth2Container *container3.OAuth2Container, securityContainer *container3.SecurityContainer, enhancers token.Enhancers, injector container3.SecurityInjector) (*container3.AuthorizationServerContainer, error) {
	enhancer := provider2.TokenEnhancer(enhancers, oauth2Container, injector)
	manager := provider2.AuthorizationAuthenticationManager(securityContainer, injector)
	authorizationServerTokenServices := provider2.AuthorizationServerTokenServices(oauth2Container, securityContainer, enhancer, manager, injector)
	consumerTokenServices := provider2.ConsumerTokenServices(oauth2Container, injector)
	authorizationServerContainer := &container3.AuthorizationServerContainer{
		AuthorizationServerTokenServices: authorizationServerTokenServices,
		ConsumerTokenServices:            consumerTokenServices,
		TokenEnhancer:                    enhancer,
		AuthenticationManager:            manager,
	}
	return authorizationServerContainer, nil
}
