package xlog_abc

import (
	"go.uber.org/zap"
)

type Field = zap.Field
type Option = zap.Option
type Xlog interface {
	DebugFn(msg string, fn func(fields *Fields))
	InfoFn(msg string, fn func(fields *Fields))
	WarnFn(msg string, fn func(fields *Fields))
	ErrorFn(msg string, fn func(fields *Fields))
	DPanicFn(msg string, fn func(fields *Fields))
	PanicFn(msg string, fn func(fields *Fields))
	FatalFn(msg string, fn func(fields *Fields))

	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Warning(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	DPanic(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)

	Debugf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Warningf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Panicf(format string, a ...interface{})
	Fatalf(format string, a ...interface{})
	DPanicf(format string, a ...interface{})

	Named(s string, opts ...zap.Option) Xlog
	With(fields ...Field) Xlog
	Sync() error
}
