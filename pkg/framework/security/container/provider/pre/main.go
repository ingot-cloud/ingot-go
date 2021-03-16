package pre

import (
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/authentication"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security/container/provider/process"
)

// All 所有实例
var All = wire.NewSet(
	/* CommonContainer start */
	wire.Struct(new(container.CommonContainer), "*"),

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
	wire.Struct(new(container.OAuth2Container), "*"),

	// Fields
	TokenStore,
	JwtAccessTokenConverter,
	AccessTokenConverter,
	UserAuthenticationConverter,
	/* OAuth2Container end */

	/* AuthorizationServerContainer start */
	wire.Struct(new(container.AuthorizationServerContainer), "*"),

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
	wire.Struct(new(container.ResourceServerContainer), "*"),

	// Fields
	ResourceAuthenticationManager,
	ResourceServerConfigurer,
	ResourceServerTokenServices,
	TokenExtractor,
	/* ResourceServerContainer end */

	/* AuthProvidersContainer start */
	wire.Struct(new(container.AuthProvidersContainer), "*"),

	// Fields
	DaoAuthenticationProvider,
	BasicAuthenticationProvider,
	wire.Struct(new(ProvidersImpl), "Basic", "Dao"),
	wire.Bind(new(authentication.Providers), new(*ProvidersImpl)),
	/* AuthProvidersContainer end */

	// 容器相关
	wire.NewSet(wire.Struct(new(container.NilSecurityInjector), "*")),

	wire.Struct(new(container.SecurityContainerImpl), "*"),
	wire.Bind(new(container.SecurityContainerPre), new(*container.SecurityContainerImpl)),

	wire.Struct(new(container.SecurityContainerPreProxyImpl), "*"),
	wire.Bind(new(container.SecurityContainerPreProxy), new(*container.SecurityContainerPreProxyImpl)),

	InjectCustomInstance,
)

// InjectCustomInstance 注入自定义实例
func InjectCustomInstance(proxy container.SecurityContainerPreProxy) container.SecurityContainerCombine {
	injector := proxy.GetSecurityInjector()
	sc := proxy.GetSecurityContainer()
	return process.DoPre(injector, sc)
}
