package xlog

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey struct{}

func AppendCtx(ctx context.Context, fields ...zap.Field) context.Context {
	var fieldList, _ = ctx.Value(ctxKey{}).([]zap.Field)
	return context.WithValue(ctx, ctxKey{}, append(fieldList, fields...))
}
