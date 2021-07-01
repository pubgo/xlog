package xlog

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

var loggerList []*xlog

var defaultZap *zap.Logger
var defaultLog Xlog
var defaultWLog Xlog

func init() {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	defaultZap = xerror.ExitErr(cfg.Build()).(*zap.Logger)
	defaultZap = defaultZap.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	defaultLog = &xlog{zl: defaultZap.WithOptions(zap.AddCallerSkip(1))}
	defaultWLog = defaultLog.Named("", zap.AddCallerSkip(-1))
}

func GetLogger(name string, opts ...zap.Option) Xlog {
	xerror.Assert(name == "", "[name] is null")

	var zl = defaultZap.Named(name).WithOptions(opts...)
	var xl = &xlog{opts: opts, name: name, zl: zl}
	loggerList = append(loggerList, xl)
	return xl
}

// SetDefault 设置默认的zap logger
func SetDefault(logger *zap.Logger) (err error) {
	xerror.RespErr(&err)

	xerror.Assert(logger == nil, "[logger] should not be nil")
	defaultZap = logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	defaultLog = &xlog{zl: defaultZap.WithOptions(zap.AddCallerSkip(1))}
	defaultWLog = defaultLog.Named("", zap.AddCallerSkip(-1))

	// 初始化log依赖
	for i := range loggerList {
		var xl = loggerList[i]
		xl.zl = defaultZap.Named(xl.name).WithOptions(xl.opts...)
		xl.initLogger()
	}

	return
}

func Debug(args ...interface{})               { defaultLog.Debug(args...) }
func Info(args ...interface{})                { defaultLog.Info(args...) }
func Warn(args ...interface{})                { defaultLog.Warn(args...) }
func Error(args ...interface{})               { defaultLog.Error(args...) }
func DPanic(args ...interface{})              { defaultLog.DPanic(args...) }
func Panic(args ...interface{})               { defaultLog.Panic(args...) }
func Fatal(args ...interface{})               { defaultLog.Fatal(args...) }
func DebugW(fn func(log Logger))              { defaultWLog.DebugW(fn) }
func InfoW(fn func(log Logger))               { defaultWLog.InfoW(fn) }
func WarnW(fn func(log Logger))               { defaultWLog.WarnW(fn) }
func ErrorW(fn func(log Logger))              { defaultWLog.ErrorW(fn) }
func DPanicW(fn func(log Logger))             { defaultWLog.DPanicW(fn) }
func PanicW(fn func(log Logger))              { defaultWLog.PanicW(fn) }
func FatalW(fn func(log Logger))              { defaultWLog.FatalW(fn) }
func Debugf(format string, a ...interface{})  { defaultLog.Debugf(format, a...) }
func Infof(format string, a ...interface{})   { defaultLog.Infof(format, a...) }
func Warnf(format string, a ...interface{})   { defaultLog.Warnf(format, a...) }
func Errorf(format string, a ...interface{})  { defaultLog.Errorf(format, a...) }
func Panicf(format string, a ...interface{})  { defaultLog.Panicf(format, a...) }
func Fatalf(format string, a ...interface{})  { defaultLog.Fatalf(format, a...) }
func DPanicf(format string, a ...interface{}) { defaultLog.DPanicf(format, a...) }
