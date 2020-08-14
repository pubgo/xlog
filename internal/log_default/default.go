package log_default

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
)

func init() {
	defaultLog = log.NewXLog(getDevLog())
}

var defaultLog internal.XLog

func GetLog() internal.XLog {
	if defaultLog == nil {
		panic("please init default log config")
	}
	return defaultLog
}

func SetDefaultZapLog(lg *zap.Logger) {
	defaultLog = log.NewXLog(lg)
}

func getDevLog() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return xerror.PanicErr(cfg.Build()).(*zap.Logger)
}

func GetDevLog() internal.XLog {
	return log.NewXLog(getDevLog())
}
