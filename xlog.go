package xlog

import (
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
	"go.uber.org/zap"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/xlog_config"
)

type XLog = internal.XLog
type Field = zap.Field

func GetDevLog() XLog {
	return log.GetDevLog()
}

func GetLog() XLog {
	return log.GetLog()
}

// 初始化加载
func init() {
	xerror.Exit(xlog_config.InitFromConfig(xlog_config.NewDevConfig()))
}

func FieldOf(fields ...Field) []Field {
	return fields
}

func Sync(ll internal.XLog) error {
	return log.Sync(ll)
}
