package boot

import (
	"context"
	"os"
	"os/signal"

	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
)

// IngotApplication 应用入口
type IngotApplication struct {
	Context   context.Context
	Container *container.Container
}

// Run 运行Application
func (app *IngotApplication) Run() {
	banner()
	httpServer := &server.HTTPServer{
		Context: app.Context,
		Router:  app.Container.Router,
		Config:  app.Container.HTTPConfig,
	}
	clean := httpServer.Run()

	app.listeningSignal(clean)
}

func (app *IngotApplication) listeningSignal(doExit func()) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.WithContext(app.Context).Info("Server exiting")

	doExit()
}
