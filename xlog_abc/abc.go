package xlog_abc

import (
	"go.uber.org/zap"
)

type M map[string]interface{}
type Xlog interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)

	DebugM(msg string, m M)
	InfoM(msg string, m M)
	WarnM(msg string, m M)
	ErrorM(msg string, m M)
	DPanicM(msg string, m M)
	PanicM(msg string, m M)
	FatalM(msg string, m M)

	Debugf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Panicf(format string, a ...interface{})
	Fatalf(format string, a ...interface{})
	DPanicf(format string, a ...interface{})

	Named(s string, opts ...zap.Option) Xlog
	With(fields ...zap.Field) Xlog
	Sync() error
}
