package xlog

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
	"go.uber.org/zap"
)

type XLog = internal.XLog

func New(zl *zap.Logger) XLog {
	xerror.Assert(zl == nil, "zap.Logger [zl] should not be nil")
	return log.NewXLog().SetZapLogger(zl)
}

func Sync(logs ...XLog) (err error) {
	defer xerror.RespErr(&err)

	if len(logs) == 0 {
		logs = append(logs, defaultLog)
	}

	for i := range logs {
		xl, ok := logs[i].(*log.XLog)
		xerror.Assert(!ok || xl == nil, "[xl] should not be nil")
		xerror.PanicF(xl.Sync(), "[xlog] %#v sync error", xl)
	}

	return nil
}
