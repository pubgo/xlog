package xlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option = zap.Option

func WithHooks(hooks ...func(zapcore.Entry) error) Option { return zap.Hooks(hooks...) }
func WithFields(fs ...Field) Option                       { return zap.Fields(fs...) }
func WithErrorOutput(w zapcore.WriteSyncer) Option        { return zap.ErrorOutput(w) }
func WithDevelopment() Option                             { return zap.Development() }
func WithCaller(enabled bool) Option                      { return zap.WithCaller(enabled) }
func WithCallerSkip(skip int) Option                      { return zap.AddCallerSkip(skip) }
func WithStacktrace(lvl zapcore.LevelEnabler) Option      { return zap.AddStacktrace(lvl) }
func WithIncreaseLevel(lvl zapcore.LevelEnabler) Option   { return zap.IncreaseLevel(lvl) }
func WithOnFatal(action zapcore.CheckWriteAction) Option  { return zap.OnFatal(action) }
