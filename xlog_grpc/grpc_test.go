package xlog_grpc

import (
	"google.golang.org/grpc/grpclog"

	"testing"
)

func TestInfoFn(t *testing.T) {
	grpclog.Info("sss")
	grpclog.Infof("ok%d", 1)

	grpclog.Component("test").InfoDepth(0, "hello")
	grpclog.Component("test").Info("hello")
}
