package util

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type SpanOption func(span opentracing.Span)

func SpanWithError(err error) SpanOption {
	return func(span opentracing.Span) {
		if err != nil {
			ext.Error.Set(span, true)
			span.LogFields(tlog.String("event", "error"), tlog.String("msg", err.Error()))
		}
	}
}

// example:
// SpanWithLog(
//    "event", "soft error",
//    "type", "cache timeout",
//    "waited.millis", 1500)
func SpanWithLog(arg ...interface{}) SpanOption {
	return func(span opentracing.Span) {
		span.LogKV(arg...)
	}
}

func Start(tracer opentracing.Tracer, spanName string, ctx context.Context) (newCtx context.Context, finish func(...SpanOption)) {
	if ctx == nil {
		ctx = context.TODO()
	}
	span, newCtx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, spanName,
		opentracing.Tag{Key: string(ext.Component), Value: "func"},
	)

	finish = func(ops ...SpanOption) {
		for _, o := range ops {
			o(span)
		}
		span.Finish()
	}

	return
}
