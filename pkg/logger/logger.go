package logger

import (
	"context"

	"go.uber.org/zap"

	"github.com/amitmahbubani/protologs/pkg/protolog"
)

const CtxKey = "ctxLogger"

func New(ctx context.Context) (context.Context, *zap.Logger) {
	// Create a default zap production logger
	logger, _ := zap.NewProduction()

	// Instantiate a context struct
	ctxField := &protolog.Context{
		Request: &protolog.Request{
			Id:         "1234567",
			Source:     "sourcevalue",
			Rpcservice: "servicev",
			Rpcpackage: "packagev",
			Rpcmethod:  "rpcv",
		},
	}

	// Add some meta. These fields are not indexed.
	meta := make([]*protolog.Meta, 0)
	meta = append(meta, &protolog.Meta{Key: "meta1", Value: "one"})
	meta = append(meta, &protolog.Meta{Key: "meta2", Value: "two"})
	ctxField.Meta = meta

	// Create a child zap logger to set to context
	logger = logger.With(zap.Object("context", ctxField))
	logger.Info("logger initialized")

	ctx = context.WithValue(ctx, CtxKey, logger)
	return ctx, logger
}

func Ctx(ctx context.Context) *zap.Logger {
	k, ok := ctx.Value(CtxKey).(*zap.Logger)
	if ok {
		return k
	}

	_, logger := New(ctx)
	return logger
}
