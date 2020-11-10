package log

import (
	"fmt"
	"strings"

	"github.com/pubgo/xerror"
	"go.uber.org/zap"

	"github.com/pubgo/xlog/internal"
)

var _ internal.XLog = (*xlog)(nil)

type xlog struct {
	zl *zap.Logger
}

func (log *xlog) Debugf(format string, a ...interface{}) {
	log.zl.Debug(fmt.Sprintf(format, a...))
}

func (log *xlog) Infof(format string, a ...interface{}) {
	log.zl.Info(fmt.Sprintf(format, a...))
}

func (log *xlog) Warnf(format string, a ...interface{}) {
	log.zl.Warn(fmt.Sprintf(format, a...))
}

func (log *xlog) Errorf(format string, a ...interface{}) {
	log.zl.Error(fmt.Sprintf(format, a...))
}

func (log *xlog) DPanicf(format string, a ...interface{}) {
	log.zl.DPanic(fmt.Sprintf(format, a...))
}

func (log *xlog) Panicf(format string, a ...interface{}) {
	log.zl.Panic(fmt.Sprintf(format, a...))
}

func (log *xlog) Fatalf(format string, a ...interface{}) {
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

func (log *xlog) SetZapLogger(zl *zap.Logger) *xlog {
	if zl == nil {
		return log
	}

	log.zl = zl
	return log
}

func (log *xlog) Named(s string, opts ...zap.Option) internal.XLog {
	if strings.TrimSpace(s) == "" {
		return log
	}

	return &xlog{log.zl.Named(s).WithOptions(opts...)}
}

func (log *xlog) Sync() error {
	return xerror.Wrap(log.zl.Sync())
}

type XLog = xlog

func NewXLog() *xlog {
	return &xlog{}
}
