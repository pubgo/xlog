package log

import (
	"github.com/pubgo/xlog/internal"
	"go.uber.org/zap"
)

var _ internal.ILog = (*xlog)(nil)

type xlog struct {
	zl *zap.Logger
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

func (log *xlog) With(fields ...zap.Field) internal.ILog {
	return &xlog{log.zl.With(fields...)}
}

func (log *xlog) Named(s string) internal.ILog {
	return &xlog{log.zl.Named(s).With(zap.Namespace(s))}
}

func (log *xlog) GetZap() *zap.Logger {
	return log.zl
}

type XLog = xlog

func NewXLog(zl *zap.Logger) *XLog {
	xl := &XLog{}
	xl.zl = zl
	return xl
}
