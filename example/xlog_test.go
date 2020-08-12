package example

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/xlog"
	"github.com/pubgo/xlog/xlog_config"
	"testing"
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
	xerror.Exit(xlog_config.InitFromJson(
		[]byte(cfg),
		xlog_config.WithDevelopment(),
		xlog_config.WithEncoding("console")),
	)
}

func TestXLog(t *testing.T) {
	initCfgFromJson()
	xlog.InfoF("hello %s", "1234")
}
