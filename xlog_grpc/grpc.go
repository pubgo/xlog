package xlog_grpc

import (
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_abc"
	"github.com/pubgo/xlog/xlog_opts"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

var _ grpclog.LoggerV2 = (*loggerWrapper)(nil)

func Init(log xlog.Xlog) {
	grpclog.SetLoggerV2(&loggerWrapper{log: log.Named("grpc", xlog_opts.AddCallerSkip(4))})
}

type loggerWrapper struct {
	log           xlog.Xlog
	printFilter   func(args ...interface{}) bool
	printfFilter  func(format string, args ...interface{}) bool
	printlnFilter func(args ...interface{}) bool
}

func (l *loggerWrapper) SetPrintFilter(filter func(args ...interface{}) bool) {
	l.printFilter = filter
}
func (l *loggerWrapper) SetPrintfFilter(filter func(format string, args ...interface{}) bool) {
	l.printfFilter = filter
}
func (l *loggerWrapper) SetPrintlnFilter(filter func(args ...interface{}) bool) {
	l.printlnFilter = filter
}

func (l *loggerWrapper) filter(args ...interface{}) bool {
	return l.printFilter != nil && l.printFilter(args...)
}

func (l *loggerWrapper) filterf(format string, args ...interface{}) bool {
	return l.printfFilter != nil && l.printfFilter(format, args...)
}

func (l *loggerWrapper) filterln(args ...interface{}) bool {
	return l.printlnFilter != nil && l.printlnFilter(args...)
}

func (l *loggerWrapper) Info(args ...interface{}) {
	if l.filter(args) {
		return
	}

	l.log.InfoW(func(log xlog_abc.Logger) { log.Print(args...) })
}

func (l *loggerWrapper) Infoln(args ...interface{}) {
	if l.filterln(args) {
		return
	}

	l.log.InfoW(func(log xlog_abc.Logger) { log.Println(args...) })
}

func (l *loggerWrapper) Infof(format string, args ...interface{}) {
	if l.filterf(format, args...) {
		return
	}

	l.log.InfoW(func(log xlog_abc.Logger) { log.Printf(format, args...) })
}

func (l *loggerWrapper) Warning(args ...interface{}) {
	if l.filter(args...) {
		return
	}

	l.log.WarnW(func(log xlog_abc.Logger) { log.Print(args...) })
}

func (l *loggerWrapper) Warningln(args ...interface{}) {
	if l.filterln(args) {
		return
	}

	l.log.WarnW(func(log xlog_abc.Logger) { log.Println(args...) })
}

func (l *loggerWrapper) Warningf(format string, args ...interface{}) {
	if l.filterf(format, args...) {
		return
	}

	l.log.WarnW(func(log xlog_abc.Logger) { log.Printf(format, args...) })
}

func (l *loggerWrapper) Error(args ...interface{}) {
	if l.filter(args...) {
		return
	}

	l.log.ErrorW(func(log xlog_abc.Logger) { log.Print(args...) })
}

func (l *loggerWrapper) Errorln(args ...interface{}) {
	if l.filterln(args) {
		return
	}

	l.log.ErrorW(func(log xlog_abc.Logger) { log.Println(args...) })
}

func (l *loggerWrapper) Errorf(format string, args ...interface{}) {
	if l.filterf(format, args...) {
		return
	}

	l.log.ErrorW(func(log xlog_abc.Logger) { log.Printf(format, args...) })
}

func (l *loggerWrapper) Fatal(args ...interface{}) {
	if l.filter(args...) {
		return
	}

	l.log.FatalW(func(log xlog_abc.Logger) { log.Print(args...) })
}

func (l *loggerWrapper) Fatalln(args ...interface{}) {
	if l.filterln(args) {
		return
	}

	l.log.FatalW(func(log xlog_abc.Logger) { log.Println(args...) })
}

func (l *loggerWrapper) Fatalf(format string, args ...interface{}) {
	if l.filterf(format, args...) {
		return
	}

	l.log.FatalW(func(log xlog_abc.Logger) { log.Printf(format, args...) })
}

func (l *loggerWrapper) V(_ int) bool { return true }
func (l *loggerWrapper) Lvl(lvl int) grpclog.LoggerV2 {
	return &loggerWrapper{log: l.log.Named("", xlog_opts.IncreaseLevel(zapcore.Level(lvl)))}
}
