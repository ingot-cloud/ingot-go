package boot

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/server"
	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"
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

	clean := server.NewHTTPServer(app.Context, container).Run()

	app.listeningSignal(clean)
	cleanFunc()
	return nil
}

func (app *IngotApplication) listeningSignal(doExit func()) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.WithContext(app.Context).Info("Server exiting")

	doExit()
}
