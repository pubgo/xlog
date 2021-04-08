package log

import (
	"fmt"
	"strings"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/xlog_abc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ xlog_abc.Xlog = (*xlog)(nil)
var _ xlog_abc.Logger = (*xlog)(nil)

func New() *xlog { return &xlog{} }

type xlog struct {
	zl  *zap.Logger
	lvl zapcore.Level
}

func (t *xlog) PrintM(m xlog_abc.M) {
	var fields = make([]zap.Field, 0, len(m))
	for k, v := range m {
		fields = append(fields, zap.Any(k, v))
	}
	t.zl.Check(t.lvl, "").Write(fields...)
}

func (t *xlog) Printf(format string, v ...interface{}) {
	t.zl.Check(t.lvl, fmt.Sprintf(format, v...)).Write()
}

func (t *xlog) Print(args ...interface{}) {
	msg, fields := fields(args...)
	t.zl.Check(t.lvl, msg).Write(fields...)
}

func (t *xlog) Println(args ...interface{}) {
	msg, fields := fields(append(args, "\n")...)
	t.zl.Check(t.lvl, msg).Write(fields...)
}

func (t *xlog) DebugW(f func(log xlog_abc.Logger)) {
	if !t.zl.Core().Enabled(zap.DebugLevel) {
		return
	}

	f(&xlog{zl: t.zl, lvl: zap.DebugLevel})
}

func (t *xlog) InfoW(f func(log xlog_abc.Logger)) {
	if !t.zl.Core().Enabled(zap.InfoLevel) {
		return
	}

	f(&xlog{zl: t.zl, lvl: zap.InfoLevel})
}

func (t *xlog) WarnW(f func(log xlog_abc.Logger)) {
	if !t.zl.Core().Enabled(zap.WarnLevel) {
		return
	}

	f(&xlog{zl: t.zl, lvl: zap.WarnLevel})
}

func (t *xlog) ErrorW(f func(log xlog_abc.Logger)) {
	if !t.zl.Core().Enabled(zap.ErrorLevel) {
		return
	}

	f(&xlog{zl: t.zl, lvl: zap.ErrorLevel})
}

func (t *xlog) DPanicW(f func(log xlog_abc.Logger)) {
	if !t.zl.Core().Enabled(zap.DPanicLevel) {
		return
	}

	f(&xlog{zl: t.zl, lvl: zap.DPanicLevel})
}

func (t *xlog) PanicW(f func(log xlog_abc.Logger)) {
	if !t.zl.Core().Enabled(zap.PanicLevel) {
		return
	}

	f(&xlog{zl: t.zl, lvl: zap.PanicLevel})
}

func (t *xlog) FatalW(f func(log xlog_abc.Logger)) {
	if !t.zl.Core().Enabled(zap.FatalLevel) {
		return
	}

	f(&xlog{zl: t.zl, lvl: zap.FatalLevel})
}

func fields(args ...interface{}) (string, []zap.Field) {
	var msg = ""

	if len(args) == 0 {
		return msg, nil
	}

	var fields = make([]zap.Field, 0, len(args))
	var fields1 = make([]interface{}, 0, len(args))
	for i := range args {
		field := args[i]
		if field == nil {
			continue
		}

		if f, ok := field.(zap.Field); ok {
			fields = append(fields, f)
			continue
		}

		if f, ok := field.(xlog_abc.M); ok {
			for k, v := range f {
				fields = append(fields, zap.Any(k, v))
			}
			continue
		}

		fields1 = append(fields1, field)
	}

	if len(fields1) > 0 {
		msg = fmt.Sprint(fields1...)
	}

	return msg, fields
}

func (t *xlog) Debug(args ...interface{}) {
	if !t.zl.Core().Enabled(zap.DebugLevel) {
		return
	}

	msg, fields := fields(args...)
	t.zl.Debug(msg, fields...)
}

func (t *xlog) Info(args ...interface{}) {
	if !t.zl.Core().Enabled(zap.InfoLevel) {
		return
	}

	msg, fields := fields(args...)
	t.zl.Info(msg, fields...)
}

func (t *xlog) Warn(args ...interface{}) {
	if !t.zl.Core().Enabled(zap.WarnLevel) {
		return
	}

	msg, fields := fields(args...)
	t.zl.Warn(msg, fields...)
}

func (t *xlog) Error(args ...interface{}) {
	if !t.zl.Core().Enabled(zap.ErrorLevel) {
		return
	}

	msg, fields := fields(args...)
	t.zl.Error(msg, fields...)
}

func (t *xlog) DPanic(args ...interface{}) {
	if !t.zl.Core().Enabled(zap.DPanicLevel) {
		return
	}

	msg, fields := fields(args...)
	t.zl.DPanic(msg, fields...)
}

func (t *xlog) Panic(args ...interface{}) {
	if !t.zl.Core().Enabled(zap.PanicLevel) {
		return
	}

	msg, fields := fields(args...)
	t.zl.Panic(msg, fields...)
}

func (t *xlog) Fatal(args ...interface{}) {
	if !t.zl.Core().Enabled(zap.FatalLevel) {
		return
	}

	msg, fields := fields(args...)
	t.zl.Fatal(msg, fields...)
}

func (t *xlog) Debugf(format string, a ...interface{}) {
	if !t.zl.Core().Enabled(zap.DebugLevel) {
		return
	}

	t.zl.Debug(fmt.Sprintf(format, a...))
}

func (t *xlog) Infof(format string, a ...interface{}) {
	if !t.zl.Core().Enabled(zap.InfoLevel) {
		return
	}

	t.zl.Info(fmt.Sprintf(format, a...))
}

func (t *xlog) Warnf(format string, a ...interface{}) {
	if !t.zl.Core().Enabled(zap.WarnLevel) {
		return
	}

	t.zl.Warn(fmt.Sprintf(format, a...))
}

func (t *xlog) Errorf(format string, a ...interface{}) {
	if !t.zl.Core().Enabled(zap.ErrorLevel) {
		return
	}

	t.zl.Error(fmt.Sprintf(format, a...))
}

func (t *xlog) DPanicf(format string, a ...interface{}) {
	if !t.zl.Core().Enabled(zap.DPanicLevel) {
		return
	}

	t.zl.DPanic(fmt.Sprintf(format, a...))
}

func (t *xlog) Panicf(format string, a ...interface{}) {
	if !t.zl.Core().Enabled(zap.PanicLevel) {
		return
	}

	t.zl.Panic(fmt.Sprintf(format, a...))
}

func (t *xlog) Fatalf(format string, a ...interface{}) {
	if !t.zl.Core().Enabled(zap.FatalLevel) {
		return
	}

	t.zl.Fatal(fmt.Sprintf(format, a...))
}

func (t *xlog) With(fields ...zap.Field) xlog_abc.Xlog { return &xlog{zl: t.zl.With(fields...)} }
func (t *xlog) Sync() error                            { return xerror.Wrap(t.zl.Sync()) }

func (t *xlog) SetZapLogger(zl *zap.Logger) *xlog {
	if zl == nil {
		t.zl.Warn("[zl] is nil")
		return t
	}

	t.zl = zl
	return t
}

func (t *xlog) Named(s string, opts ...zap.Option) xlog_abc.Xlog {
	zl := t.zl
	if len(opts) > 0 {
		zl = zl.WithOptions(opts...)
	}

	if strings.TrimSpace(s) != "" {
		zl = zl.Named(s)
	}

	return &xlog{zl: zl}
}
