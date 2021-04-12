package xlog

import (
	"os"
	"testing"

	"github.com/pubgo/xlog/xlog_opts"
)

var logs Xlog

func TestMain(m *testing.M) {
	Watch(func(log1 Xlog) {
		logs = log1.Named("test", xlog_opts.Fields(String("name", "hello")))
	})

	os.Exit(m.Run())
}

func TestInfo(t *testing.T) {
	logs.Info("test")
	Info("test")
}

func TestInfoFn(t *testing.T) {
	Debug("ok11")
	logs.Debug("dddd2")

	DebugW(func(log Logger) {
		log.Println("ok111")
		log.Println("ok111")
	})

	logs.DebugW(func(log Logger) {
		log.Println("sss")
		log.Print("ok")
		log.Printf("ok%d", 1)
		log.PrintM(M{
			"hello": 2,
		})
	})

	logs.InfoW(func(log Logger) {
		log.Println("sss")
		log.Println("sss")
		log.Print("ok")
		log.Printf("ok%d", 2)
	})
}
