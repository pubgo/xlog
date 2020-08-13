package internal

import (
	"go.uber.org/zap"
)

type Field = zap.Field

type ILog interface {
	Debug(msg string, fields ...Field)
	DebugF(format string, a ...interface{})
	Info(msg string, fields ...Field)
	InfoF(format string, a ...interface{})
	Warn(msg string, fields ...Field)
	WarnF(format string, a ...interface{})
	Error(msg string, fields ...Field)
	ErrorF(format string, a ...interface{})
	DPanic(msg string, fields ...Field)
	DPanicF(format string, a ...interface{})
	Panic(msg string, fields ...Field)
	PanicF(format string, a ...interface{})
	Fatal(msg string, fields ...Field)
	FatalF(format string, a ...interface{})
	Named(s string) ILog
	With(fields ...Field) ILog
}
