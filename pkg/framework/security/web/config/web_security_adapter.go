package config

// WebSecurityConfigurerAdapter 安全配置适配器
type WebSecurityConfigurerAdapter struct {
	AdditionalConfigurers []HTTPSecurityConfigurer
}

// Configure Web安全配置
func (adapter *WebSecurityConfigurerAdapter) Configure(web *WebSecurity) error {
	http, err := adapter.getHTTP()
	if err != nil {
		return err
	}

	web.AddSecurityFilterChainBuilder(http)

	return nil
}

func (adapter *WebSecurityConfigurerAdapter) getHTTP() (*HTTPSecurity, error) {

	http := &HTTPSecurity{}

	err := adapter.applyDefaultConfiguration(http)
	if err != nil {
		return nil, err
	}

	err = adapter.applyAdditionalConfiguration(http)
	if err != nil {
		return nil, err
	}

	return http, nil
}

func (adapter *WebSecurityConfigurerAdapter) applyDefaultConfiguration(http *HTTPSecurity) error {
	// 应用默认配置
	// http.A()  http.B()
	return nil
}

func (adapter *WebSecurityConfigurerAdapter) applyAdditionalConfiguration(http *HTTPSecurity) error {
	for _, configurer := range adapter.AdditionalConfigurers {
		if err := configurer(http); err != nil {
			return err
		}
	}
	return nil
}
