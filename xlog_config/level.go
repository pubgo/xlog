package xlog_config

import "go.uber.org/zap/zapcore"

type Level = zapcore.Level

const (
	DebugLevel  Level = zapcore.DebugLevel
	InfoLevel         = zapcore.InfoLevel
	WarnLevel         = zapcore.WarnLevel
	ErrorLevel        = zapcore.ErrorLevel
	DPanicLevel       = zapcore.DPanicLevel
	PanicLevel        = zapcore.PanicLevel
	FatalLevel        = zapcore.FatalLevel
)
