package newrelictracing

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

const (
	segmentKey key = "segment"
)

func NewSegment(ctx context.Context, name string) (*newrelic.Segment, context.Context) {
	txn := ctx.Value(txnKey).(*newrelic.Transaction)
	if txn == nil {
		newTxn, newCtx := NewTransaction(ctx, name)
		ctx = newCtx
		txn = newTxn
	}

	segment := txn.StartSegment(name)
	return segment, context.WithValue(ctx, segmentKey, txn)
}

func ContextWithSegment(ctx context.Context, segment *newrelic.Segment) context.Context {
	return context.WithValue(ctx, segmentKey, segment)
}
