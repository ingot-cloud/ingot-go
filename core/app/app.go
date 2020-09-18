package app

import (
	"context"
	"fmt"
	"ingot/common/log"
	"ingot/config"
	"ingot/core/injector"
	"ingot/core/server"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Run start app
func Run() error {
	banner()

	ctx := log.NewTagContext(context.Background(), "Application")

	cleanupFunc, err := initModule(ctx)
	if err != nil {
		return err
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.WithContext(ctx).Info("Server exiting")

	cleanupFunc()
	return nil
}

func banner() {
	fmt.Println("                      _                           _")
	fmt.Println("  ___    ___    ___  (_)  _ __     __ _    ___   | |_        ___    ___    _ __ ___")
	fmt.Println(" / __|  / _ \\  / __| | | | '_ \\   / _` |  / _ \\  | __|      / __|  / _ \\  | '_ ` _ \\")
	fmt.Println(" \\__ \\ |  __/ | (__  | | | | | | | (_| | | (_) | | |_   _  | (__  | (_) | | | | | | |")
	fmt.Println(" |___/  \\___|  \\___| |_| |_| |_|  \\__, |  \\___/   \\__| (_)  \\___|  \\___/  |_| |_| |_|")
	fmt.Println("                                  |___/")
	fmt.Println(":: Power by Motcwang ::")
}

func initModule(ctx context.Context) (func(), error) {

	loggerCleanFunc, err := log.InitLogger()
	if err != nil {
		return nil, err
	}

	container, containerCleanupFunc, err := injector.BuildContainer()
	if err != nil {
		return nil, err
	}

	httpServerCleanupFunc := runHTTPServer(ctx, config.CONFIG.Server, container.Engine)

	return func() {
		httpServerCleanupFunc()
		containerCleanupFunc()
		loggerCleanFunc()
	}, nil
}

func runHTTPServer(ctx context.Context, conf config.Server, handler http.Handler) func() {
	httpServer := server.HTTPServer(handler, conf)

	go func() {
		log.WithContext(ctx).Infof("=== HTTP server started successfully, address=%s ===", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithContext(ctx).Fatalf("listen: %s\n", err)
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(5*time.Second))
		defer cancel()

		httpServer.SetKeepAlivesEnabled(false)
		if err := httpServer.Shutdown(ctx); err != nil {
			log.WithContext(ctx).Errorf(err.Error())
		}
	}
}
