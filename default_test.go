package xlog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

	ll := getDefault().Named("hello")
	assert.NotNil(t, ll)
	ll.Info("test")
}
