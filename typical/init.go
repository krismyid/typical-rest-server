package typical

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/typical-go/typical-rest-server/app"
	"github.com/typical-go/typical-rest-server/app/controller"
	"github.com/typical-go/typical-rest-server/app/repository"
	"github.com/typical-go/typical-rest-server/config"
	"github.com/typical-go/typical-rest-server/experimental/typictx"
	"github.com/typical-go/typical-rest-server/typical/module"
)

// Context instance of Context
var Context typictx.Context

func init() {
	Context = typictx.Context{
		Name:        "Typical-RESTful-Server",
		Version:     "0.1.0",
		Description: "Example of typical and scalable RESTful API Server for Go",

		TypiApp: typictx.TypiApp{
			Config: &config.AppConfig{},
			ConfigLoadFunc: func() (config config.AppConfig, err error) {
				err = envconfig.Process(Context.ConfigPrefixOrDefault(), &config)
				return
			},
			Constructors: []interface{}{
				app.NewServer,
				controller.NewBookController,
				repository.NewBookRepository,
			},
			StartFunc: startApplication,
			StopFunc:  gracefulShutdown,
			TestTargets: []string{
				"./app/controller",
				"./app/repository",
			},
			MockTargets: []string{
				"./app/repository/book_repo.go",
			},
		},

		Modules: []*typictx.Module{
			module.NewPostgres(),
		},
	}

}
