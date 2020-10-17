package opentracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	otracing "github.com/opentracing/opentracing-go"
)

// ContextWithSpan - Returns Context with Span
func ContextWithSpan(ctx context.Context, span otracing.Span) context.Context {
	return otracing.ContextWithSpan(ctx, span)
}

// SpanFromContext - Returns Span from Context
func SpanFromContext(ctx context.Context) otracing.Span {
	return otracing.SpanFromContext(ctx)
}

// StartSpanFromContext - Starts Span from Context
func StartSpanFromContext(ctx context.Context, operationName string, opts ...otracing.StartSpanOption) (otracing.Span, context.Context) {
	return otracing.StartSpanFromContext(ctx, operationName, opts...)
}

// CreateChildSpanFromContext - Creates Child Span from Context
// Extracts parent span from context
func CreateChildSpanFromContext(ctx context.Context, operation string) otracing.Span {
	parentSpan := SpanFromContext(ctx)

	return otracing.StartSpan(
		operation,
		opentracing.ChildOf(parentSpan.Context()),
	)
}
