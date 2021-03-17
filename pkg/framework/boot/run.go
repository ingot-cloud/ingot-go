package boot

import (
	"context"

	"github.com/ingot-cloud/ingot-go/pkg/framework/container"
)

// Run 运行服务
func Run(context context.Context, factory container.Factory) error {
	app := &IngotApplication{
		Context: context,
		Factory: factory,
	}
	return app.Run()
}
