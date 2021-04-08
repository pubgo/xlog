package xlog_grpc

import (
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_abc"
	"github.com/pubgo/xlog/xlog_opts"

	"os"
	"testing"
)

var log *loggerWrapper

func TestMain(m *testing.M) {
	xlog.Watch(func(log1 xlog_abc.Xlog) {
		log = &loggerWrapper{log1.Named("test",
			xlog_opts.AddCallerSkip(3),
			xlog_opts.Fields(xlog.String("name", "hello")))}
	})

	os.Exit(m.Run())
}

func TestInfoFn(t *testing.T) {
	log.Info("sss")
	log.Infof("ok%d", 1)
}
