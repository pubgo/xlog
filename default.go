package xlog

import (
	"fmt"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

func GetDevLog() XLog {
	zl, err := xlog_config.NewZapLoggerFromConfig(xlog_config.NewDevConfig())
	if err != nil {
		xerror.Exit(err)
	}

	zl = zl.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1)).Named("xlog")
	return log.NewXLog().SetZapLogger(zl)
}

var defaultLog = func() *log.XLog {
	return GetDevLog().Named("debug", zap.AddCallerSkip(1)).(*log.XLog)
}()

func SetLog(zl XLog) error {
	if zl == nil {
		return xerror.New("the params should not be nil")
	}

	defaultLog = zl.(*log.XLog)
	return nil
}

func Debug(msg string, fields ...internal.Field) {
	defaultLog.Debug(msg, fields...)
}

func DebugF(format string, a ...interface{}) {
	defaultLog.Debug(fmt.Sprintf(format, a...))
}

func InfoF(format string, a ...interface{}) {
	defaultLog.Info(fmt.Sprintf(format, a...))
}

func WarnF(format string, a ...interface{}) {
	defaultLog.Warn(fmt.Sprintf(format, a...))
}

func ErrorF(format string, a ...interface{}) {
	defaultLog.Error(fmt.Sprintf(format, a...))
}

func DPanicF(format string, a ...interface{}) {
	defaultLog.DPanic(fmt.Sprintf(format, a...))
}

func PanicF(format string, a ...interface{}) {
	defaultLog.Panic(fmt.Sprintf(format, a...))
}

func FatalF(format string, a ...interface{}) {
	defaultLog.Fatal(fmt.Sprintf(format, a...))
}

func Info(msg string, fields ...internal.Field) {
	defaultLog.Info(msg, fields...)
}

func Warn(msg string, fields ...internal.Field) {
	defaultLog.Warn(msg, fields...)
}

func Error(msg string, fields ...internal.Field) {
	defaultLog.Error(msg, fields...)
}

func DPanic(msg string, fields ...internal.Field) {
	defaultLog.DPanic(msg, fields...)
}

func Panic(msg string, fields ...internal.Field) {
	defaultLog.Panic(msg, fields...)
}

func Fatal(msg string, fields ...internal.Field) {
	defaultLog.Fatal(msg, fields...)
}

func With(fields ...zap.Field) internal.XLog {
	return defaultLog.With(fields...)
}

func Named(s string, opts ...zap.Option) internal.XLog {
	return defaultLog.Named(s, opts...)
}
