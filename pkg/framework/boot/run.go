package boot

import (
	"context"

	"github.com/ingot-cloud/ingot-go/pkg/framework/boot/container"
)

// Run 运行服务
func Run(context context.Context, container *container.Container, cleanFunc func()) {
	app := &IngotApplication{
		Context:   context,
		Container: container,
	}
	app.Run()
	cleanFunc()
}
