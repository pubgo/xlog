package xlog_config

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

func SetGlobalLevel(l zapcore.Level) {
	if globalLevel != nil {
		globalLevel.SetLevel(l)
	}
}

var _ zapcore.Core = (*filterCore)(nil)

type filterCore struct {
	filterPrefix []string
	filterSuffix []string
	zapcore.Core
}

func (t *filterCore) Enabled(lvl zapcore.Level) bool {
	return t.Core.Enabled(lvl)
}

func (t *filterCore) With(fields []zapcore.Field) zapcore.Core {
	return &filterCore{Core: t.Core.With(fields), filterPrefix: t.filterPrefix, filterSuffix: t.filterSuffix}
}

func (t *filterCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	for i := range t.filterPrefix {
		if strings.HasPrefix(ent.LoggerName, t.filterPrefix[i]) {
			return nil
		}
	}

	for i := range t.filterSuffix {
		if strings.HasSuffix(ent.LoggerName, t.filterSuffix[i]) {
			return nil
		}
	}

	return t.Core.Write(ent, fields)
}

func (t *filterCore) Check(ent zapcore.Entry, cc *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	var filter bool
	for i := range t.filterPrefix {
		if strings.HasPrefix(ent.LoggerName, t.filterPrefix[i]) {
			filter = true
			cc.AddCore(ent, t)
		}
	}

	for i := range t.filterSuffix {
		if strings.HasSuffix(ent.LoggerName, t.filterSuffix[i]) {
			filter = true
			cc.AddCore(ent, t)
		}
	}

	if filter {
		return cc
	}

	return t.Core.Check(ent, cc)
}
