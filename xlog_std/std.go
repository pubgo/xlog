package xlog_std

import (
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_opts"

	"io"
	"log"
)

func New(name string, opts ...xlog_opts.Option) *log.Logger {
	return log.New(&w{l: xlog.Named(name, append(opts, xlog_opts.AddCallerSkip(3))...)}, "", log.LstdFlags|log.Llongfile)
}

var _ io.Writer = (*w)(nil)

type w struct {
	l xlog.Xlog
}

func (t *w) Write(p []byte) (n int, err error) {
	t.l.Info(string(p))
	return 0, err
}
