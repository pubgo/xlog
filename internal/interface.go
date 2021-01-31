package internal

import (
	"go.uber.org/zap"
)

type Field = zap.Field
type Option = zap.Option
type XLog interface {
	Debug(msg string, fields ...Field)
	DebugFn(msg string, fn func(fields Fields))
	Debugf(format string, a ...interface{})
	Info(msg string, fields ...Field)
	InfoFn(msg string, fn func(fields Fields))
	Infof(format string, a ...interface{})
	Warn(msg string, fields ...Field)
	WarnFn(msg string, fn func(fields Fields))
	Warning(msg string, fields ...Field)
	Warningf(format string, a ...interface{})
	Error(msg string, fields ...Field)
	ErrorFn(msg string, fn func(fields Fields))
	Errorf(format string, a ...interface{})
	DPanic(msg string, fields ...Field)
	DPanicFn(msg string, fn func(fields Fields))
	DPanicf(format string, a ...interface{})
	Panic(msg string, fields ...Field)
	PanicFn(msg string, fn func(fields Fields))
	Panicf(format string, a ...interface{})
	Fatal(msg string, fields ...Field)
	FatalFn(msg string, fn func(fields Fields))
	Fatalf(format string, a ...interface{})
	Named(s string, opts ...zap.Option) XLog
	With(fields ...Field) XLog
	Sync() error
}
