package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/config"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/api"
	"github.com/ingot-cloud/ingot-go/pkg/framework/core/web/ingot"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// HTTPServer http web server
type HTTPServer struct {
	Context        context.Context
	Config         config.HTTPConfig
	HTTPConfigurer api.HTTPConfigurer
	Container      container.Container
}

// NewHTTPServer 创建 http 服务
func NewHTTPServer(context context.Context, c container.Container) *HTTPServer {
	return &HTTPServer{
		Context:        context,
		Config:         c.GetHTTPConfig(),
		HTTPConfigurer: c.GetHTTPConfigurer(),
		Container:      c,
	}
}

// Run 运行Http Web服务
func (server *HTTPServer) Run() func() {
	engine := server.buildHTTPHandler()
	return server.runHTTPServer(engine)
}

// BuildHTTPHandler to get gin.Engine
func (server *HTTPServer) buildHTTPHandler() *gin.Engine {
	gin.SetMode(server.Config.Mode)

	engine := gin.New()
	enableDefaultMiddleware(engine)
	enableSecurityMiddleware(engine, server.Container)

	server.defaultHTTPApi(engine)
	return engine
}

func (server *HTTPServer) defaultHTTPApi(engine *gin.Engine) {
	// 设置 prefix
	routerGroup := engine.Group(server.Config.Prefix)
	ingotRouter := ingot.NewRouter(routerGroup)
	server.HTTPConfigurer.Configure(ingotRouter)
	for _, api := range server.HTTPConfigurer.GetAPI() {
		api.Apply(ingotRouter)
	}
}

func (server *HTTPServer) runHTTPServer(handler http.Handler) func() {
	httpServer := &http.Server{
		Addr:         server.Config.Address,
		Handler:      handler,
		ReadTimeout:  server.Config.ReadTimeout * time.Second,
		WriteTimeout: server.Config.WriteTimeout * time.Second,
	}

	go func() {
		log.WithContext(server.Context).Infof("=== HTTP server started successfully, address=%s ===", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithContext(server.Context).Fatalf("listen: %s\n", err)
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(server.Context, time.Second*time.Duration(5*time.Second))
		defer cancel()

		httpServer.SetKeepAlivesEnabled(false)
		if err := httpServer.Shutdown(ctx); err != nil {
			log.WithContext(ctx).Errorf(err.Error())
		}
	}
}
