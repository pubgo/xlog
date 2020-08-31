package log

import (
	"fmt"
	"github.com/pubgo/xerror"
	"go.uber.org/zap"

	"github.com/pubgo/xlog/internal"
)

var _ internal.XLog = (*xlog)(nil)

type xlog struct {
	zl *zap.Logger
}

func (log *xlog) DebugF(format string, a ...interface{}) {
	log.zl.Debug(fmt.Sprintf(format, a...))
}

func (log *xlog) InfoF(format string, a ...interface{}) {
	log.zl.Info(fmt.Sprintf(format, a...))
}

func (log *xlog) WarnF(format string, a ...interface{}) {
	log.zl.Warn(fmt.Sprintf(format, a...))
}

func (log *xlog) ErrorF(format string, a ...interface{}) {
	log.zl.Error(fmt.Sprintf(format, a...))
}

func (log *xlog) DPanicF(format string, a ...interface{}) {
	log.zl.DPanic(fmt.Sprintf(format, a...))
}

func (log *xlog) PanicF(format string, a ...interface{}) {
	log.zl.Panic(fmt.Sprintf(format, a...))
}

func (log *xlog) FatalF(format string, a ...interface{}) {
	log.zl.Fatal(fmt.Sprintf(format, a...))
}

func (log *xlog) Debug(msg string, fields ...internal.Field) {
	log.zl.Debug(msg, fields...)
}

func (log *xlog) Info(msg string, fields ...internal.Field) {
	log.zl.Info(msg, fields...)
}

func (log *xlog) Warn(msg string, fields ...internal.Field) {
	log.zl.Warn(msg, fields...)
}

func (log *xlog) Error(msg string, fields ...internal.Field) {
	log.zl.Error(msg, fields...)
}

func (log *xlog) DPanic(msg string, fields ...internal.Field) {
	log.zl.DPanic(msg, fields...)
}

func (log *xlog) Panic(msg string, fields ...internal.Field) {
	log.zl.Panic(msg, fields...)
}

func (log *xlog) Fatal(msg string, fields ...internal.Field) {
	log.zl.Fatal(msg, fields...)
}

func (log *xlog) With(fields ...zap.Field) internal.XLog {
	return &xlog{log.zl.With(fields...)}
}

func (log *xlog) Named(s string) internal.XLog {
	return &xlog{log.zl.Named(s).With(zap.Namespace(s))}
}

func (log *xlog) Sync() error {
	return xerror.Wrap(log.zl.Sync())
}

type XLog = xlog

func NewXLog(zl *zap.Logger) *xlog {
	xl := &xlog{}
	xl.zl = zl
	return xl
}
