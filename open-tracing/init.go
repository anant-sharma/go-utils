package opentracing

import (
	"context"
	"io"
	"time"

	otracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// GlobalTracer - Returns Instance of Global Tracer
func GlobalTracer() otracing.Tracer {
	return otracing.GlobalTracer()
}

// NewTracer - Creates New Tracer
func NewTracer(ctx context.Context, serviceName string) (otracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			CollectorEndpoint:   "http://instance.chipserver.ml:14268/api/traces",
		},
	}
	return cfg.New(
		serviceName,
		config.Logger(jaeger.StdLogger),
	)
}

// Init - Initialize Global Tracing
func Init(ctx context.Context, serviceName string) error {
	tracer, _, err := NewTracer(ctx, serviceName)

	if err != nil {
		return err
	}

	otracing.SetGlobalTracer(tracer)
	return nil
}
