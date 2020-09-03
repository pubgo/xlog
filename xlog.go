package xlog

import (
	"go.uber.org/zap"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
)

type XLog = internal.XLog

func New(zl *zap.Logger) XLog {
	if zl == nil {
		panic("zap.Logger should not be nil")
	}
	return log.NewXLog().SetZapLogger(zl)
}

func Sync(ll XLog) (err error) {
	defer xerror.RespErr(&err)
	if ll == nil {
		return xerror.New("the params should not be nil")
	}

	xl, ok := ll.(*log.XLog)
	if !ok || xl == nil {
		return xerror.Fmt("the params should be log.XLog type, got(%v)", xl)
	}

	return xerror.Wrap(xl.Sync())
}
