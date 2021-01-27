package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// Config struct
type Config struct {
	App    App           `yaml:"app"`
	Server server.Config `yaml:"server"`
	Log    log.Config    `yaml:"log"`
	Gorm   Gorm          `yaml:"gorm"`
	MySQL  MySQL         `yaml:"mysql"`
	Redis  Redis         `yaml:"redis"`
	Auth   Auth          `yaml:"auth"`
}

// App struct
type App struct {
	Name string `yaml:"name"`
}

// Gorm config
type Gorm struct {
	Debug             bool   `yaml:"debug"`
	DBType            string `yaml:"dbType"`
	MaxLifetime       int    `yaml:"maxLifetime"`
	MaxOpenConns      int    `yaml:"maxOpenConns"`
	MaxIdleConns      int    `yaml:"maxIdleConns"`
	EnableAutoMigrate bool   `yaml:"enableAutoMigrate"`
}

// MySQL config
type MySQL struct {
	Host                      string `yaml:"host"`
	Port                      int    `yaml:"port"`
	User                      string `yaml:"user"`
	Password                  string `yaml:"password"`
	DBName                    string `yaml:"dbName"`
	Parameters                string `yaml:"parameters"`
	DefaultStringSize         uint   `yaml:"defaultStringSize"`
	DisableDatetimePrecision  bool   `yaml:"disableDatetimePrecision"`
	DontSupportRenameIndex    bool   `yaml:"dontSupportRenameIndex"`
	DontSupportRenameColumn   bool   `yaml:"dontSupportRenameColumn"`
	SkipInitializeWithVersion bool   `yaml:"skipInitializeWithVersion"`
}

// Redis config
type Redis struct {
	Address   string `yaml:"address"`
	DB        int    `yaml:"db"`
	Password  string `yaml:"password"`
	KeyPrefix string `yaml:"keyPrefix"`
	SSL       bool   `yaml:"ssl"`
}

// Auth config
type Auth struct {
	PermitUrls []string `yaml:"permitUrls"`
	Jwt        Jwt      `yaml:"jwt"`
}

// Jwt config
type Jwt struct {
	SigningMethod string `yaml:"signingMethod"`
	SigningKey    string `yaml:"signingKey"`
	Expired       int    `yaml:"expired"`
}
