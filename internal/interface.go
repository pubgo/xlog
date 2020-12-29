package internal

import (
	"go.uber.org/zap"
	"log"
)

type Field = zap.Field
type Option = zap.Option
type XLog interface {
	Debug(msg string, fields ...Field)
	Debugf(format string, a ...interface{})
	Info(msg string, fields ...Field)
	Infof(format string, a ...interface{})
	Warn(msg string, fields ...Field)
	Warning(msg string, fields ...Field)
	Warningf(format string, a ...interface{})
	Error(msg string, fields ...Field)
	Errorf(format string, a ...interface{})
	DPanic(msg string, fields ...Field)
	DPanicf(format string, a ...interface{})
	Panic(msg string, fields ...Field)
	Panicf(format string, a ...interface{})
	Fatal(msg string, fields ...Field)
	Fatalf(format string, a ...interface{})
	Named(s string, opts ...zap.Option) XLog
	With(fields ...Field) XLog
}

func init() {
	log.Fatal()
}
