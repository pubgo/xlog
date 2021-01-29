package xlog

import (
	"errors"
	"fmt"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal/log"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

var logWatchers []func(log XLog)

func Watch(fn func(logs XLog)) {
	defer xerror.RespExit()

	xerror.Assert(fn == nil, "[fn] should not be nil")

	fn(GetDefault())

	logWatchers = append(logWatchers, fn)
}

func GetDevLog() XLog {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	zapL := xerror.PanicErr(xlog_config.NewZapLogger(cfg)).(*zap.Logger)
	devLog := New(zapL.WithOptions(WithCaller(true), WithCallerSkip(1)))
	return devLog.Named("xlog")
}

var defaultLog = GetDevLog().(*log.XLog)

func GetDefault() *log.XLog {
	if defaultLog != nil {
		return defaultLog
	}

	xerror.Exit(errors.New("please init defaultLog"))
	return nil
}

func SetDefault(zl XLog) (err error) {
	xerror.RespErr(&err)

	xerror.Assert(zl == nil, "the params should not be nil")

	defaultLog = zl.(*log.XLog)

	// 初始化log依赖
	for i := range logWatchers {
		logWatchers[i](defaultLog)
	}
	return nil
}

func Debug(msg string, fields ...Field)       { GetDefault().Debug(msg, fields...) }
func Debugf(format string, a ...interface{})  { GetDefault().Debug(fmt.Sprintf(format, a...)) }
func Infof(format string, a ...interface{})   { GetDefault().Info(fmt.Sprintf(format, a...)) }
func Warnf(format string, a ...interface{})   { GetDefault().Warn(fmt.Sprintf(format, a...)) }
func Errorf(format string, a ...interface{})  { GetDefault().Error(fmt.Sprintf(format, a...)) }
func DPanicf(format string, a ...interface{}) { GetDefault().DPanic(fmt.Sprintf(format, a...)) }
func Panicf(format string, a ...interface{})  { GetDefault().Panic(fmt.Sprintf(format, a...)) }
func Fatalf(format string, a ...interface{})  { GetDefault().Fatal(fmt.Sprintf(format, a...)) }
func Info(msg string, fields ...Field)        { GetDefault().Info(msg, fields...) }
func Warn(msg string, fields ...Field)        { GetDefault().Warn(msg, fields...) }
func Error(msg string, fields ...Field)       { GetDefault().Error(msg, fields...) }
func DPanic(msg string, fields ...Field)      { GetDefault().DPanic(msg, fields...) }
func Panic(msg string, fields ...Field)       { GetDefault().Panic(msg, fields...) }
func Fatal(msg string, fields ...Field)       { GetDefault().Fatal(msg, fields...) }
func With(fields ...zap.Field) XLog           { return GetDefault().With(fields...) }
func Named(s string, opts ...zap.Option) XLog { return GetDefault().Named(s, opts...) }
