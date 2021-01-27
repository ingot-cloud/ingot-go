package config

// BuildWebSecurity 构建 WebSecurity
func BuildWebSecurity(configurers WebSecurityConfigurers) (*WebSecurity, error) {
	webSecurity := &WebSecurity{}

	for _, configurer := range configurers {
		webSecurity.Apply(configurer)
	}

	return webSecurity, nil
}
