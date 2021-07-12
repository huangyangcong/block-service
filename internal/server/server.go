package server

import (
	"github.com/huangyangcong/block-service/internal/conf"

	"github.com/go-kratos/consul/registry"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewHandleOption, NewTracerProvider, NewRegister)

// Get trace provider
func NewTracerProvider(c *conf.Trace) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.NewRawExporter(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.ServiceNameKey.String("github.com/huangyangcong/block-service"),
			attribute.String("environment", "development"),
			attribute.Int64("ID", 1),
		)),
	)
	return tp, nil
}

// New Registry
func NewRegister(c *conf.Registry) (*registry.Registry, error) {
	// Get a new client
	config := api.DefaultConfig()
	if c.Url != "" {
		config.Address = c.Url
	}
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	r := registry.New(client)
	return r, nil
}
