package xlog

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey struct{}

func WithCtx(ctx context.Context, fields ...zap.Field) context.Context {
	return context.WithValue(ctx, ctxKey{}, fields)
}
