package main

import (
	"encoding/json"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_config"
)

var log xlog.Xlog

func init() {
	xlog.Watch(func(logs xlog.Xlog) {
		log = logs.Named("test")
	})
}

func main() {
	log.Info("hello", xlog.String("ss", "hello1"))
	xlog.Info("hello", xlog.String("ss", "hello1"))
	xlog.Infof("hello %s", "1234")
	xlog.InfoM("hello %s", xlog.M{
		"test": "ok",
	})
	xlog.Warn("test")

	initCfgFromJson()

	xlog.Info("hello", xlog.String("ss", "hello1"))
	xlog.Infof("hello %s", "1234")
	xlog.InfoM("hello %s", xlog.M{
		"test": "ok",
	})
	xlog.Warn("test")
	log.Info("hello", xlog.String("ss", "hello1"))
	xlog.WarnW(func(log xlog.Logger) {
		log.Println("hello w")
	})
	log.WarnW(func(log xlog.Logger) {
		log.Print("hhhh jnjnjnj")
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
	xerror.Exit(xlog.SetDefault(xlog.New(zl)))
	xlog.Warn("test")
}
