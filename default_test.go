package xlog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var logs XLog

func TestMain(m *testing.M) {
	Watch(func(log1 XLog) {
		logs = log1.Named("test")
	})

	os.Exit(m.Run())
}

func TestInfo(t *testing.T) {
	logs.Info("test")

	assert.Nil(t, SetDefault(GetDevLog().Named("hello")))
	logs.Info("test")
}
