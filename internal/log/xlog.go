package log

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"go.uber.org/zap"
)

var fieldsPool = sync.Pool{
	New: func() interface{} {
		return make([]zap.Field, 1)
	},
}

func getFields() []zap.Field { return fieldsPool.Get().([]zap.Field) }
func put(fields []zap.Field) { fieldsPool.Put(fields[:0]) }

var _ internal.XLog = (*xlog)(nil)

type xlog struct{ *zap.Logger }
type log func(msg string, fields ...zap.Field)

func logHandle(log log, msg string, fn func(fields internal.Fields)) {
	var fields = getFields()
	defer put(fields)

	fn(fields)
	log(msg, fields...)
}

func (log *xlog) DebugFn(msg string, fn func(fields internal.Fields)) {
	logHandle(log.Logger.Debug, msg, fn)
}

func (log *xlog) InfoFn(msg string, fn func(fields internal.Fields)) {
	logHandle(log.Logger.Info, msg, fn)
}

func (log *xlog) WarnFn(msg string, fn func(fields internal.Fields)) {
	logHandle(log.Logger.Warn, msg, fn)
}

func (log *xlog) ErrorFn(msg string, fn func(fields internal.Fields)) {
	logHandle(log.Logger.Error, msg, fn)
}

func (log *xlog) DPanicFn(msg string, fn func(fields internal.Fields)) {
	logHandle(log.Logger.DPanic, msg, fn)
}

func (log *xlog) PanicFn(msg string, fn func(fields internal.Fields)) {
	logHandle(log.Logger.Panic, msg, fn)
}

func (log *xlog) FatalFn(msg string, fn func(fields internal.Fields)) {
	logHandle(log.Logger.Fatal, msg, fn)
}

type XLog = xlog

func New() *xlog                                           { return &xlog{} }
func (log *xlog) Warningf(format string, a ...interface{}) { log.Warningf(format, a...) }
func (log *xlog) Debugf(format string, a ...interface{})   { log.Logger.Debug(fmt.Sprintf(format, a...)) }
func (log *xlog) Infof(format string, a ...interface{})    { log.Logger.Info(fmt.Sprintf(format, a...)) }
func (log *xlog) Warnf(format string, a ...interface{})    { log.Logger.Warn(fmt.Sprintf(format, a...)) }
func (log *xlog) Errorf(format string, a ...interface{})   { log.Logger.Error(fmt.Sprintf(format, a...)) }
func (log *xlog) DPanicf(format string, a ...interface{}) {
	log.Logger.DPanic(fmt.Sprintf(format, a...))
}
func (log *xlog) Panicf(format string, a ...interface{})       { log.Logger.Panic(fmt.Sprintf(format, a...)) }
func (log *xlog) Fatalf(format string, a ...interface{})       { log.Logger.Fatal(fmt.Sprintf(format, a...)) }
func (log *xlog) Warning(msg string, fields ...internal.Field) { log.Logger.Warn(msg, fields...) }
func (log *xlog) With(fields ...zap.Field) internal.XLog       { return &xlog{log.Logger.With(fields...)} }
func (log *xlog) Sync() error                                  { return xerror.Wrap(log.Logger.Sync()) }

func (log *xlog) SetZapLogger(zl *zap.Logger) *xlog {
	if zl == nil {
		log.Warn("[zl] is nil")
		return log
	}

	log.Logger = zl
	return log
}

func (log *xlog) Named(s string, opts ...zap.Option) internal.XLog {
	if strings.TrimSpace(s) == "" {
		return log
	}

	return &xlog{log.Logger.Named(s).WithOptions(opts...)}
}
