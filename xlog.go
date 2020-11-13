package xlog

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
	"go.uber.org/zap"
)

type XLog = internal.XLog

func New(zl *zap.Logger) XLog {
	if zl == nil {
		xerror.Panic(xerror.New("zap.Logger [zl] should not be nil"))
	}
	return log.NewXLog().SetZapLogger(zl)
}

func Sync(ll ...XLog) (err error) {
	defer xerror.RespErr(&err)
	if len(ll) == 0 {
		ll = append(ll, defaultLog)
	}

	for i := range ll {
		xl, ok := ll[i].(*log.XLog)
		if !ok || xl == nil {
			return xerror.Fmt("the params should be log.XLog type, got(%v)", xl)
		}

		xerror.Panic(xl.Sync())
	}

	return nil
}
