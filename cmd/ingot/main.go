package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ingot-cloud/ingot-go/internal/app"
	"github.com/ingot-cloud/ingot-go/pkg/framework/log"

	"github.com/urfave/cli/v2"
)

// VERSION 版本号
var VERSION = "0.1.0"

func main() {
	ctx := log.NewTagContext(context.Background(), "Application")

	app := &cli.App{
		Name:    "Ingot",
		Usage:   "Ingot",
		Version: VERSION,
		Commands: []*cli.Command{
			serverCmd(ctx),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.WithContext(ctx).Errorf(err.Error())
	}
}

func serverCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "Start service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "Configuration files(.yaml) in the 'config' directory",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "model",
				Aliases:  []string{"m"},
				Usage:    "Casbin model configuration file",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("config = %s", c.String("conf"))
			return app.Run(ctx, &app.Options{
				ConfigFile:      c.String("conf"),
				CasbinModelFile: c.String("model"),
			})
		},
	}
}
