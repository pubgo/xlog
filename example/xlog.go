package main

import (
	"encoding/json"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

var log = xlog.GetLogger("test", zap.Fields(zap.String("name", "test_hello")))

func main() {
	log.Info("hello", zap.String("ss", "hello1"))
	log.Info("hello", zap.String("ss", "hello1"))
	log.Infof("hello %s", "1234")
	log.Info("hello %s", xlog.M{
		"test": "ok",
	})
	log.Warn("test")

	initCfgFromJson()

	log.Info("hello", zap.String("ss", "hello1"))
	log.Infof("hello %s", "1234")
	log.Info("hello %s", xlog.M{
		"test": "ok",
	})
	log.Warn("test")
	log.Info("hello", zap.String("ss", "hello1"))
	log.WarnW(func(logs xlog.Logger) {
		logs.Println("hello w")
	})
	log.WarnW(func(logs xlog.Logger) {
		logs.Print("hhhh jnjnjnj")
	})
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
	cfg1.Encoding = "console"
	zl, err := cfg1.Build()

	xerror.Exit(err)
	xerror.Exit(xlog.SetDefault(zl))
	log.Warn("test")
}
