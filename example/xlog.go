package main

import (
	"encoding/json"
	"fmt"
	stdLog "log"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
)

var log = xlog.GetLogger("test.12345", zap.Fields(zap.String("name", "test_hello")))

func main() {
	stdLog.Println("hello std")
	zap.L().Info("hello test")
	xlog.Info("hello", zap.String("ss", "hello1"))
	xlog.Error("hello", zap.String("ss", "hello1"))
	xlog.Info("hello %s", xlog.M{
		"test": "ok",
	})
	xlog.InfoW(func(log xlog.Logger) {
		log.Print("ok")
	})
	log.Info("hello", zap.String("ss", "hello1"))
	log.Error("hello", zap.String("ss", "hello1"))
	log.Info("hello", zap.String("ss", "hello1"))
	log.Infof("hello %s", "1234")
	log.Info("hello %s", xlog.M{
		"test": "ok",
	})
	log.Warn("test")
	log.InfoW(func(log xlog.Logger) {
		log.Print("ok")
	})

	initCfgFromJson()

	fmt.Printf("ok-----------------------------------------\n\n")
	stdLog.Println("hello std")
	zap.L().Info("hello test")
	xlog.Info("hello", zap.String("ss", "hello1"))
	xlog.Error("hello", zap.String("ss", "hello1"))
	xlog.Info("hello %s", xlog.M{
		"test": "ok",
	})
	xlog.InfoW(func(log xlog.Logger) {
		log.Print("ok")
	})
	log.Info("hello", zap.String("ss", "hello1"))
	log.Error("hello", zap.String("ss", "hello1"))
	log.Info("hello", zap.String("ss", "hello1"))
	log.Infof("hello %s", "1234")
	log.Info("hello %s", xlog.M{
		"test": "ok",
	})
	log.Warn("test")
	log.InfoW(func(log xlog.Logger) {
		log.Print("ok")
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
        "encoding": "console",
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
        "initialFields": {"hello":"world"},
		"filterSuffix":["test.12345"]
}`

	var cfg1 xlog_config.Config
	xerror.Exit(json.Unmarshal([]byte(cfg), &cfg1))
	cfg1.Encoding = "console"
	//cfg1.Encoding = "json"
	zl, err := cfg1.Build()

	xerror.Exit(err)
	xerror.Exit(xlog.SetDefault(zl))
	log.Warn("test")
}
