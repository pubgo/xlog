package xlog_std

import (
	"log"
	"os"
	"testing"
)

var ll *log.Logger

func TestMain(m *testing.M) {
	ll = New("hello")
	os.Exit(m.Run())
}

func TestInfoFn(t *testing.T) {
	ll.Println("hello")
}
