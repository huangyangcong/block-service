// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"block-service/internal/biz"
	"block-service/internal/conf"
	"block-service/internal/data"
	"block-service/internal/schedule"
	"block-service/internal/server"
	"block-service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, trace *conf.Trace, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	registryRegistry, err := server.NewRegister(registry)
	if err != nil {
		return nil, nil, err
	}
	httpServer := server.NewHTTPServer(confServer)
	tracerProvider, err := server.NewTracerProvider(trace)
	if err != nil {
		return nil, nil, err
	}
	grpcServer := server.NewGRPCServer(tracerProvider, confServer, logger)
	handleOption := server.NewHandleOption(logger)
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase, logger)
	services := service.NewServices(httpServer, grpcServer, handleOption, greeterService)
	cron := schedule.NewSchedule(logger)
	boxPrice := schedule.NewBoxPrice(cron, logger)
	routes := schedule.NewScheduleRoutes(boxPrice)
	app := newApp(logger, registryRegistry, services, routes)
	return app, func() {
		cleanup()
	}, nil
}
