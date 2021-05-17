package xlog

import "context"

type xlogCtx struct{}

func WithCtx(ctx context.Context, log Xlog) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	if log == nil {
		log = getDefault()
	}

	return context.WithValue(ctx, xlogCtx{}, log)
}

func FromCtx(ctx context.Context) Xlog {
	var val = ctx.Value(xlogCtx{})
	if val == nil {
		return getDefault()
	}

	return val.(Xlog)
}
