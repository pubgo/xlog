package xlog_std

import (
	"github.com/pubgo/xlog"
	"go.uber.org/zap"

	"io"
	"log"
)

func New(name string, opts ...zap.Option) *log.Logger {
	return log.New(&w{l: xlog.GetLogger(name, append(opts, zap.AddCallerSkip(3))...)}, "", log.LstdFlags|log.Llongfile)
}

var _ io.Writer = (*w)(nil)

type w struct {
	l xlog.Xlog
}

func (t *w) Write(p []byte) (n int, err error) {
	t.l.Info(string(p))
	return 0, err
}
