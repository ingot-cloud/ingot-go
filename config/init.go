package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// CONFIG 配置
var CONFIG Config

const loadConfigFile = "config/config.yml"

func init() {
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

	if err := v.Unmarshal(&CONFIG); err != nil {
		fmt.Println(err)
	}
}
