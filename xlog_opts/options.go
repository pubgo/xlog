package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option = zap.Option

func WrapCore(f func(zapcore.Core) zapcore.Core) Option { return zap.WrapCore(f) }
func Hooks(hooks ...func(zapcore.Entry) error) Option   { return zap.Hooks(hooks...) }
func Fields(fs ...zap.Field) Option                     { return zap.Fields(fs...) }
func ErrorOutput(w zapcore.WriteSyncer) Option          { return zap.ErrorOutput(w) }
func Development() Option                               { return zap.Development() }
func AddCaller() Option                                 { return zap.AddCaller() }
func WithCaller(enabled bool) Option                    { return zap.WithCaller(enabled) }
func AddCallerSkip(skip int) Option                     { return zap.AddCallerSkip(skip) }
func AddStacktrace(lvl zapcore.LevelEnabler) Option     { return zap.AddStacktrace(lvl) }
func IncreaseLevel(lvl zapcore.LevelEnabler) Option     { return zap.IncreaseLevel(lvl) }
func OnFatal(action zapcore.CheckWriteAction) Option    { return zap.OnFatal(action) }
