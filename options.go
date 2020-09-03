package xlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option = zap.Option

func Hooks(hooks ...func(zapcore.Entry) error) Option {
	return zap.Hooks(hooks...)
}

func Fields(fs ...Field) Option {
	return zap.Fields(fs...)
}

func ErrorOutput(w zapcore.WriteSyncer) Option {
	return zap.ErrorOutput(w)
}

func Development() Option {
	return zap.Development()
}

func AddCaller() Option {
	return zap.AddCaller()
}

func WithCaller(enabled bool) Option {
	return zap.WithCaller(enabled)
}

func AddCallerSkip(skip int) Option {
	return zap.AddCallerSkip(skip)
}

func AddStacktrace(lvl zapcore.LevelEnabler) Option {
	return zap.AddStacktrace(lvl)
}

func IncreaseLevel(lvl zapcore.LevelEnabler) Option {
	return zap.IncreaseLevel(lvl)
}
