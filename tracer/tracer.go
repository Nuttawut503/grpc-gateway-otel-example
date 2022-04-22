package tracer

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var (
	_environment = "production"
	_url         = "http://localhost:14268/api/traces"
)

func NewTraceProvider(serviceName string) (*trace.TracerProvider, error) {
	// set env for this
	// OTEL_EXPORTER_JAEGER_AGENT_HOST = localhost
	// OTEL_EXPORTER_JAEGER_AGENT_PORT = 6831
	// OTEL_EXPORTER_JAEGER_ENDPOINT = http://localhost:14268/api/traces
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		return nil, err
	}
	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			attribute.String("environment", _environment),
		)),
	)
	otel.SetTracerProvider(provider)
	return provider, nil
}
