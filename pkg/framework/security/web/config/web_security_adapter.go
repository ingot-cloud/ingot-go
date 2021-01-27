package config

// WebSecurityConfigurerAdapter 安全配置适配器
type WebSecurityConfigurerAdapter struct {
	AdditionalConfigurer HTTPSecurityConfigurer
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

	err = adapter.AdditionalConfigurer.Configure(http)

	return http, err
}

func (adapter *WebSecurityConfigurerAdapter) applyDefaultConfiguration(http *HTTPSecurity) error {
	// 应用默认配置
	// http.A()  http.B()
	return nil
}
