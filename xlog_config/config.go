package xlog_config

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/xlog_errs"
)

type Config config
type Option func(opts *config)
type config struct {
	Level             string                 `json:"level" yaml:"level" toml:"level"`
	Development       bool                   `json:"development" yaml:"development" toml:"development"`
	DisableCaller     bool                   `json:"disableCaller" yaml:"disableCaller" toml:"disableCaller"`
	DisableStacktrace bool                   `json:"disableStacktrace" yaml:"disableStacktrace" toml:"disableStacktrace"`
	Sampling          *samplingConfig        `json:"sampling" yaml:"sampling" toml:"sampling"`
	Encoding          string                 `json:"encoding" yaml:"encoding" toml:"encoding"`
	EncoderConfig     encoderConfig          `json:"encoderConfig" yaml:"encoderConfig" toml:"encoderConfig"`
	OutputPaths       []string               `json:"outputPaths" yaml:"outputPaths" toml:"outputPaths"`
	ErrorOutputPaths  []string               `json:"errorOutputPaths" yaml:"errorOutputPaths" toml:"errorOutputPaths"`
	InitialFields     map[string]interface{} `json:"initialFields" yaml:"initialFields" toml:"initialFields"`
}

type encoderConfig struct {
	MessageKey     string `json:"messageKey" yaml:"messageKey" toml:"messageKey"`
	LevelKey       string `json:"levelKey" yaml:"levelKey" toml:"levelKey"`
	TimeKey        string `json:"timeKey" yaml:"timeKey" toml:"timeKey"`
	NameKey        string `json:"nameKey" yaml:"nameKey" toml:"nameKey"`
	CallerKey      string `json:"callerKey" yaml:"callerKey" toml:"callerKey"`
	StacktraceKey  string `json:"stacktraceKey" yaml:"stacktraceKey" toml:"stacktraceKey"`
	LineEnding     string `json:"lineEnding" yaml:"lineEnding" toml:"lineEnding"`
	EncodeLevel    string `json:"levelEncoder" yaml:"levelEncoder" toml:"levelEncoder"`
	EncodeTime     string `json:"timeEncoder" yaml:"timeEncoder" toml:"timeEncoder"`
	EncodeDuration string `json:"durationEncoder" yaml:"durationEncoder" toml:"durationEncoder"`
	EncodeCaller   string `json:"callerEncoder" yaml:"callerEncoder" toml:"callerEncoder"`
	EncodeName     string `json:"nameEncoder" yaml:"nameEncoder" toml:"nameEncoder"`
}

type samplingConfig struct {
	Initial    int `json:"initial" yaml:"initial" toml:"initial"`
	Thereafter int `json:"thereafter" yaml:"thereafter" toml:"thereafter"`
}

func (t config) toZapLogger() (_ *zap.Logger, err error) {
	defer xerror.RespErr(&err)

	zapCfg := zap.Config{}
	xerror.Panic(json.Unmarshal(xerror.PanicBytes(json.Marshal(&t)), &zapCfg))

	var ok bool

	if zapCfg.EncoderConfig.EncodeLevel, ok = levelEncoder[t.EncoderConfig.EncodeLevel]; !ok {
		if t.EncoderConfig.EncodeLevel != "" {
			xerror.PanicF(xlog_errs.ErrParamsInValid, "EncodeLevel: %s", t.EncoderConfig.EncodeLevel)
		}
		zapCfg.EncoderConfig.EncodeLevel = levelEncoder[defaultKey]
	}

	if zapCfg.EncoderConfig.EncodeTime, ok = timeEncoder[t.EncoderConfig.EncodeTime]; !ok {
		if t.EncoderConfig.EncodeTime != "" {
			xerror.PanicF(xlog_errs.ErrParamsInValid, "EncodeTime: %s", t.EncoderConfig.EncodeTime)
		}
		zapCfg.EncoderConfig.EncodeTime = timeEncoder[defaultKey]
	}

	if zapCfg.EncoderConfig.EncodeDuration, ok = durationEncoder[t.EncoderConfig.EncodeDuration]; !ok {
		if t.EncoderConfig.EncodeDuration != "" {
			xerror.PanicF(xlog_errs.ErrParamsInValid, "EncodeDuration: %s", t.EncoderConfig.EncodeDuration)
		}
		zapCfg.EncoderConfig.EncodeDuration = durationEncoder[defaultKey]
	}

	if zapCfg.EncoderConfig.EncodeCaller, ok = callerEncoder[t.EncoderConfig.EncodeCaller]; !ok {
		if t.EncoderConfig.EncodeCaller != "" {
			xerror.PanicF(xlog_errs.ErrParamsInValid, "EncodeCaller: %s", t.EncoderConfig.EncodeCaller)
		}
		zapCfg.EncoderConfig.EncodeCaller = callerEncoder[defaultKey]
	}

	if zapCfg.EncoderConfig.EncodeName, ok = nameEncoder[t.EncoderConfig.EncodeName]; !ok {
		if t.EncoderConfig.EncodeName != "" {
			xerror.PanicF(xlog_errs.ErrParamsInValid, "EncodeName: %s", t.EncoderConfig.EncodeName)
		}
		zapCfg.EncoderConfig.EncodeName = nameEncoder[defaultKey]
	}

	return xerror.PanicErr(zapCfg.Build()).(*zap.Logger), nil
}

func NewZapLoggerFromOption(opts ...Option) (_ *zap.Logger, err error) {
	defer xerror.RespErr(&err)

	cfg := config(NewProdConfig())
	for _, opt := range opts {
		opt(&cfg)
	}

	return xerror.PanicErr(cfg.toZapLogger()).(*zap.Logger), nil
}

func NewZapLoggerFromConfig(conf Config, opts ...Option) (_ *zap.Logger, err error) {
	defer xerror.RespErr(&err)

	cfg := config(conf)
	for _, opt := range opts {
		opt(&cfg)
	}

	return xerror.PanicErr(cfg.toZapLogger()).(*zap.Logger), nil
}

func NewZapLoggerFromJson(conf []byte, opts ...Option) (_ *zap.Logger, err error) {
	defer xerror.RespErr(&err)

	var cfg config
	xerror.Panic(json.Unmarshal(conf, &cfg))

	for _, opt := range opts {
		opt(&cfg)
	}

	return xerror.PanicErr(cfg.toZapLogger()).(*zap.Logger), nil
}

func NewDevConfig() Config {
	return Config{
		Level:       DebugLevel.String(),
		Development: true,
		Encoding:    "console",
		EncoderConfig: encoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			EncodeLevel:    "capitalColor",
			EncodeTime:     "iso8601",
			EncodeDuration: "string",
			EncodeCaller:   "default",
			LineEnding:     zapcore.DefaultLineEnding,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewProdConfig() Config {
	return Config{
		Level:       InfoLevel.String(),
		Development: false,
		Sampling: &samplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: encoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    "default",
			EncodeTime:     "default",
			EncodeDuration: "default",
			EncodeCaller:   "default",
			LineEnding:     zapcore.DefaultLineEnding,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
