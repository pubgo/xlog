package xlog_grpc

import (
	"github.com/pubgo/xlog"
	"google.golang.org/grpc/grpclog"
)

func Init(log xlog.Xlog) {
	grpclog.SetLoggerV2(&loggerWrapper{log: log})
}

type loggerWrapper struct{ log xlog.Xlog }

func (l *loggerWrapper) Info(args ...interface{})                    { l.log.Info(args...) }
func (l *loggerWrapper) Infoln(args ...interface{})                  { l.log.Info(args...) }
func (l *loggerWrapper) Infof(format string, args ...interface{})    { l.log.Infof(format, args...) }
func (l *loggerWrapper) Warning(args ...interface{})                 { l.log.Warn(args...) }
func (l *loggerWrapper) Warningln(args ...interface{})               { l.log.Warn(args...) }
func (l *loggerWrapper) Warningf(format string, args ...interface{}) { l.log.Warnf(format, args...) }
func (l *loggerWrapper) Error(args ...interface{})                   { l.log.Error(args...) }
func (l *loggerWrapper) Errorln(args ...interface{})                 { l.log.Error(args...) }
func (l *loggerWrapper) Errorf(format string, args ...interface{})   { l.log.Errorf(format, args...) }
func (l *loggerWrapper) Fatal(args ...interface{})                   { l.log.Fatal(args...) }
func (l *loggerWrapper) Fatalln(args ...interface{})                 { l.log.Fatal(args...) }
func (l *loggerWrapper) Fatalf(format string, args ...interface{})   { l.log.Fatalf(format, args...) }
func (l *loggerWrapper) V(v int) bool                                { return true }
