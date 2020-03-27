package typical

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
	"github.com/typical-go/typical-rest-server/pkg/typredis"
	"github.com/typical-go/typical-rest-server/server/config"
	"github.com/typical-go/typical-rest-server/server/repository"
	"github.com/typical-go/typical-rest-server/server/service"
)

func init() {
	typapp.AppendConstructor(
		typapp.NewConstructor(repository.NewBookRepo),
		typapp.NewConstructor(service.NewBookService),
		typapp.NewConstructor(func(cfgMngr typcore.ConfigManager) (*config.Config, error) {
			cfg, err := cfgMngr.RetrieveConfig("APP")
			if err != nil {
				return nil, err
			}
			return cfg.(*config.Config), nil
		}),
		typapp.NewConstructor(func(cfgMngr typcore.ConfigManager) (*typpostgres.Config, error) {
			cfg, err := cfgMngr.RetrieveConfig("PG")
			if err != nil {
				return nil, err
			}
			return cfg.(*typpostgres.Config), nil
		}),
		typapp.NewConstructor(func(cfgMngr typcore.ConfigManager) (*typredis.Config, error) {
			cfg, err := cfgMngr.RetrieveConfig("REDIS")
			if err != nil {
				return nil, err
			}
			return cfg.(*typredis.Config), nil
		}),
	)
}
