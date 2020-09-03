# xlog
xlog is a simple and easy-to-use logger with zap logger as the bottom abstraction

## 介绍
xlog是一个对zap logger的简单的封装, 意在简化配置, 增强可控, 简洁易用。


## example
```go
package example_test

import (
	"testing"
	"time"

	"github.com/pubgo/dix"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/xlog_config"
)

var log = xlog.GetLog()

func init() {
	dix.Go(func(log1 xlog.XLog) {
		log = log1.
			Named("service").With(xlog.String("key", "service")).
			Named("hello").With(xlog.String("key", "hello")).
			Named("world").With(xlog.String("key", "world"))
	})
}

func TestExample(t *testing.T) {
	log.Debug("hello",
		xlog.Any("hss", "ss"),
	)

	dix.Go(initCfgFromJsonDebug(time.Now().Format("2006-01-02 15:04:05")))

	log.Info("hello",
		xlog.Any("hss", "ss"),
	)
	//fmt.Println(dix.Graph())
}

func initCfgFromJsonDebug(name string) internal.XLog {
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

	zl, err := xlog_config.NewZapLoggerFromJson([]byte(cfg), xlog_config.WithEncoding("console"))
	xerror.Exit(err)
	return xlog.New(zl.WithOptions(xlog.AddCallerSkip(1)))
}
```
