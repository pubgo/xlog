package example

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_abc"
	"github.com/pubgo/xlog/xlog_config"
)

var log xlog_abc.Xlog

func TestMain(m *testing.M) {
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
	xerror.Exit(xlog.Init(zl))
}

func TestLog(t *testing.T) {
	xlog.Infof("hello %s", "1234")
	log.Info("hello")
	log.InfoFn("hello", func(fields *xlog_abc.Fields) {
		fields.String("ss", "hello1")
	})

	xlog.InfoFn("", func(fields *xlog_abc.Fields) {
		fields.String("ss", "hello1")
	})

	initCfgFromJson()
	xlog.Infof("hello %s", "1234")
	log.Info("hello")
}
