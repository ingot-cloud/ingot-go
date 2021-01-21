package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// CONFIG 配置
var CONFIG Config

// LoadConfig 加载配置
func LoadConfig(configFile string) (func(), error) {
	fmt.Printf("Use profile = %s\n", configFile)

	// loadConfigFile := fmt.Sprintf("config/%s", configFile)
	loadConfigFile := configFile

	v := viper.New()
	v.SetConfigFile(loadConfigFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file(%s) changed: %s", loadConfigFile, e.Name)
		if err := v.Unmarshal(&CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	return func() {
		// 空的clear方法
	}, v.Unmarshal(&CONFIG)
}
