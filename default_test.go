package xlog

import (
	"os"
	"testing"

	"github.com/pubgo/xlog/xlog_abc"
)

var logs xlog_abc.Xlog

func TestMain(m *testing.M) {
	Watch(func(log1 xlog_abc.Xlog) {
		logs = log1.Named("test")
	})

	os.Exit(m.Run())
}

func TestInfo(t *testing.T) {
	logs.Info("test")
	Info("test")
}
