// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/ingot-cloud/ingot-go/internal/app/api"
	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/core/http"
	"github.com/ingot-cloud/ingot-go/internal/app/model/dao"
	"github.com/ingot-cloud/ingot-go/internal/app/provider"
	"github.com/ingot-cloud/ingot-go/internal/app/provider/factory"
	"github.com/ingot-cloud/ingot-go/internal/app/service"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	config2 "github.com/ingot-cloud/ingot-go/pkg/framework/security/web/config"
)

// Injectors from wire.go:

func BuildConfiguration(options *config.Options) (*config.Config, error) {
	configConfig, err := provider.NewConfig(options)
	if err != nil {
		return nil, err
	}
	return configConfig, nil
}

func BuildContainer(config3 *config.Config, options *config.Options) (*container.Container, func(), error) {
	httpConfig, err := factory.HTTPConfigSet(config3)
	if err != nil {
		return nil, nil, err
	}
	authentication, cleanup, err := factory.NewAuthentication(config3)
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := factory.NewGorm(config3)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	role := &dao.Role{
		DB: db,
	}
	roleAuthority := &dao.RoleAuthority{
		DB: db,
	}
	authority := &dao.Authority{
		DB: db,
	}
	user := &dao.User{
		DB: db,
	}
	roleUser := &dao.RoleUser{
		DB: db,
	}
	permission := &service.Permission{
		RoleDao:          role,
		RoleAuthorityDao: roleAuthority,
		AuthorityDao:     authority,
		UserDao:          user,
		RoleUserDao:      roleUser,
	}
	casbinAdapterService := &service.CasbinAdapterService{
		PermissionService: permission,
	}
	syncedEnforcer, cleanup3, err := factory.NewCasbin(options, casbinAdapterService)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	auth, err := factory.AuthConfigSet(config3)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	encoder, cleanup4, err := factory.NewPasswordEncoder()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceAuth := &service.Auth{
		UserDao:         user,
		RoleUserDao:     roleUser,
		RoleDao:         role,
		Auth:            authentication,
		PasswordEncoder: encoder,
	}
	apiAuth := &api.Auth{
		AuthService: serviceAuth,
	}
	apiConfig := &http.APIConfig{
		Auth:           authentication,
		CasbinEnforcer: syncedEnforcer,
		HTTPConfig:     httpConfig,
		AuthConfig:     auth,
		AuthAPI:        apiAuth,
	}
	webSecurityConfigurers, err := provider.NewWebSecurityConfigurers()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	filter, err := config2.BuildWebSecurityFilter(webSecurityConfigurers)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	containerContainer := &container.Container{
		HTTPConfig:     httpConfig,
		HTTPConfigurer: apiConfig,
		Filter:         filter,
	}
	return containerContainer, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
