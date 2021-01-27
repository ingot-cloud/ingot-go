package provider

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
	"github.com/spf13/viper"
)

// LoadConfig 加载配置
func LoadConfig(options *config.Options) (*config.Config, error) {
	configFile := options.ConfigFile
	fmt.Printf("Use profile = %s\n", configFile)

	loadConfigFile := configFile

	var config config.Config

	v := viper.New()
	v.SetConfigFile(loadConfigFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file(%s) changed: %s", loadConfigFile, e.Name)
		if err := v.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})

	err := v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// HTTPConfigSet 单独注入 http config
func HTTPConfigSet(config *config.Config) (server.Config, error) {
	return config.Server, nil
}

// AuthConfigSet 单独注入 auth config
func AuthConfigSet(config *config.Config) (config.Auth, error) {
	return config.Auth, nil
}

// ConfigSet 需要单独注入的配置
var ConfigSet = wire.NewSet(
	HTTPConfigSet,
	AuthConfigSet,
)
