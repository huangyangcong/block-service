// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/huangyangcong/block-service/internal/biz"
	"github.com/huangyangcong/block-service/internal/conf"
	"github.com/huangyangcong/block-service/internal/data"
	"github.com/huangyangcong/block-service/internal/schedule"
	"github.com/huangyangcong/block-service/internal/server"
	"github.com/huangyangcong/block-service/internal/service"
	"github.com/huangyangcong/block-service/internal/util"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Trace, *conf.Registry, *conf.Email, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, schedule.ProviderSet, util.ProviderSet, newApp))
}
