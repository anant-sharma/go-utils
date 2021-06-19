package newrelictracing

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type key string

const (
	txnKey key = "transaction"
)

func NewTransaction(ctx context.Context, name string) (*newrelic.Transaction, context.Context) {
	txn, ok := ctx.Value(txnKey).(*newrelic.Transaction)
	if ok {
		txn.AddAttribute("NewTransactionName", name)
		return txn, ctx
	}

	newTxn := App.StartTransaction(name)
	return newTxn, context.WithValue(ctx, txnKey, newTxn)
}

func ContextWithTransaction(ctx context.Context, txn *newrelic.Transaction) context.Context {
	return context.WithValue(ctx, txnKey, txn)
}
