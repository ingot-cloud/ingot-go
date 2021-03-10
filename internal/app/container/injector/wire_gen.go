// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	container3 "github.com/ingot-cloud/ingot-go/internal/app/container"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider"
	"github.com/ingot-cloud/ingot-go/internal/app/container/provider/factory"
	"github.com/ingot-cloud/ingot-go/internal/app/core/http"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/service"
	"github.com/ingot-cloud/ingot-go/internal/app/core/security/token"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/internal/app/service/impl"
	container2 "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	dao2 "github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication/provider/dao"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	provider2 "github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/preset"
)

// Injectors from app.go:

func BuildConfiguration(options *config.Options) (*config.Config, error) {
	configConfig, err := provider.NewConfig(options)
	if err != nil {
		return nil, err
	}
	return configConfig, nil
}

func BuildContainerInjector(config2 *config.Config, options *config.Options) (container.SecurityInjector, func(), error) {
	nilSecurityInjector := &container.NilSecurityInjector{}
	httpConfig, err := factory.HTTPConfig(config2)
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
	permission := &impl.Permission{
		RoleDao:          role,
		RoleAuthorityDao: roleAuthority,
		AuthorityDao:     authority,
		UserDao:          user,
		RoleUserDao:      roleUser,
	}
	casbinAdapterService := &impl.CasbinAdapterService{
		PermissionService: permission,
	}
	syncedEnforcer, cleanup2, err := factory.NewCasbin(options, casbinAdapterService)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	security, err := factory.SecurityConfig(config2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	test := &api.Test{}
	apiConfig := &http.APIConfig{
		CasbinEnforcer: syncedEnforcer,
		SecurityConfig: security,
		TestAPI:        test,
	}
	webSecurityConfigurersImpl := &preset.WebSecurityConfigurersImpl{}
	encoder := preset.PasswordEncoder()
	userCache := preset.UserCache()
	preChecker := preset.PreChecker()
	postChecker := preset.PostChecker()
	userdetailsService := preset.UserDetailsService()
	clientdetailsService := preset.ClientDetailsService()
	commonContainer := &container.CommonContainer{
		WebSecurityConfigurers: webSecurityConfigurersImpl,
		PasswordEncoder:        encoder,
		UserCache:              userCache,
		PreChecker:             preChecker,
		PostChecker:            postChecker,
		UserDetailsService:     userdetailsService,
		ClientDetailsService:   clientdetailsService,
	}
	oAuth2, err := factory.OAuth2Config(config2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userAuthenticationConverter := preset.UserAuthenticationConverter()
	accessTokenConverter := preset.AccessTokenConverter(oAuth2, userAuthenticationConverter)
	jwtAccessTokenConverter := preset.JwtAccessTokenConverter(oAuth2, accessTokenConverter)
	store := preset.TokenStore(jwtAccessTokenConverter)
	oAuth2Container := &container.OAuth2Container{
		Config:                      oAuth2,
		TokenStore:                  store,
		JwtAccessTokenConverter:     jwtAccessTokenConverter,
		AccessTokenConverter:        accessTokenConverter,
		UserAuthenticationConverter: userAuthenticationConverter,
	}
	resourceServerTokenServices := preset.ResourceServerTokenServices(store)
	resourceManager := preset.ResourceAuthenticationManager(oAuth2Container, resourceServerTokenServices)
	tokenExtractor := preset.TokenExtractor()
	resourceServerConfigurer := preset.ResourceServerConfigurer(tokenExtractor, resourceManager)
	resourceServerContainer := &container.ResourceServerContainer{
		AuthenticationManager:       resourceManager,
		ResourceServerConfigurer:    resourceServerConfigurer,
		ResourceServerTokenServices: resourceServerTokenServices,
		TokenExtractor:              tokenExtractor,
	}
	authenticationProvider := preset.BasicAuthenticationProvider(commonContainer)
	daoAuthenticationProvider := &dao2.AuthenticationProvider{
		PasswordEncoder:          encoder,
		UserDetailsService:       userdetailsService,
		UserCache:                userCache,
		PreAuthenticationChecks:  preChecker,
		PostAuthenticationChecks: postChecker,
	}
	providersImpl := &preset.ProvidersImpl{
		Basic: authenticationProvider,
		Dao:   daoAuthenticationProvider,
	}
	authProvidersContainer := &container.AuthProvidersContainer{
		Providers: providersImpl,
	}
	authorizationManager := preset.AuthorizationAuthenticationManager(authProvidersContainer)
	authorizationServerConfigurer := preset.AuthorizationServerConfigurer(authorizationManager)
	enhancer := preset.TokenEnhancer(oAuth2Container)
	authorizationServerTokenServices := preset.AuthorizationServerTokenServices(oAuth2, store, commonContainer, enhancer, authorizationManager)
	consumerTokenServices := preset.ConsumerTokenServices(store)
	passwordTokenGranter := preset.PasswordTokenGranter(authorizationServerTokenServices, authorizationManager)
	granter := preset.TokenGranter(passwordTokenGranter)
	tokenEndpoint := preset.TokenEndpoint(granter, commonContainer)
	oAuth2HTTPConfigurer := preset.TokenEndpointHTTPConfigurer(tokenEndpoint)
	authorizationServerContainer := &container.AuthorizationServerContainer{
		AuthenticationManager:            authorizationManager,
		AuthorizationServerConfigurer:    authorizationServerConfigurer,
		AuthorizationServerTokenServices: authorizationServerTokenServices,
		ConsumerTokenServices:            consumerTokenServices,
		TokenEndpoint:                    tokenEndpoint,
		TokenEndpointHTTPConfigurer:      oAuth2HTTPConfigurer,
		TokenEnhancer:                    enhancer,
		TokenGranter:                     granter,
		PasswordTokenGranter:             passwordTokenGranter,
	}
	securityContainerImpl := &container.SecurityContainerImpl{
		CommonContainer:              commonContainer,
		OAuth2Container:              oAuth2Container,
		ResourceServerContainer:      resourceServerContainer,
		AuthorizationServerContainer: authorizationServerContainer,
		AuthProvidersContainer:       authProvidersContainer,
	}
	defaultPre := &container2.DefaultPre{
		HTTPConfig:        httpConfig,
		HTTPConfigurer:    apiConfig,
		SecurityContainer: securityContainerImpl,
	}
	oauthClientDetails := &dao.OauthClientDetails{
		DB: db,
	}
	clientDetails := &service.ClientDetails{
		OauthClientDetailsDao: oauthClientDetails,
	}
	userDetail := &impl.UserDetail{}
	userDetails := &service.UserDetails{
		UserDetailService: userDetail,
	}
	requestMatcher := provider.PermitURLMatcher(security)
	resourceServerAdapter := provider.ResourceServerAdapter(tokenExtractor, resourceManager, requestMatcher)
	ingotEnhancer := &token.IngotEnhancer{}
	appContainer := &container3.AppContainer{
		NilSecurityInjector:   nilSecurityInjector,
		DefaultPre:            defaultPre,
		SecurityConfig:        security,
		ClientDetailsService:  clientDetails,
		UserDetailsService:    userDetails,
		ResourceServerAdapter: resourceServerAdapter,
		IngotEnhancer:         ingotEnhancer,
	}
	return appContainer, func() {
		cleanup2()
		cleanup()
	}, nil
}

func BuildContainer(config2 *config.Config, options *config.Options, securityInjector container.SecurityInjector) (container2.Container, func(), error) {
	httpConfig, err := factory.HTTPConfig(config2)
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
	permission := &impl.Permission{
		RoleDao:          role,
		RoleAuthorityDao: roleAuthority,
		AuthorityDao:     authority,
		UserDao:          user,
		RoleUserDao:      roleUser,
	}
	casbinAdapterService := &impl.CasbinAdapterService{
		PermissionService: permission,
	}
	syncedEnforcer, cleanup2, err := factory.NewCasbin(options, casbinAdapterService)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	security, err := factory.SecurityConfig(config2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	test := &api.Test{}
	apiConfig := &http.APIConfig{
		CasbinEnforcer: syncedEnforcer,
		SecurityConfig: security,
		TestAPI:        test,
	}
	webSecurityConfigurersImpl := &provider2.WebSecurityConfigurersImpl{
		Injector: securityInjector,
	}
	encoder := provider2.PasswordEncoder(securityInjector)
	userCache := provider2.UserCache(securityInjector)
	preChecker := provider2.PreChecker(securityInjector)
	postChecker := provider2.PostChecker(securityInjector)
	userdetailsService := provider2.UserDetailsService(securityInjector)
	clientdetailsService := provider2.ClientDetailsService(securityInjector)
	commonContainer := &container.CommonContainer{
		WebSecurityConfigurers: webSecurityConfigurersImpl,
		PasswordEncoder:        encoder,
		UserCache:              userCache,
		PreChecker:             preChecker,
		PostChecker:            postChecker,
		UserDetailsService:     userdetailsService,
		ClientDetailsService:   clientdetailsService,
	}
	oAuth2, err := factory.OAuth2Config(config2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userAuthenticationConverter := provider2.UserAuthenticationConverter(securityInjector)
	accessTokenConverter := provider2.AccessTokenConverter(oAuth2, userAuthenticationConverter, securityInjector)
	jwtAccessTokenConverter := provider2.JwtAccessTokenConverter(oAuth2, accessTokenConverter, securityInjector)
	store := provider2.TokenStore(jwtAccessTokenConverter, securityInjector)
	oAuth2Container := &container.OAuth2Container{
		Config:                      oAuth2,
		TokenStore:                  store,
		JwtAccessTokenConverter:     jwtAccessTokenConverter,
		AccessTokenConverter:        accessTokenConverter,
		UserAuthenticationConverter: userAuthenticationConverter,
	}
	resourceServerTokenServices := provider2.ResourceServerTokenServices(store, securityInjector)
	resourceManager := provider2.ResourceAuthenticationManager(oAuth2Container, resourceServerTokenServices, securityInjector)
	tokenExtractor := provider2.TokenExtractor(securityInjector)
	resourceServerConfigurer := provider2.ResourceServerConfigurer(tokenExtractor, resourceManager, securityInjector)
	resourceServerContainer := &container.ResourceServerContainer{
		AuthenticationManager:       resourceManager,
		ResourceServerConfigurer:    resourceServerConfigurer,
		ResourceServerTokenServices: resourceServerTokenServices,
		TokenExtractor:              tokenExtractor,
	}
	authenticationProvider := provider2.BasicAuthenticationProvider(commonContainer)
	daoAuthenticationProvider := &dao2.AuthenticationProvider{
		PasswordEncoder:          encoder,
		UserDetailsService:       userdetailsService,
		UserCache:                userCache,
		PreAuthenticationChecks:  preChecker,
		PostAuthenticationChecks: postChecker,
	}
	providersImpl := &provider2.ProvidersImpl{
		Injector: securityInjector,
		Basic:    authenticationProvider,
		Dao:      daoAuthenticationProvider,
	}
	authProvidersContainer := &container.AuthProvidersContainer{
		Providers: providersImpl,
	}
	authorizationManager := provider2.AuthorizationAuthenticationManager(authProvidersContainer, securityInjector)
	authorizationServerConfigurer := provider2.AuthorizationServerConfigurer(authorizationManager, securityInjector)
	enhancer := provider2.TokenEnhancer(oAuth2Container, securityInjector)
	authorizationServerTokenServices := provider2.AuthorizationServerTokenServices(oAuth2, store, commonContainer, enhancer, authorizationManager, securityInjector)
	consumerTokenServices := provider2.ConsumerTokenServices(store, securityInjector)
	passwordTokenGranter := provider2.PasswordTokenGranter(authorizationServerTokenServices, authorizationManager, securityInjector)
	granter := provider2.TokenGranter(passwordTokenGranter, securityInjector)
	tokenEndpoint := provider2.TokenEndpoint(granter, commonContainer, securityInjector)
	oAuth2HTTPConfigurer := provider2.TokenEndpointHTTPConfigurer(tokenEndpoint, securityInjector)
	authorizationServerContainer := &container.AuthorizationServerContainer{
		AuthenticationManager:            authorizationManager,
		AuthorizationServerConfigurer:    authorizationServerConfigurer,
		AuthorizationServerTokenServices: authorizationServerTokenServices,
		ConsumerTokenServices:            consumerTokenServices,
		TokenEndpoint:                    tokenEndpoint,
		TokenEndpointHTTPConfigurer:      oAuth2HTTPConfigurer,
		TokenEnhancer:                    enhancer,
		TokenGranter:                     granter,
		PasswordTokenGranter:             passwordTokenGranter,
	}
	securityContainerImpl := &container.SecurityContainerImpl{
		CommonContainer:              commonContainer,
		OAuth2Container:              oAuth2Container,
		ResourceServerContainer:      resourceServerContainer,
		AuthorizationServerContainer: authorizationServerContainer,
		AuthProvidersContainer:       authProvidersContainer,
	}
	printSecurityInjector := provider2.PrintInjectInstance(securityContainerImpl)
	defaultContainer := &container2.DefaultContainer{
		HTTPConfig:         httpConfig,
		HTTPConfigurer:     apiConfig,
		SecurityInjector:   securityInjector,
		SecurityContainer:  securityContainerImpl,
		DebugPrintInjector: printSecurityInjector,
	}
	return defaultContainer, func() {
		cleanup2()
		cleanup()
	}, nil
}
