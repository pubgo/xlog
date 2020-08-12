# xlog
xlog is a simple and easy-to-use logger with zap logger as the bottom abstraction

## 介绍
xlog是一个对zap logger的简单的封装, 意在简化配置, 增加可控, 简单易用。


## example
```go
package example_test

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_config"
	"testing"
)

var fields = xlog.FieldOf(
	xlog.String("key", "value"),
)
var log = xlog.GetDevLog().With(fields...)

func init() {
	//initCfgFromJson()
	initCfgFromJsonDebug()
	log = xlog.GetLog().
		Named("service").With(fields...).
		Named("hello").With(fields...).
		Named("world").With(fields...)
}

func TestExample(t *testing.T) {
	log.Debug("hello",
		xlog.Any("hss", "ss"),
	)

	log.Info("hello",
		xlog.Any("hss", "ss"),
	)

	log.Error("hello",
		xlog.Any("hss", "ss"),
	)

	log.Info("hello",
		xlog.Any("hss", "ss"),
	)
}

func initCfgFromJsonDebug() {
	cfg := `{
        "level": "debug",
        "development": true,
        "disableCaller": false,
        "disableStacktrace": false,
        "sampling": null,
        "encoding": "console",
        "encoderConfig": {
                "messageKey": "M",
                "levelKey": "L",
                "timeKey": "T",
                "nameKey": "N",
                "callerKey": "C",
                "stacktraceKey": "S",
                "lineEnding": "\n",
                "levelEncoder": "capitalColor",
                "timeEncoder": "iso8601",
                "durationEncoder": "string",
                "callerEncoder": "default",
                "nameEncoder": ""
        },
        "outputPaths": [
                "stderr"
        ],
        "errorOutputPaths": [
                "stderr"
        ],
        "initialFields": null
}`
	xerror.Exit(xlog_config.InitFromJson(
		[]byte(cfg),
	))
}
```
