package xlog_abc

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fields []zap.Field

func (t *Fields) Skip() *Fields { *t = append(*t, zap.Skip()); return t }
func (t *Fields) Binary(key string, val []byte) *Fields {
	*t = append(*t, zap.Binary(key, val))
	return t
}

func (t *Fields) Bool(key string, val bool) *Fields {
	*t = append(*t, zap.Bool(key, val))
	return t
}

func (t *Fields) ByteString(key string, val []byte) *Fields {
	*t = append(*t, zap.ByteString(key, val))
	return t
}

func (t *Fields) Complex128(key string, val complex128) *Fields {
	*t = append(*t, zap.Complex128(key, val))
	return t
}

func (t *Fields) Complex64(key string, val complex64) *Fields {
	*t = append(*t, zap.Complex64(key, val))
	return t
}

func (t *Fields) Float64(key string, val float64) *Fields {
	*t = append(*t, zap.Float64(key, val))
	return t
}

func (t *Fields) Float32(key string, val float32) *Fields {
	*t = append(*t, zap.Float32(key, val))
	return t
}

func (t *Fields) Int(key string, val int) *Fields {
	*t = append(*t, zap.Int(key, val))
	return t
}

func (t *Fields) Int64(key string, val int64) *Fields {
	*t = append(*t, zap.Int64(key, val))
	return t
}

func (t *Fields) Int32(key string, val int32) *Fields {
	*t = append(*t, zap.Int32(key, val))
	return t
}

func (t *Fields) Int16(key string, val int16) *Fields {
	*t = append(*t, zap.Int16(key, val))
	return t
}

func (t *Fields) Int8(key string, val int8) *Fields {
	*t = append(*t, zap.Int8(key, val))
	return t
}

func (t *Fields) String(key string, val string) *Fields {
	*t = append(*t, zap.String(key, val))
	return t
}

func (t *Fields) Msg(format string, a ...interface{}) *Fields {
	*t = append(*t, zap.String("msg", fmt.Sprintf(format, a...)))
	return t
}

func (t *Fields) Uint(key string, val uint) *Fields {
	*t = append(*t, zap.Uint(key, val))
	return t
}

func (t *Fields) Uint64(key string, val uint64) *Fields {
	*t = append(*t, zap.Uint64(key, val))
	return t
}

func (t *Fields) Uint32(key string, val uint32) *Fields {
	*t = append(*t, zap.Uint32(key, val))
	return t
}

func (t *Fields) Uint16(key string, val uint16) *Fields {
	*t = append(*t, zap.Uint16(key, val))
	return t
}

func (t *Fields) Uint8(key string, val uint8) *Fields {
	*t = append(*t, zap.Uint8(key, val))
	return t
}

func (t *Fields) Reflect(key string, val interface{}) *Fields {
	*t = append(*t, zap.Reflect(key, val))
	return t
}

func (t *Fields) Stringer(key string, val fmt.Stringer) *Fields {
	*t = append(*t, zap.Stringer(key, val))
	return t
}

func (t *Fields) Time(key string, val time.Time) *Fields {
	*t = append(*t, zap.Time(key, val))
	return t
}

func (t *Fields) Stack(key string) *Fields {
	*t = append(*t, zap.Stack(key))
	return t
}

func (t *Fields) Duration(key string, val time.Duration) *Fields {
	*t = append(*t, zap.Duration(key, val))
	return t
}

func (t *Fields) Object(key string, val zapcore.ObjectMarshaler) *Fields {
	*t = append(*t, zap.Object(key, val))
	return t
}

func (t *Fields) Any(key string, value interface{}) *Fields {
	*t = append(*t, zap.Any(key, value))
	return t
}
