package log_default

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	defaultLog = log.NewXLog(getDevLog())
}

var defaultLog internal.ILog

func GetLog() internal.ILog {
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

func GetDevLog() internal.ILog {
	return log.NewXLog(getDevLog())
}

