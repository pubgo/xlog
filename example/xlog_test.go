package example

import (
	"encoding/json"
	"testing"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_config"
)

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
                "callerEncoder": "default",
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

func TestXLog(t *testing.T) {
	initCfgFromJson()
	xlog.Infof("hello %s", "1234")
}

func TestDevLog(t *testing.T) {
	xlog.Infof("hello %s", "1234")
}
