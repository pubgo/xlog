package xlog

import (
	"go.uber.org/zap"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal/log"
	"github.com/pubgo/xlog/xlog_config"
)

var defaultLog = func() *log.XLog {
	zl, err := xlog_config.NewZapLoggerFromConfig(xlog_config.NewDevConfig())
	if err != nil {
		xerror.Exit(err)
	}

	zl = zl.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	return log.NewXLog().SetZapLogger(zl)
}()

func GetLog() XLog {
	return defaultLog
}

func SetLog(zl XLog) error {
	if zl == nil {
		return xerror.New("the params should not be nil")
	}
	defaultLog = zl.(*log.XLog)
	return nil
}

var (
	Debug   = defaultLog.Debug
	DebugF  = defaultLog.DebugF
	Info    = defaultLog.Info
	InfoF   = defaultLog.InfoF
	Warn    = defaultLog.Warn
	WarnF   = defaultLog.WarnF
	Error   = defaultLog.Error
	ErrorF  = defaultLog.ErrorF
	DPanic  = defaultLog.DPanic
	DPanicF = defaultLog.DPanicF
	Panic   = defaultLog.Panic
	PanicF  = defaultLog.PanicF
	Named   = defaultLog.Named
	With    = defaultLog.With
	FatalF  = defaultLog.FatalF
	Fatal   = defaultLog.Fatal
)
