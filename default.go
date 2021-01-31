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

type Xlog = internal.XLog

var logWatchers []func(log Xlog)

func Watch(fn func(logs Xlog)) {
	defer xerror.RespExit()

	xerror.Assert(fn == nil, "[fn] should not be nil")

	fn(getDefault())

	logWatchers = append(logWatchers, fn)
}

var defaultLog = func() *log.XLog {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	zapL := xerror.ExitErr(xlog_config.NewZapLogger(cfg)).(*zap.Logger)
	zl := log.New().SetZapLogger(zapL.WithOptions(WithCaller(true), WithCallerSkip(1)))
	return zl.Named("xlog").(*log.XLog)
}()

func New(zl *zap.Logger) Xlog {
	xerror.Assert(zl == nil, "[xlog] [zl] should not be nil")
	return log.New().SetZapLogger(zl)
}

func getDefault() *log.XLog {
	if defaultLog != nil {
		return defaultLog
	}

	xerror.Exit(errors.New("please init defaultLog"))
	return nil
}

func Init(zl *zap.Logger) (err error) {
	xerror.RespErr(&err)

	xerror.Assert(zl == nil, "[xlog] [zl] should not be nil")
	defaultLog = log.New().SetZapLogger(zl.WithOptions(WithCaller(true), WithCallerSkip(1)))
	// 初始化log依赖
	for i := range logWatchers {
		logWatchers[i](defaultLog)
	}
	return nil
}

type Fields = internal.Fields

func Debug(msg string, fields ...Field)       { getDefault().Debug(msg, fields...) }
func Info(msg string, fields ...Field)        { getDefault().Info(msg, fields...) }
func Warn(msg string, fields ...Field)        { getDefault().Warn(msg, fields...) }
func Error(msg string, fields ...Field)       { getDefault().Error(msg, fields...) }
func DPanic(msg string, fields ...Field)      { getDefault().DPanic(msg, fields...) }
func Panic(msg string, fields ...Field)       { getDefault().Panic(msg, fields...) }
func Fatal(msg string, fields ...Field)       { getDefault().Fatal(msg, fields...) }
func With(fields ...zap.Field) Xlog           { return getDefault().With(fields...) }
func Named(s string, opts ...zap.Option) Xlog { return getDefault().Named(s, opts...) }

func Debugf(format string, a ...interface{})  { getDefault().Debug(fmt.Sprintf(format, a...)) }
func Infof(format string, a ...interface{})   { getDefault().Info(fmt.Sprintf(format, a...)) }
func Warnf(format string, a ...interface{})   { getDefault().Warn(fmt.Sprintf(format, a...)) }
func Errorf(format string, a ...interface{})  { getDefault().Error(fmt.Sprintf(format, a...)) }
func DPanicf(format string, a ...interface{}) { getDefault().DPanic(fmt.Sprintf(format, a...)) }
func Panicf(format string, a ...interface{})  { getDefault().Panic(fmt.Sprintf(format, a...)) }
func Fatalf(format string, a ...interface{})  { getDefault().Fatal(fmt.Sprintf(format, a...)) }

func FatalFn(msg string, fn func(fields Fields))  { getDefault().FatalFn(msg, fn) }
func PanicFn(msg string, fn func(fields Fields))  { getDefault().PanicFn(msg, fn) }
func DPanicFn(msg string, fn func(fields Fields)) { getDefault().DPanicFn(msg, fn) }
func ErrorFn(msg string, fn func(fields Fields))  { getDefault().ErrorFn(msg, fn) }
func WarnFn(msg string, fn func(fields Fields))   { getDefault().WarnFn(msg, fn) }
func InfoFn(msg string, fn func(fields Fields))   { getDefault().InfoFn(msg, fn) }
func DebugFn(msg string, fn func(fields Fields))  { getDefault().DebugFn(msg, fn) }
