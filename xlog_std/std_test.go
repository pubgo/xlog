package xlog_std

import (
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_abc"

	"log"
	"os"
	"testing"
)

var ll *log.Logger

func TestMain(m *testing.M) {
	xlog.Watch(func(log1 xlog_abc.Xlog) {
		ll = New("hello")
	})

	os.Exit(m.Run())
}

func TestInfoFn(t *testing.T) {
	ll.Println("hello")
}
