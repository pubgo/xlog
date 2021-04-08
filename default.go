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

	fn(getDefault().Named("", zap.AddCallerSkip(-1)))

	logWatchers = append(logWatchers, fn)
}

func init() {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	zapL := xerror.ExitErr(xlog_config.NewZapLogger(cfg)).(*zap.Logger)
	xerror.Exit(SetDefault(New(zapL)))
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

func SetDefault(lg xlog_abc.Xlog) (err error) {
	xerror.RespErr(&err)

	xerror.Assert(lg == nil, "[xlog] [zl] should not be nil")
	logW := lg.Named("", zap.WithCaller(true), zap.AddCallerSkip(1))
	defaultLog = logW.Named("", zap.AddCallerSkip(1))

	// 初始化log依赖
	for i := range logWatchers {
		logWatchers[i](logW)
	}
	return nil
}

func With(fields ...zap.Field) xlog_abc.Xlog           { return getDefault().With(fields...) }
func Named(s string, opts ...zap.Option) xlog_abc.Xlog { return getDefault().Named(s, opts...) }
func Sync() error                                      { return xerror.Wrap(getDefault().Sync(), "[xlog] sync error") }

func Debug(fields ...interface{})  { getDefault().Debug(fields...) }
func Info(fields ...interface{})   { getDefault().Info(fields...) }
func Warn(fields ...interface{})   { getDefault().Warn(fields...) }
func Error(fields ...interface{})  { getDefault().Error(fields...) }
func DPanic(fields ...interface{}) { getDefault().DPanic(fields...) }
func Panic(fields ...interface{})  { getDefault().Panic(fields...) }
func Fatal(fields ...interface{})  { getDefault().Fatal(fields...) }

func DebugM(msg string, m M)  { getDefault().Debug(msg, m) }
func InfoM(msg string, m M)   { getDefault().Info(msg, m) }
func WarnM(msg string, m M)   { getDefault().Warn(msg, m) }
func ErrorM(msg string, m M)  { getDefault().Error(msg, m) }
func DPanicM(msg string, m M) { getDefault().DPanic(msg, m) }
func PanicM(msg string, m M)  { getDefault().Panic(msg, m) }
func FatalM(msg string, m M)  { getDefault().Fatal(msg, m) }

func DebugW(fn func(log xlog_abc.Logger)) {
	getDefault().Named("", xlog_opts.AddCallerSkip(-1)).DebugW(fn)
}

func InfoW(fn func(log xlog_abc.Logger)) {
	getDefault().Named("", xlog_opts.AddCallerSkip(-1)).InfoW(fn)
}

func WarnW(fn func(log xlog_abc.Logger)) {
	getDefault().Named("", xlog_opts.AddCallerSkip(-1)).WarnW(fn)
}

func ErrorW(fn func(log xlog_abc.Logger)) {
	getDefault().Named("", xlog_opts.AddCallerSkip(-1)).ErrorW(fn)
}

func DPanicW(fn func(log xlog_abc.Logger)) {
	getDefault().Named("", xlog_opts.AddCallerSkip(-1)).DPanicW(fn)
}

func PanicW(fn func(log xlog_abc.Logger)) {
	getDefault().Named("", xlog_opts.AddCallerSkip(-1)).PanicW(fn)
}

func FatalW(fn func(log xlog_abc.Logger)) {
	getDefault().Named("", xlog_opts.AddCallerSkip(-1)).FatalW(fn)
}

func Debugf(format string, a ...interface{})  { getDefault().Debugf(format, a...) }
func Infof(format string, a ...interface{})   { getDefault().Infof(format, a...) }
func Warnf(format string, a ...interface{})   { getDefault().Warnf(format, a...) }
func Errorf(format string, a ...interface{})  { getDefault().Errorf(format, a...) }
func DPanicf(format string, a ...interface{}) { getDefault().DPanicf(format, a...) }
func Panicf(format string, a ...interface{})  { getDefault().Panicf(format, a...) }
func Fatalf(format string, a ...interface{})  { getDefault().Fatalf(format, a...) }
