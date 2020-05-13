package typical

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"github.com/typical-go/typical-rest-server/server"
	"github.com/typical-go/typical-rest-server/server/config"
	"github.com/typical-go/typical-rest-server/server/repository"
	"github.com/typical-go/typical-rest-server/server/service"
)

func init() {
	typgo.Provide(
		&typgo.Constructor{Name: "", Fn: server.Connect},
		&typgo.Constructor{Name: "", Fn: repository.NewBookRepo},
		&typgo.Constructor{Name: "", Fn: service.NewBookService},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *typredis.Config, err error) {
				cfg = new(typredis.Config)
				if err = typgo.ProcessConfig("REDIS", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *typpostgres.Config, err error) {
				cfg = new(typpostgres.Config)
				if err = typgo.ProcessConfig("PG", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
		&typgo.Constructor{
			Name: "",
			Fn: func() (cfg *config.Config, err error) {
				cfg = new(config.Config)
				if err = typgo.ProcessConfig("APP", cfg); err != nil {
					return nil, err
				}
				return
			},
		},
	)
	typgo.Destroy(
		&typgo.Destructor{Fn: server.Disconnect},
	)
}
