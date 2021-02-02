package config

import "time"

// HTTPConfig 服务配置
type HTTPConfig struct {
	Mode         string        `yaml:"mode"`
	Address      string        `yaml:"address"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	Prefix       string        `yaml:"prefix"`
}
