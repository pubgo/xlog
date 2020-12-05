package xlog

import (
	"errors"
	"fmt"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

func GetDevLog() XLog {
	zl, err := xlog_config.NewZapLoggerFromConfig(xlog_config.NewDevConfig())
	xerror.Exit(err)

	zl = zl.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1)).Named("xlog")
	return New(zl)
}

var defaultLog = func() *log.XLog {
	return GetDevLog().Named("debug", zap.AddCallerSkip(1)).(*log.XLog)
}()

func GetDefault() *log.XLog {
	if defaultLog != nil {
		return defaultLog
	}

	xerror.Exit(errors.New("please init defaultLog"))
	return nil
}

func SetDefault(zl XLog) error {
	if zl == nil {
		return xerror.New("the params should not be nil")
	}

	defaultLog = zl.(*log.XLog)
	return nil
}

func Debug(msg string, fields ...internal.Field) {
	GetDefault().Debug(msg, fields...)
}

func Debugf(format string, a ...interface{}) {
	GetDefault().Debug(fmt.Sprintf(format, a...))
}

func Infof(format string, a ...interface{}) {
	GetDefault().Info(fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...interface{}) {
	GetDefault().Warn(fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...interface{}) {
	GetDefault().Error(fmt.Sprintf(format, a...))
}

func DPanicf(format string, a ...interface{}) {
	GetDefault().DPanic(fmt.Sprintf(format, a...))
}

func Panicf(format string, a ...interface{}) {
	GetDefault().Panic(fmt.Sprintf(format, a...))
}

func Fatalf(format string, a ...interface{}) {
	GetDefault().Fatal(fmt.Sprintf(format, a...))
}

func Info(msg string, fields ...internal.Field) {
	GetDefault().Info(msg, fields...)
}

func Warn(msg string, fields ...internal.Field) {
	GetDefault().Warn(msg, fields...)
}

func Error(msg string, fields ...internal.Field) {
	GetDefault().Error(msg, fields...)
}

func DPanic(msg string, fields ...internal.Field) {
	GetDefault().DPanic(msg, fields...)
}

func Panic(msg string, fields ...internal.Field) {
	GetDefault().Panic(msg, fields...)
}

func Fatal(msg string, fields ...internal.Field) {
	GetDefault().Fatal(msg, fields...)
}

func With(fields ...zap.Field) internal.XLog {
	return GetDefault().With(fields...)
}

func Named(s string, opts ...zap.Option) internal.XLog {
	return GetDefault().Named(s, opts...)
}
