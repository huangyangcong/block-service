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

func NewServices(hs *http.Server, gs *grpc.Server, m http.HandleOption,
	g *GreeterService) *Services {
	v1.RegisterGreeterServer(gs, g)
	hs.HandlePrefix("/", v1.NewGreeterHandler(g, m))
	return &Services{
		GreeterService: g,
	}
}
