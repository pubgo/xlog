package xlog_grpc

import (
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_abc"
	"github.com/pubgo/xlog/xlog_opts"
	"google.golang.org/grpc/grpclog"
)

func Init(log xlog.Xlog) {
	grpclog.SetLoggerV2(&loggerWrapper{log: log.Named("grpc", xlog_opts.AddCallerSkip(3))})
}

type loggerWrapper struct{ log xlog.Xlog }

func (l *loggerWrapper) Info(args ...interface{}) {
	l.log.InfoW(func(log xlog_abc.Logger) { log.Print(args...) })
}

func (l *loggerWrapper) Infoln(args ...interface{}) {
	l.log.InfoW(func(log xlog_abc.Logger) { log.Println(args...) })
}

func (l *loggerWrapper) Infof(format string, args ...interface{}) {
	l.log.InfoW(func(log xlog_abc.Logger) { log.Printf(format, args...) })
}

func (l *loggerWrapper) Warning(args ...interface{}) {
	l.log.WarnW(func(log xlog_abc.Logger) { log.Print(args...) })
}

func (l *loggerWrapper) Warningln(args ...interface{}) {
	l.log.WarnW(func(log xlog_abc.Logger) { log.Println(args...) })
}

func (l *loggerWrapper) Warningf(format string, args ...interface{}) {
	l.log.WarnW(func(log xlog_abc.Logger) { log.Printf(format, args...) })
}

func (l *loggerWrapper) Error(args ...interface{}) {
	l.log.ErrorW(func(log xlog_abc.Logger) { log.Print(args...) })
}

func (l *loggerWrapper) Errorln(args ...interface{}) {
	l.log.ErrorW(func(log xlog_abc.Logger) { log.Println(args...) })
}

func (l *loggerWrapper) Errorf(format string, args ...interface{}) {
	l.log.ErrorW(func(log xlog_abc.Logger) { log.Printf(format, args...) })
}

func (l *loggerWrapper) Fatal(args ...interface{}) {
	l.log.FatalW(func(log xlog_abc.Logger) { log.Print(args...) })
}

func (l *loggerWrapper) Fatalln(args ...interface{}) {
	l.log.FatalW(func(log xlog_abc.Logger) { log.Println(args...) })
}

func (l *loggerWrapper) Fatalf(format string, args ...interface{}) {
	l.log.FatalW(func(log xlog_abc.Logger) { log.Printf(format, args...) })
}

func (l *loggerWrapper) V(_ int) bool { return true }
