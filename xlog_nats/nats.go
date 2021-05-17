package xlog_nats

import (
	"log"

	"go.uber.org/zap"
)

type logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	fatal *log.Logger
}

func (l *logger) Noticef(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.fatal.Printf(format, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *logger) Tracef(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

// NewLog will create a new logger which implement nats-server Logger.
func NewLog(zl *zap.Logger) *logger {
	l := &logger{}
	l.debug, _ = zap.NewStdLogAt(zl, zap.DebugLevel)
	l.info, _ = zap.NewStdLogAt(zl, zap.InfoLevel)
	l.warn, _ = zap.NewStdLogAt(zl, zap.WarnLevel)
	l.error, _ = zap.NewStdLogAt(zl, zap.ErrorLevel)
	l.fatal, _ = zap.NewStdLogAt(zl, zap.FatalLevel)
	return l
}
