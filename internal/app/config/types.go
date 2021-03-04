package config

import (
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	oauth2Config "github.com/ingot-cloud/ingot-go/pkg/framework/security/oauth2/model"
)

// Config struct
type Config struct {
	App      App               `yaml:"app"`
	Server   config.HTTPConfig `yaml:"server"`
	Log      log.Config        `yaml:"log"`
	Gorm     Gorm              `yaml:"gorm"`
	MySQL    MySQL             `yaml:"mysql"`
	Redis    Redis             `yaml:"redis"`
	Security Security          `yaml:"security"`
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

// Security config
type Security struct {
	EnableResourceServer      bool                `yaml:"enableResourceServer"`
	EnableAuthorizationServer bool                `yaml:"enableAuthorizationServer"`
	PermitURLs                []string            `yaml:"permitUrls"`
	OAuth2                    oauth2Config.OAuth2 `yaml:"oauth2"`
}
