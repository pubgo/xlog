package example

import (
	"encoding/json"
	"github.com/pubgo/xlog/xlog_fields"
	"os"
	"testing"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_abc"
	"github.com/pubgo/xlog/xlog_config"
)

var log xlog_abc.Xlog

func TestMain(m *testing.M) {
	xlog.With()
	xlog.Watch(func(logs xlog_abc.Xlog) {
		log = logs.Named("test")
	})

	os.Exit(m.Run())
}

func initCfgFromJson() {
	cfg := `{
        "level": "info",
        "development": false,
        "disableCaller": false,
        "disableStacktrace": false,
        "sampling": {
                "initial": 100,
                "thereafter": 100
        },
        "encoding": "json",
        "encoderConfig": {
                "messageKey": "msg",
                "levelKey": "level",
                "timeKey": "ts",
                "nameKey": "logger",
                "callerKey": "caller",
                "stacktraceKey": "stacktrace",
                "lineEnding": "\n",
                "levelEncoder": "default",
                "timeEncoder": "default",
                "durationEncoder": "default",
                "callerEncoder": "full",
                "nameEncoder": "default"
        },
        "outputPaths": ["stderr"],
        "errorOutputPaths": ["stderr"],
        "initialFields": null
}`

	var cfg1 xlog_config.Config
	xerror.Exit(json.Unmarshal([]byte(cfg), &cfg1))
	zl, err := xlog_config.NewZapLogger(cfg1, func(opts *xlog_config.Config) {
		opts.Encoding = "console"
	})

	xerror.Exit(err)
	xerror.Exit(xlog.SetDefault(xlog.New(zl)))
}

func TestLog(t *testing.T) {
	xlog.Info("hello", xlog_fields.String("ss", "hello1"))
	xlog.Infof("hello %s", "1234")
	xlog.InfoM("hello %s", xlog.M{
		"test": "ok",
	})

	initCfgFromJson()
	
	xlog.Info("hello", xlog_fields.String("ss", "hello1"))
	xlog.Infof("hello %s", "1234")
	xlog.InfoM("hello %s", xlog.M{
		"test": "ok",
	})
}
