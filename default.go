package xlog

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

var loggerList []*xlog

var defaultZap *zap.Logger

func init() {
	cfg := xlog_config.NewDevConfig()
	cfg.EncoderConfig.EncodeCaller = "full"
	defaultZap = xerror.ExitErr(cfg.Build()).(*zap.Logger)
	defaultZap = defaultZap.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
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

	// 初始化log依赖
	for i := range loggerList {
		var xl = loggerList[i]
		xl.zl = defaultZap.Named(xl.name).WithOptions(xl.opts...)
		xl.initLogger()
	}

	return
}
