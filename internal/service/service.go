package service

import (
	v1 "block-service/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewServices)

type Services struct {
	GreeterService *GreeterService
}

func NewServices(hs *http.Server, gs *grpc.Server, m http.ServerOption,
	g *GreeterService) *Services {
	v1.RegisterGreeterServer(gs, g)
	v1.RegisterGreeterHTTPServer(hs, g)
	return &Services{
		GreeterService: g,
	}
}
