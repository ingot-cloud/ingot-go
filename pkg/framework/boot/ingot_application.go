package boot

import (
	"context"
	"os"
	"os/signal"

	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
	"github.com/ingot-cloud/ingot-go/pkg/framework/security"
)

// IngotApplication 应用入口
type IngotApplication struct {
	Context context.Context
	Factory container.Factory
}

// Run 运行Application
func (app *IngotApplication) Run() error {
	banner()

	container, cleanFunc, err := app.Factory(app.Context)
	if err != nil {
		return nil
	}

	httpServer := &server.HTTPServer{
		Context: app.Context,
		Router:  container.Router,
		Config:  container.HTTPConfig,
		SecurityHandler: &security.Handler{
			Filter: container.Filter,
		},
	}
	clean := httpServer.Run()

	app.listeningSignal(clean)
	cleanFunc()
	return nil
}

func (app *IngotApplication) listeningSignal(doExit func()) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.WithContext(app.Context).Info("Server exiting")

	doExit()
}
