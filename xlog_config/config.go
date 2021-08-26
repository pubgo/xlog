package xlog_config

import (
	"encoding/json"
	"sort"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLevel *zap.AtomicLevel

type option func(opts *Config)

type encoderConfig struct {
	MessageKey       string `json:"messageKey" yaml:"messageKey" toml:"messageKey"`
	LevelKey         string `json:"levelKey" yaml:"levelKey" toml:"levelKey"`
	TimeKey          string `json:"timeKey" yaml:"timeKey" toml:"timeKey"`
	NameKey          string `json:"nameKey" yaml:"nameKey" toml:"nameKey"`
	CallerKey        string `json:"callerKey" yaml:"callerKey" toml:"callerKey"`
	StacktraceKey    string `json:"stacktraceKey" yaml:"stacktraceKey" toml:"stacktraceKey"`
	LineEnding       string `json:"lineEnding" yaml:"lineEnding" toml:"lineEnding"`
	EncodeLevel      string `json:"levelEncoder" yaml:"levelEncoder" toml:"levelEncoder"`
	EncodeTime       string `json:"timeEncoder" yaml:"timeEncoder" toml:"timeEncoder"`
	EncodeDuration   string `json:"durationEncoder" yaml:"durationEncoder" toml:"durationEncoder"`
	EncodeCaller     string `json:"callerEncoder" yaml:"callerEncoder" toml:"callerEncoder"`
	EncodeName       string `json:"nameEncoder" yaml:"nameEncoder" toml:"nameEncoder"`
	ConsoleSeparator string `json:"consoleSeparator" yaml:"consoleSeparator"`
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
	FilterPrefix      []string               `json:"filterPrefix" yaml:"filterPrefix" toml:"filterPrefix"`
	FilterSuffix      []string               `json:"filterSuffix" yaml:"filterSuffix" toml:"filterSuffix"`
}

func (t Config) handleOpts(opts ...option) Config {
	for _, opt := range opts {
		opt(&t)
	}
	return t
}

func (t Config) Build(opts ...zap.Option) (_ *zap.Logger, err error) {
	defer xerror.RespErr(&err)

	zapCfg := zap.Config{}
	var dt = xerror.PanicBytes(json.Marshal(&t))
	xerror.Panic(json.Unmarshal(dt, &zapCfg))

	// 保留全局log level
	globalLevel = &zapCfg.Level

	key := internal.Default(t.EncoderConfig.EncodeLevel, defaultKey)
	zapCfg.EncoderConfig.EncodeLevel = levelEncoder[key]

	key = internal.Default(t.EncoderConfig.EncodeTime, defaultKey)
	zapCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(key)
	var te, ok = timeEncoder[key]
	if ok {
		zapCfg.EncoderConfig.EncodeTime = te
	}

	key = internal.Default(t.EncoderConfig.EncodeDuration, defaultKey)
	zapCfg.EncoderConfig.EncodeDuration = durationEncoder[key]

	key = internal.Default(t.EncoderConfig.EncodeCaller, defaultKey)
	zapCfg.EncoderConfig.EncodeCaller = callerEncoder[key]

	key = internal.Default(t.EncoderConfig.EncodeName, defaultKey)
	zapCfg.EncoderConfig.EncodeName = nameEncoder[key]

	var log = xerror.PanicErr(zapCfg.Build(opts...)).(*zap.Logger)
	log = log.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return &filterCore{Core: core, filterPrefix: t.FilterPrefix, filterSuffix: t.FilterSuffix}
	}))

	return log, nil
}

func NewDevConfig(opts ...option) Config {
	cfg := Config{
		Level:             "debug",
		Development:       true,
		Encoding:          "console",
		DisableStacktrace: true,
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
		Encoding:          "json",
		DisableStacktrace: true,
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

// MergeOutputPaths merges logging output paths, resolving conflicts.
func MergeOutputPaths(cfg zap.Config) zap.Config {
	outputs := make(map[string]struct{})
	for _, v := range cfg.OutputPaths {
		outputs[v] = struct{}{}
	}
	outputSlice := make([]string, 0)
	if _, ok := outputs["/dev/null"]; ok {
		// "/dev/null" to discard all
		outputSlice = []string{"/dev/null"}
	} else {
		for k := range outputs {
			outputSlice = append(outputSlice, k)
		}
	}
	cfg.OutputPaths = outputSlice
	sort.Strings(cfg.OutputPaths)

	errOutputs := make(map[string]struct{})
	for _, v := range cfg.ErrorOutputPaths {
		errOutputs[v] = struct{}{}
	}
	errOutputSlice := make([]string, 0)
	if _, ok := errOutputs["/dev/null"]; ok {
		// "/dev/null" to discard all
		errOutputSlice = []string{"/dev/null"}
	} else {
		for k := range errOutputs {
			errOutputSlice = append(errOutputSlice, k)
		}
	}
	cfg.ErrorOutputPaths = errOutputSlice
	sort.Strings(cfg.ErrorOutputPaths)

	return cfg
}
