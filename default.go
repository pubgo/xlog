package xlog

import (
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
	"github.com/pubgo/xlog/xlog_config"
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

func getDefault() *log.XLog {
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
	getDefault().Debug(msg, fields...)
}

func Debugf(format string, a ...interface{}) {
	getDefault().Debug(fmt.Sprintf(format, a...))
}

func Infof(format string, a ...interface{}) {
	getDefault().Info(fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...interface{}) {
	getDefault().Warn(fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...interface{}) {
	getDefault().Error(fmt.Sprintf(format, a...))
}

func DPanicf(format string, a ...interface{}) {
	getDefault().DPanic(fmt.Sprintf(format, a...))
}

func Panicf(format string, a ...interface{}) {
	getDefault().Panic(fmt.Sprintf(format, a...))
}

func Fatalf(format string, a ...interface{}) {
	getDefault().Fatal(fmt.Sprintf(format, a...))
}

func Info(msg string, fields ...internal.Field) {
	getDefault().Info(msg, fields...)
}

func Warn(msg string, fields ...internal.Field) {
	getDefault().Warn(msg, fields...)
}

func Error(msg string, fields ...internal.Field) {
	getDefault().Error(msg, fields...)
}

func DPanic(msg string, fields ...internal.Field) {
	getDefault().DPanic(msg, fields...)
}

func Panic(msg string, fields ...internal.Field) {
	getDefault().Panic(msg, fields...)
}

func Fatal(msg string, fields ...internal.Field) {
	getDefault().Fatal(msg, fields...)
}

func With(fields ...zap.Field) internal.XLog {
	return getDefault().With(fields...)
}

func Named(s string, opts ...zap.Option) internal.XLog {
	return getDefault().Named(s, opts...)
}
