package app

import (
	"context"

	"github.com/ingot-cloud/ingot-go/internal/app/config"
	"github.com/ingot-cloud/ingot-go/internal/app/container"
	"github.com/ingot-cloud/ingot-go/internal/app/container/injector"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot"
	bootContainer "github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	securityContainer "github.com/ingot-cloud/ingot-go/pkg/framework/security/container"
)

// Run start app
func Run(ctx context.Context, options *config.Options) error {
	// 初始化模块
	factory := func(ctx context.Context) (*bootContainer.Container, func(), error) {
		return initContainer(ctx, options)
	}

	return boot.Run(ctx, factory)
}

func initContainer(ctx context.Context, options *config.Options) (*bootContainer.Container, func(), error) {

	// 初始化 config
	config, err := injector.BuildConfiguration(options)
	if err != nil {
		return nil, nil, err
	}

	// 初始化 log
	loggerCleanFunc, err := log.InitLogger(config.Log)
	if err != nil {
		return nil, nil, err
	}

	// AppContainer
	appContainer, appCleanup, err := injector.BuildAppContainer(config, options)
	if err != nil {
		return nil, nil, err
	}

	// 安全容器
	securityAllContainer, err := createSecurityAllContainer(appContainer, config)
	if err != nil {
		return nil, nil, err
	}

	// boot 容器
	container, containerCleanupFunc, err := injector.BuildContainer(appContainer, securityAllContainer)
	if err != nil {
		return nil, nil, err
	}

	return container, func() {
		containerCleanupFunc()
		appCleanup()
		loggerCleanFunc()
	}, nil
}

func createSecurityAllContainer(app *container.AppContainer, config *config.Config) (securityContainer.SecurityAllContainer, error) {
	allContainer := &securityContainer.DefaultSecurityAllContainer{}

	securityContainer, err := injector.BuildSecurityContainer(app)
	if err != nil {
		return nil, err
	}
	allContainer.SecurityContainer = securityContainer

	oauth2Container, err := injector.BuildOAuth2Container(config.Security.OAuth2, app)
	if err != nil {
		return nil, err
	}
	allContainer.OAuth2Container = oauth2Container

	if config.Security.EnableResourceServer {
		resourceServerContainer, err := injector.BuildResourceServerContainer(oauth2Container, app)
		if err != nil {
			return nil, err
		}
		allContainer.ResourceServerContainer = resourceServerContainer
	}
	if config.Security.EnableAuthorizationServer {
		authorizationServerContainer, err := injector.BuildAuthorizationServerContainer(oauth2Container, securityContainer, app)
		if err != nil {
			return nil, err
		}
		allContainer.AuthorizationServerContainer = authorizationServerContainer
	}

	return allContainer, nil
}
