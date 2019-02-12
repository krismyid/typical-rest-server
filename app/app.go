package app

import (
	"github.com/imantung/typical-go-server/config"
	"github.com/imantung/typical-go-server/db"
	"github.com/urfave/cli"
	"go.uber.org/dig"
)

func container() *dig.Container {
	container := dig.New()
	container.Provide(config.NewConfig)
	container.Provide(db.Connect)
	container.Provide(newServer)
	return container
}

func triggerAction(invokeFunc interface{}) interface{} {
	return func(ctx *cli.Context) error {
		container := container()
		container.Provide(ctx.Args)
		return container.Invoke(invokeFunc)
	}
}
