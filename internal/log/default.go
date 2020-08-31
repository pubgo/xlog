package log

import (
	"github.com/pubgo/xlog/internal"
	"go.uber.org/zap"
)

var defaultLog = &xlog{}

func GetLog() internal.XLog {
	if defaultLog == nil {
		panic("please init default log config")
	}
	return defaultLog
}

func SetDefaultZapLog(zl *zap.Logger) {
	defaultLog.zl = zl
}
