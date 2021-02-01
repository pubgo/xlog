package log

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/xlog_abc"
	"go.uber.org/zap"
)

var fieldsPool = sync.Pool{New: func() interface{} { return make([]zap.Field, 0, 1) }}

func getFields() xlog_abc.Fields { return fieldsPool.Get().([]zap.Field) }
func put(fields []zap.Field)     { fieldsPool.Put(fields[:0]) }

var _ xlog_abc.Xlog = (*xlog)(nil)

type xlog struct{ zl *zap.Logger }

func (log *xlog) DebugFn(msg string, fn func(fields *xlog_abc.Fields)) {
	var fields = getFields()
	defer put(fields)

	fn(&fields)
	log.zl.Debug(msg, fields...)
}

func (log *xlog) InfoFn(msg string, fn func(fields *xlog_abc.Fields)) {
	var fields = getFields()
	defer put(fields)

	fn(&fields)
	log.zl.Info(msg, fields...)
}

func (log *xlog) WarnFn(msg string, fn func(fields *xlog_abc.Fields)) {
	var fields = getFields()
	defer put(fields)

	fn(&fields)
	log.zl.Warn(msg, fields...)
}

func (log *xlog) ErrorFn(msg string, fn func(fields *xlog_abc.Fields)) {
	var fields = getFields()
	defer put(fields)

	fn(&fields)
	log.zl.Error(msg, fields...)
}

func (log *xlog) DPanicFn(msg string, fn func(fields *xlog_abc.Fields)) {
	var fields = getFields()
	defer put(fields)

	fn(&fields)
	log.zl.DPanic(msg, fields...)
}

func (log *xlog) PanicFn(msg string, fn func(fields *xlog_abc.Fields)) {
	var fields = getFields()
	defer put(fields)

	fn(&fields)
	log.zl.Panic(msg, fields...)
}

func (log *xlog) FatalFn(msg string, fn func(fields *xlog_abc.Fields)) {
	var fields = getFields()
	defer put(fields)

	fn(&fields)
	log.zl.Fatal(msg, fields...)
}

func New() *xlog { return &xlog{} }

func (log *xlog) Debug(msg string, fields ...xlog_abc.Field)   { log.zl.Debug(msg, fields...) }
func (log *xlog) Info(msg string, fields ...xlog_abc.Field)    { log.zl.Info(msg, fields...) }
func (log *xlog) Warn(msg string, fields ...xlog_abc.Field)    { log.zl.Warn(msg, fields...) }
func (log *xlog) Warning(msg string, fields ...xlog_abc.Field) { log.zl.Warn(msg, fields...) }
func (log *xlog) Error(msg string, fields ...xlog_abc.Field)   { log.zl.Error(msg, fields...) }
func (log *xlog) DPanic(msg string, fields ...xlog_abc.Field)  { log.zl.DPanic(msg, fields...) }
func (log *xlog) Panic(msg string, fields ...xlog_abc.Field)   { log.zl.Panic(msg, fields...) }
func (log *xlog) Fatal(msg string, fields ...xlog_abc.Field)   { log.zl.Fatal(msg, fields...) }

func (log *xlog) Warningf(format string, a ...interface{}) { log.zl.Warn(fmt.Sprintf(format, a...)) }
func (log *xlog) Debugf(format string, a ...interface{})   { log.zl.Debug(fmt.Sprintf(format, a...)) }
func (log *xlog) Infof(format string, a ...interface{})    { log.zl.Info(fmt.Sprintf(format, a...)) }
func (log *xlog) Warnf(format string, a ...interface{})    { log.zl.Warn(fmt.Sprintf(format, a...)) }
func (log *xlog) Errorf(format string, a ...interface{})   { log.zl.Error(fmt.Sprintf(format, a...)) }
func (log *xlog) DPanicf(format string, a ...interface{})  { log.zl.DPanic(fmt.Sprintf(format, a...)) }
func (log *xlog) Panicf(format string, a ...interface{})   { log.zl.Panic(fmt.Sprintf(format, a...)) }
func (log *xlog) Fatalf(format string, a ...interface{})   { log.zl.Fatal(fmt.Sprintf(format, a...)) }

func (log *xlog) With(fields ...zap.Field) xlog_abc.Xlog { return &xlog{log.zl.With(fields...)} }
func (log *xlog) Sync() error                            { return xerror.Wrap(log.zl.Sync()) }

func (log *xlog) SetZapLogger(zl *zap.Logger) *xlog {
	if zl == nil {
		log.Warn("[zl] is nil")
		return log
	}

	log.zl = zl
	return log
}

func (log *xlog) Named(s string, opts ...zap.Option) xlog_abc.Xlog {
	zl := log.zl
	if len(opts) > 0 {
		zl = zl.WithOptions(opts...)
	}

	if strings.TrimSpace(s) != "" {
		zl = zl.Named(s)
	}

	return &xlog{zl: zl}
}
