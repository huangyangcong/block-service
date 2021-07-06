// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"block-service/internal/biz"
	"block-service/internal/conf"
	"block-service/internal/data"
	"block-service/internal/schedule"
	"block-service/internal/server"
	"block-service/internal/service"
	"block-service/internal/util"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Trace, *conf.Registry, *conf.Email, log.Logger) *kratos.App {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, schedule.ProviderSet, util.ProviderSet, newApp))
}
