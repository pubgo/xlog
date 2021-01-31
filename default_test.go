package xlog

import (
	"os"
	"testing"
)

var logs Xlog

func TestMain(m *testing.M) {
	Watch(func(log1 Xlog) {
		logs = log1.Named("test")
	})

	os.Exit(m.Run())
}

func TestInfo(t *testing.T) {
	logs.Info("test")
	Info("test")
}
