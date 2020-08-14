package log

import (
	"go.uber.org/zap"
)

type XLog = xlog

func NewXLog(zl *zap.Logger) *xlog {
	xl := &xlog{}
	xl.zl = zl
	return xl
}
