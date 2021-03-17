package provider

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container/di"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/container/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
)

// All 所有实例
var All = wire.NewSet(
	/* CommonContainer start */
	wire.Struct(new(securityContainer.CommonContainer), "*"),

	// Fields
	PasswordEncoder,
	UserCache,
	PreChecker,
	PostChecker,
	UserDetailsService,
	ClientDetailsService,
	wire.Struct(new(WebSecurityConfigurersImpl)),
	wire.Bind(new(security.WebSecurityConfigurers), new(*WebSecurityConfigurersImpl)),
	/* CommonContainer end */

	/* OAuth2Container start */
	wire.Struct(new(securityContainer.OAuth2Container), "*"),

	// Fields
	TokenStore,
	JwtAccessTokenConverter,
	AccessTokenConverter,
	UserAuthenticationConverter,
	/* OAuth2Container end */

	/* AuthorizationServerContainer start */
	wire.Struct(new(securityContainer.AuthorizationServerContainer), "*"),

	// Fields
	AuthorizationAuthenticationManager,
	AuthorizationServerConfigurer,
	AuthorizationServerTokenServices,
	ConsumerTokenServices,
	TokenEndpoint,
	TokenEndpointHTTPConfigurer,
	TokenEnhancer,
	TokenGranter,
	PasswordTokenGranter,
	/* AuthorizationServerContainer end */

	/* ResourceServerContainer start */
	wire.Struct(new(securityContainer.ResourceServerContainer), "*"),

	// Fields
	ResourceAuthenticationManager,
	ResourceServerConfigurer,
	ResourceServerTokenServices,
	TokenExtractor,
	/* ResourceServerContainer end */

	/* AuthProvidersContainer start */
	wire.Struct(new(securityContainer.AuthProvidersContainer), "*"),

	// Fields
	DaoAuthenticationProvider,
	BasicAuthenticationProvider,
	wire.Struct(new(ProvidersImpl), "Basic", "Dao"),
	wire.Bind(new(authentication.Providers), new(*ProvidersImpl)),
	/* AuthProvidersContainer end */

	// 容器相关
	wire.NewSet(wire.Struct(new(securityContainer.NilSecurityInjector), "*")),

	wire.Struct(new(securityContainer.SecurityContainerImpl), "*"),
	wire.Bind(new(securityContainer.SecurityContainer), new(*securityContainer.SecurityContainerImpl)),

	// wire.Struct(new(container.SecurityContainerPreProxyImpl), "*"),
	// wire.Bind(new(container.SecurityContainerPreProxy), new(*container.SecurityContainerPreProxyImpl)),

	// InjectCustomInstance,
)

// InjectCustomInstance 注入自定义实例
// func InjectCustomInstance(proxy container.SecurityContainerPreProxy) container.SecurityContainerCombine {
// 	injector := proxy.GetSecurityInjector()
// 	sc := proxy.GetSecurityContainer()
// 	return process.DoPre(injector, sc)
// }

var ProviderSet = di.NewSet(
	/* CommonContainer start */
	di.Struct(new(securityContainer.CommonContainer)),

	// Fields
	di.Func(PasswordEncoder),
	di.Func(UserCache),
	di.Func(PreChecker),
	di.Func(PostChecker),
	di.Func(UserDetailsService),
	di.Func(ClientDetailsService),
	di.Struct(new(WebSecurityConfigurersImpl)),
	di.Bind(new(security.WebSecurityConfigurers), new(WebSecurityConfigurersImpl)),
	/* CommonContainer end */

	/* OAuth2Container start */
	di.Struct(new(securityContainer.OAuth2Container)),

	// Fields
	di.Func(TokenStore),
	di.Func(JwtAccessTokenConverter),
	di.Func(AccessTokenConverter),
	di.Func(UserAuthenticationConverter),
	/* OAuth2Container end */

	/* AuthorizationServerContainer start */
	di.Struct(new(securityContainer.AuthorizationServerContainer)),

	// Fields
	di.Func(AuthorizationAuthenticationManager),
	di.Func(AuthorizationServerConfigurer),
	di.Func(AuthorizationServerTokenServices),
	di.Func(ConsumerTokenServices),
	di.Func(TokenEndpoint),
	di.Func(TokenEndpointHTTPConfigurer),
	di.Func(TokenEnhancer),
	di.Func(TokenGranter),
	di.Func(PasswordTokenGranter),
	/* AuthorizationServerContainer end */

	/* ResourceServerContainer start */
	di.Struct(new(securityContainer.ResourceServerContainer)),

	// Fields
	di.Func(ResourceAuthenticationManager),
	di.Func(ResourceServerConfigurer),
	di.Func(ResourceServerTokenServices),
	di.Func(TokenExtractor),
	/* ResourceServerContainer end */

	/* AuthProvidersContainer start */
	di.Struct(new(securityContainer.AuthProvidersContainer)),

	// Fields
	di.Func(DaoAuthenticationProvider),
	di.Func(BasicAuthenticationProvider),
	di.Struct(new(ProvidersImpl), "Basic", "Dao"),
	di.Bind(new(authentication.Providers), new(ProvidersImpl)),
	/* AuthProvidersContainer end */

)
