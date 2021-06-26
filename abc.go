package xlog

import (
	"go.uber.org/zap"
)

type M map[string]interface{}
type Xlog interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	DebugW(func(log Logger))
	InfoW(func(log Logger))
	WarnW(func(log Logger))
	ErrorW(func(log Logger))
	DPanicW(func(log Logger))
	PanicW(func(log Logger))
	FatalW(func(log Logger))

	Debugf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Panicf(format string, a ...interface{})
	Fatalf(format string, a ...interface{})
	DPanicf(format string, a ...interface{})

	Named(name string, opts ...zap.Option) Xlog
	Zap() *zap.Logger
}

type Logger interface {
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
	PrintM(m M)
}
