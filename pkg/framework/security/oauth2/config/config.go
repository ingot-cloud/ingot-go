package config

// OAuth2 配置
type OAuth2 struct {
	// 是否包含 grantType
	IncludeGrantType    bool                `yaml:"includeGrantType"`
	Jwt                 Jwt                 `yaml:"jwt"`
	ResourceServer      ResourceServer      `yaml:"resourceServer"`
	AuthorizationServer AuthorizationServer `yaml:"authorizationServer"`
}

// Jwt config
type Jwt struct {
	SigningMethod string `yaml:"signingMethod"`
	SigningKey    string `yaml:"signingKey"`
}

// ResourceServer 资源服务器配置
type ResourceServer struct {
	Enable     bool   `yaml:"enable"`
	ResourceID string `yaml:"resourceID"`
}

// AuthorizationServer 授权服务器配置
type AuthorizationServer struct {
	Enable bool `yaml:"enable"`
	// 是否支持RefreshToken
	SupportRefreshToken bool `yaml:"supportRefreshToken"`
	// 是否重复使用RefreshToken
	ReuseRefreshToken bool `yaml:"reuseRefreshToken"`
}
