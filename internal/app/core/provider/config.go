package provider

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/spf13/viper"
)

// NewConfig 加载配置
func NewConfig(options *config.Options) (*config.Config, error) {
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
