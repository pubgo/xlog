package xlog_config

import (
	"encoding/json"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type option func(opts *Config)

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

type Config struct {
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

func (t Config) handleOpts(opts ...option) Config {
	for _, opt := range opts {
		opt(&t)
	}
	return t
}

func (t Config) toZapLogger() (_ *zap.Logger, err error) {
	defer xerror.RespErr(&err)

	zapCfg := zap.Config{}
	xerror.Panic(json.Unmarshal(xerror.PanicBytes(json.Marshal(&t)), &zapCfg))

	key := t.EncoderConfig.EncodeLevel
	key = internal.If(key != "", key, defaultKey).(string)
	zapCfg.EncoderConfig.EncodeLevel = levelEncoder[key]

	key = t.EncoderConfig.EncodeTime
	key = internal.If(key != "", key, defaultKey).(string)
	zapCfg.EncoderConfig.EncodeTime = timeEncoder[key]

	key = t.EncoderConfig.EncodeDuration
	key = internal.If(key != "", key, defaultKey).(string)
	zapCfg.EncoderConfig.EncodeDuration = durationEncoder[key]

	key = t.EncoderConfig.EncodeCaller
	key = internal.If(key != "", key, defaultKey).(string)
	zapCfg.EncoderConfig.EncodeCaller = callerEncoder[key]

	key = t.EncoderConfig.EncodeName
	key = internal.If(key != "", key, defaultKey).(string)
	zapCfg.EncoderConfig.EncodeName = nameEncoder[key]

	return xerror.PanicErr(zapCfg.Build()).(*zap.Logger), nil
}

func NewZapLogger(conf Config, opts ...option) (*zap.Logger, error) {
	return conf.handleOpts(opts...).toZapLogger()
}

func NewDevConfig(opts ...option) Config {
	cfg := Config{
		Level:       "debug",
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
			EncodeCaller:   "full",
			LineEnding:     zapcore.DefaultLineEnding,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return cfg.handleOpts(opts...)
}

func NewProdConfig(opts ...option) Config {
	cfg := Config{
		Level:       "info",
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

	return cfg.handleOpts(opts...)
}
