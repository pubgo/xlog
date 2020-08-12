package example_test

import (
	"github.com/pubgo/dix"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/internal"
	"github.com/pubgo/xlog/xlog_config"
	"testing"
	"time"
)

var log = xlog.GetDevLog()

func init() {
	xerror.Exit(dix.Dix(func(log1 *xlog.XLog) {
		log = log1.
			Named("service").With(xlog.String("key", "value1")).
			Named("hello").With(xlog.String("key", "value2")).
			Named("world").With(xlog.String("key", "value3"))
	}))
}

func TestExample(t *testing.T) {
	for {
		//fmt.Println(dix.Graph())

		log.Debug("hello",
			xlog.Any("hss", "ss"),
		)

		log.Info("hello",
			xlog.Any("hss", "ss"),
		)
		time.Sleep(time.Second)
		xerror.Exit(dix.Dix(initCfgFromJsonDebug(time.Now().Format("2006-01-02 15:04:05"))))
	}
}

func initCfgFromJsonDebug(name string) internal.ILog {
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

	xx, err := xlog_config.NewFromJson(
		[]byte(cfg),
	)
	xerror.Exit(err)
	return xx.Named(name)
}
