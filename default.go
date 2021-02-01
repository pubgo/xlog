package xlog

import (
	"errors"
	"fmt"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal/log"
	"github.com/pubgo/xlog/xlog_abc"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

var logWatchers []func(log xlog_abc.Xlog)

func Watch(fn func(logs xlog_abc.Xlog)) {
	defer xerror.RespExit()

	xerror.Assert(fn == nil, "[fn] should not be nil")

	fn(getDefault().Named("", zap.AddCallerSkip(-1)))

	logWatchers = append(logWatchers, fn)
}

func init() {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	zapL := xerror.ExitErr(xlog_config.NewZapLogger(cfg)).(*zap.Logger)
	xerror.Exit(Init(zapL))
}

var defaultLog xlog_abc.Xlog

func New(zl *zap.Logger) xlog_abc.Xlog {
	xerror.Assert(zl == nil, "[xlog] [zl] should not be nil")
	return log.New().SetZapLogger(zl)
}

func getDefault() xlog_abc.Xlog {
	if defaultLog != nil {
		return defaultLog
	}

	xerror.Exit(errors.New("please init defaultLog"))
	return nil
}

func Init(zl *zap.Logger) (err error) {
	xerror.RespErr(&err)

	xerror.Assert(zl == nil, "[xlog] [zl] should not be nil")
	logW := log.New().SetZapLogger(zl.WithOptions(zap.WithCaller(true), zap.AddCallerSkip(1)))
	defaultLog = logW.Named("", zap.AddCallerSkip(1))

	// 初始化log依赖
	for i := range logWatchers {
		logWatchers[i](logW)
	}
	return nil
}

func Debug(msg string, fields ...zap.Field)            { getDefault().Debug(msg, fields...) }
func Info(msg string, fields ...zap.Field)             { getDefault().Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)             { getDefault().Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field)            { getDefault().Error(msg, fields...) }
func DPanic(msg string, fields ...zap.Field)           { getDefault().DPanic(msg, fields...) }
func Panic(msg string, fields ...zap.Field)            { getDefault().Panic(msg, fields...) }
func Fatal(msg string, fields ...zap.Field)            { getDefault().Fatal(msg, fields...) }
func With(fields ...zap.Field) xlog_abc.Xlog           { return getDefault().With(fields...) }
func Named(s string, opts ...zap.Option) xlog_abc.Xlog { return getDefault().Named(s, opts...) }

func Debugf(format string, a ...interface{})  { getDefault().Debug(fmt.Sprintf(format, a...)) }
func Infof(format string, a ...interface{})   { getDefault().Info(fmt.Sprintf(format, a...)) }
func Warnf(format string, a ...interface{})   { getDefault().Warn(fmt.Sprintf(format, a...)) }
func Errorf(format string, a ...interface{})  { getDefault().Error(fmt.Sprintf(format, a...)) }
func DPanicf(format string, a ...interface{}) { getDefault().DPanic(fmt.Sprintf(format, a...)) }
func Panicf(format string, a ...interface{})  { getDefault().Panic(fmt.Sprintf(format, a...)) }
func Fatalf(format string, a ...interface{})  { getDefault().Fatal(fmt.Sprintf(format, a...)) }

func FatalFn(msg string, fn func(fields *xlog_abc.Fields))  { getDefault().FatalFn(msg, fn) }
func PanicFn(msg string, fn func(fields *xlog_abc.Fields))  { getDefault().PanicFn(msg, fn) }
func DPanicFn(msg string, fn func(fields *xlog_abc.Fields)) { getDefault().DPanicFn(msg, fn) }
func ErrorFn(msg string, fn func(fields *xlog_abc.Fields))  { getDefault().ErrorFn(msg, fn) }
func WarnFn(msg string, fn func(fields *xlog_abc.Fields))   { getDefault().WarnFn(msg, fn) }
func InfoFn(msg string, fn func(fields *xlog_abc.Fields))   { getDefault().InfoFn(msg, fn) }
func DebugFn(msg string, fn func(fields *xlog_abc.Fields))  { getDefault().DebugFn(msg, fn) }
