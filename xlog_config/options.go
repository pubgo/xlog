package xlog_config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Hooks registers functions which will be called each time the Logger writes
// out an Entry. Repeated use of Hooks is additive.
//
// Hooks are useful for simple side effects, like capturing metrics for the
// number of emitted logs. More complex side effects, including anything that
// requires access to the Entry's structured fields, should be implemented as
// a zapcore.Core instead. See zapcore.RegisterHooks for details.
func WithHooks(hooks ...func(zapcore.Entry) error) Option {
	return func(opts *config) {
		opts.zapOpts = append(opts.zapOpts, zap.Hooks(hooks...))
	}
}

// Fields adds fields to the Logger.
func WithFields(fs ...zap.Field) Option {
	return func(opts *config) {
		opts.zapOpts = append(opts.zapOpts, zap.Fields(fs...))
	}
}

// ErrorOutput sets the destination for errors generated by the Logger. Note
// that this option only affects internal errors; for sample code that sends
// error-level logs to a different location from info- and debug-level logs,
// see the package-level AdvancedConfiguration example.
//
// The supplied WriteSyncer must be safe for concurrent use. The Open and
// zapcore.Lock functions are the simplest ways to protect files with a mutex.
func WithErrorOutput(w zapcore.WriteSyncer) Option {
	return func(opts *config) {
		opts.zapOpts = append(opts.zapOpts, zap.ErrorOutput(w))
	}
}

// Development puts the logger in development mode, which makes DPanic-level
// logs panic instead of simply logging an error.
func WithDevelopment() Option {
	return func(opts *config) {
		opts.zapOpts = append(opts.zapOpts, zap.Development())
	}
}

// AddCaller configures the Logger to annotate each message with the filename
// and line number of zap's caller.
func WithCaller() Option {
	return func(opts *config) {
		opts.zapOpts = append(opts.zapOpts, zap.AddCaller())
	}
}

// AddCallerSkip increases the number of callers skipped by caller annotation
// (as enabled by the AddCaller option). When building wrappers around the
// Logger and SugaredLogger, supplying this Option prevents zap from always
// reporting the wrapper code as the caller.
func WithCallerSkip(skip int) Option {
	return func(opts *config) {
		opts.zapOpts = append(opts.zapOpts, zap.AddCallerSkip(skip))
	}
}

// AddStacktrace configures the Logger to record a stack trace for all messages at
// or above a given level.
func WithStacktrace(lvl zapcore.LevelEnabler) Option {
	return func(opts *config) {
		opts.zapOpts = append(opts.zapOpts, zap.AddStacktrace(lvl))
	}
}

// WithEncoding ...
func WithEncoding(enc string) Option {
	return func(opts *config) {
		opts.Encoding = enc
	}
}

func WithLevel(ll Level) Option {
	return func(opts *config) {
		opts.Level = zap.NewAtomicLevelAt(ll)
	}
}