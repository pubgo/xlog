package xlog

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey struct{}

func AppendCtx(ctx context.Context, fields ...zap.Field) context.Context {
	var fieldList, ok = ctx.Value(ctxKey{}).([]zap.Field)
	if !ok {
		fieldList = make([]zap.Field, 0, 3)
	}
	return context.WithValue(ctx, ctxKey{}, append(fieldList, fields...))
}
