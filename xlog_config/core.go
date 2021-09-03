package xlog_config

import (
	"strings"
	"sync"

	"go.uber.org/zap/zapcore"
)

var globalMutex sync.RWMutex

func SetGlobalLevel(l zapcore.Level) {
	if globalLevel == nil {
		return
	}
	globalLevel.SetLevel(l)
}

func hasPrefix(name string) bool {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	for key := range globalPrefix {
		if strings.HasPrefix(name, key) {
			return true
		}
	}
	return false
}

func hasSuffix(name string) bool {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	for key := range globalSuffix {
		if strings.HasSuffix(name, key) {
			return true
		}
	}
	return false
}

func DelPrefix(key string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	delete(globalPrefix, key)
}

func SetPrefix(key string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	globalPrefix[key] = struct{}{}
}

func SetSuffix(key string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	globalSuffix[key] = struct{}{}
}

func DelSuffix(key string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	delete(globalSuffix, key)
}

var _ zapcore.Core = (*filterCore)(nil)

type filterCore struct {
	zapcore.Core
}

func (t *filterCore) xlog() {}

func (t *filterCore) With(fields []zapcore.Field) zapcore.Core {
	var _, ok = t.Core.(interface{ xlog() })
	if ok {
		return t.Core.With(fields)
	}

	return &filterCore{Core: t.Core.With(fields)}
}

func (t *filterCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	if hasPrefix(ent.LoggerName) || hasSuffix(ent.LoggerName) {
		return nil
	}

	return t.Core.Write(ent, fields)
}

func (t *filterCore) Check(ent zapcore.Entry, cc *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if hasPrefix(ent.LoggerName) || hasSuffix(ent.LoggerName) {
		cc.AddCore(ent, t)
		return cc
	}

	return t.Core.Check(ent, cc)
}
