package xlog

import (
	"errors"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal/log"
	"github.com/pubgo/xlog/xlog_abc"
	"github.com/pubgo/xlog/xlog_config"
	"github.com/pubgo/xlog/xlog_opts"
	"go.uber.org/zap"
)

type Xlog = xlog_abc.Xlog
type M = xlog_abc.M

var logWatchers []func(log xlog_abc.Xlog)
var defaultLog xlog_abc.Xlog

func Watch(fn func(logs xlog_abc.Xlog)) {
	defer xerror.RespExit()

	xerror.Assert(fn == nil, "[fn] should not be nil")

	fn(getDefault())

	logWatchers = append(logWatchers, fn)
}

func init() {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	zapL := xerror.ExitErr(xlog_config.NewZapLogger(cfg)).(*zap.Logger)
	defaultLog = New(zapL).Named("", zap.WithCaller(true), zap.AddCallerSkip(1))
}

func New(zl *zap.Logger) xlog_abc.Xlog {
	xerror.Assert(zl == nil, "[xlog] [zl] should not be nil")
	return log.New().SetZapLogger(zl)
}

func getDefault() xlog_abc.Xlog {
	if defaultLog != nil {
		return defaultLog
	}

	xerror.Exit(errors.New("[xlog] please init defaultLog"))
	return nil
}

func getDefaultNext() xlog_abc.Xlog {
	return getDefault().Named("", xlog_opts.AddCallerSkip(1))
}

func SetDefault(lg xlog_abc.Xlog) (err error) {
	xerror.RespErr(&err)

	xerror.Assert(lg == nil, "[xlog] [lg] should not be nil")
	defaultLog = lg.Named("", zap.WithCaller(true), zap.AddCallerSkip(1))

	// 初始化log依赖
	for i := range logWatchers {
		logWatchers[i](defaultLog)
	}
	return nil
}

func With(fields ...zap.Field) xlog_abc.Xlog           { return getDefault().With(fields...) }
func Named(s string, opts ...zap.Option) xlog_abc.Xlog { return getDefault().Named(s, opts...) }
func Sync() error                                      { return xerror.Wrap(getDefault().Sync(), "[xlog] sync error") }

func Debug(fields ...interface{})  { getDefaultNext().Debug(fields...) }
func Info(fields ...interface{})   { getDefaultNext().Info(fields...) }
func Warn(fields ...interface{})   { getDefaultNext().Warn(fields...) }
func Error(fields ...interface{})  { getDefaultNext().Error(fields...) }
func DPanic(fields ...interface{}) { getDefaultNext().DPanic(fields...) }
func Panic(fields ...interface{})  { getDefaultNext().Panic(fields...) }
func Fatal(fields ...interface{})  { getDefaultNext().Fatal(fields...) }

func DebugM(msg string, m M)  { getDefaultNext().Debug(msg, m) }
func InfoM(msg string, m M)   { getDefaultNext().Info(msg, m) }
func WarnM(msg string, m M)   { getDefaultNext().Warn(msg, m) }
func ErrorM(msg string, m M)  { getDefaultNext().Error(msg, m) }
func DPanicM(msg string, m M) { getDefaultNext().DPanic(msg, m) }
func PanicM(msg string, m M)  { getDefaultNext().Panic(msg, m) }
func FatalM(msg string, m M)  { getDefaultNext().Fatal(msg, m) }

func Debugf(format string, a ...interface{})  { getDefaultNext().Debugf(format, a...) }
func Infof(format string, a ...interface{})   { getDefaultNext().Infof(format, a...) }
func Warnf(format string, a ...interface{})   { getDefaultNext().Warnf(format, a...) }
func Errorf(format string, a ...interface{})  { getDefaultNext().Errorf(format, a...) }
func DPanicf(format string, a ...interface{}) { getDefaultNext().DPanicf(format, a...) }
func Panicf(format string, a ...interface{})  { getDefaultNext().Panicf(format, a...) }
func Fatalf(format string, a ...interface{})  { getDefaultNext().Fatalf(format, a...) }

func DebugW(fn func(log xlog_abc.Logger))  { getDefault().DebugW(fn) }
func InfoW(fn func(log xlog_abc.Logger))   { getDefault().InfoW(fn) }
func WarnW(fn func(log xlog_abc.Logger))   { getDefault().WarnW(fn) }
func ErrorW(fn func(log xlog_abc.Logger))  { getDefault().ErrorW(fn) }
func DPanicW(fn func(log xlog_abc.Logger)) { getDefault().DPanicW(fn) }
func PanicW(fn func(log xlog_abc.Logger))  { getDefault().PanicW(fn) }
func FatalW(fn func(log xlog_abc.Logger))  { getDefault().FatalW(fn) }
