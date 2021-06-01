module block-service

go 1.15

require (
	github.com/go-kratos/consul v0.1.0
	github.com/go-kratos/kratos/v2 v2.0.0-rc1
	github.com/google/wire v0.5.0
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/consul/api v1.8.1
	github.com/kr/text v0.2.0 // indirect
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.20.0
	go.opentelemetry.io/otel/sdk v0.20.0
	golang.org/x/sys v0.0.0-20210525143221-35b2ab0089ea // indirect
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
