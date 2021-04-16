package xlog_grpc

import (
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_abc"
	"github.com/pubgo/xlog/xlog_opts"
	"google.golang.org/grpc/grpclog"

	"os"
	"testing"
)

func TestMain(m *testing.M) {
	xlog.Watch(func(log1 xlog_abc.Xlog) {
		Init(log1.Named("test", xlog_opts.Fields(xlog.String("name", "hello"))))
	})

	os.Exit(m.Run())
}

func TestInfoFn(t *testing.T) {
	grpclog.Info("sss")
	grpclog.Infof("ok%d", 1)
}
