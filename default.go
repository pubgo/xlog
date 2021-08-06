package xlog

import (
	"log"
	"sync"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

var loggerMap sync.Map
var defaultLog Xlog
var defaultWLog Xlog

func init() {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	var zapLog = xerror.ExitErr(cfg.Build()).(*zap.Logger)
	xerror.Exit(SetDefault(zapLog))
}

func GetLogger(name string, opts ...zap.Option) Xlog {
	xerror.Assert(name == "", "[name] is null")

	if xl, ok := loggerMap.Load(name); ok {
		return xl.(*xlog)
	}

	var xl = &xlog{opts: opts, zl: zap.L().
		Named(name).
		WithOptions(zap.AddCallerSkip(1)).
		WithOptions(opts...)}
	loggerMap.Store(name, xl)
	return xl
}

// GetDefault 获取默认xlog
func GetDefault() Xlog { return defaultLog }

// SetDefault 设置默认的zap logger
func SetDefault(logger *zap.Logger) (err error) {
	xerror.RespErr(&err)

	xerror.Assert(logger == nil, "[logger] should not be nil")

	// 替换zap默认log
	zap.ReplaceGlobals(logger)

	// 替换std默认log
	var stdLog = log.Default()
	*stdLog = *zap.NewStdLog(logger)

	defaultWLog = &xlog{zl: logger.WithOptions(zap.AddCallerSkip(1))}
	defaultLog = &xlog{zl: logger.WithOptions(zap.AddCallerSkip(2))}

	// 初始化log依赖
	loggerMap.Range(func(name, xl interface{}) bool {
		xl.(*xlog).zl = zap.L().
			Named(name.(string)).
			WithOptions(zap.AddCallerSkip(1)).
			WithOptions(xl.(*xlog).opts...)

		return true
	})

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
