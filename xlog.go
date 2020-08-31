package xlog

import (
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/internal/log"
)

type XLog = internal.XLog
type Field = zap.Field

func init() {
	xerror.Exit(xlog_config.InitDevLog())
}

func GetLog() internal.XLog {
	return defaultLog()
}

func Sync(ll XLog) (err error) {
	defer xerror.RespErr(&err)
	if ll == nil {
		return xerror.New("params is should not be nil")
	}

	xl, ok := ll.(*log.XLog)
	if !ok || xl == nil {
		return xerror.Fmt("params is should be log.XLog type, got(%v)", xl)
	}

	return xerror.Wrap(xl.Sync())
}

func defaultLog() internal.XLog {
	return log.GetLog()
}

var (
	Debug   = defaultLog().Debug
	DebugF  = defaultLog().DebugF
	Info    = defaultLog().Info
	InfoF   = defaultLog().InfoF
	Warn    = defaultLog().Warn
	WarnF   = defaultLog().WarnF
	Error   = defaultLog().Error
	ErrorF  = defaultLog().ErrorF
	DPanic  = defaultLog().DPanic
	DPanicF = defaultLog().DPanicF
	Panic   = defaultLog().Panic
	PanicF  = defaultLog().PanicF
	Named   = defaultLog().Named
	With    = defaultLog().With
	FatalF  = defaultLog().FatalF
	Fatal   = defaultLog().Fatal
)
