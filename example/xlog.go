package main

import (
	"encoding/json"
	"fmt"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xlog/xlog_config"
	"go.uber.org/zap"
	stdLog "log"
)

var log = xlog.GetLogger("test.12345", zap.Fields(zap.String("name", "test_hello")))

func main() {
	xlog.ErrWith("hello-error", fmt.Errorf("hello error"))
	stdLog.Println("hello std")
	zap.L().Info("hello test")
	xlog.Info("hello", zap.String("ss", "hello1"))
	xlog.Debug("hello", zap.String("ss", "hello1"))
	xlog.Info(fmt.Sprintf)
	xlog.Info(fmt.Sprintf, zap.Logger{})
	xlog.Info(fmt.Sprintf, xerror.Wrap(xerror.Fmt("hello")))
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

	log.Named("kkkkkk").Info("hello")
	log.Named("123").Info("hello")
	log.Named("456").Info("hello")

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
	log.Info("hello1", zap.String("ss", "hello1"))
	log.Error("hello1", zap.String("ss", "hello1"))
	log.Info("hello1", zap.String("ss", "hello1"))
	log.Infof("hello1 %s", "1234")
	log.Info("hello1 %s", xlog.M{
		"test1": "ok",
	})
	log.Warn("test1")
	log.InfoW(func(log xlog.Logger) {
		log.Print("ok")
	})

	xlog_config.GlobalLevel(zap.InfoLevel)
	log.Named("aakkkkkk").Info("hello")
	log.Named("aa123").Info("hello")
	log.Named("aa456").Info("hello")
}

func initCfgFromJson() {
	defer xerror.RespExit()
	cfg := `{
        "level": "error",
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
        "initialFields": {"hello":"world"}
}`

	var cfg1 xlog_config.Config
	xerror.Exit(json.Unmarshal([]byte(cfg), &cfg1))
	cfg1.Encoding = "console"
	//cfg1.Encoding = "json"
	zl := cfg1.Build("hello")
	xerror.Exit(xlog.SetDefault(zl))
	zl.Info("test config")
}
