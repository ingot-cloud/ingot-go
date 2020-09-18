package provider

import (
	"errors"
	"fmt"
	"ingot/common/log"
	"ingot/config"
	"ingot/model/domain"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// BuildGorm init gorm
func BuildGorm() (*gorm.DB, func(), error) {

	cfg := config.CONFIG
	db, cleanFunc, err := newGormDB(cfg)
	if err != nil {
		return nil, cleanFunc, err
	}

	if cfg.Gorm.EnableAutoMigrate {
		err := autoMigrate(db, cfg.Gorm.DBType)
		if err != nil {
			return nil, cleanFunc, err
		}
	}

	return db, cleanFunc, nil
}

func newGormDB(cfg config.Config) (*gorm.DB, func(), error) {
	dbType := cfg.Gorm.DBType
	if dbType != "mysql" {
		return nil, nil, errors.New("unknown db")
	}

	mysqlCfg := cfg.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.Port, mysqlCfg.DBName, mysqlCfg.Parameters)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         mysqlCfg.DefaultStringSize,
		DisableDatetimePrecision:  mysqlCfg.DisableDatetimePrecision,
		DontSupportRenameIndex:    mysqlCfg.DontSupportRenameIndex,
		DontSupportRenameColumn:   mysqlCfg.DontSupportRenameColumn,
		SkipInitializeWithVersion: mysqlCfg.SkipInitializeWithVersion,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	if cfg.Gorm.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {
		err := sqlDB.Close()
		if err != nil {
			log.Errorf("Gorm db close error: %s", err.Error())
		}
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, cleanFunc, err
	}

	sqlDB.SetMaxIdleConns(cfg.Gorm.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Gorm.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Gorm.MaxLifetime) * time.Second)
	return db, cleanFunc, nil
}

func autoMigrate(db *gorm.DB, dbType string) error {
	if strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(domain.User),
	)
}
