package xlog_errs

import "github.com/pubgo/xerror"

var (
	Err              = xerror.New("log_default err")
	ErrParamsInValid = Err.New("params is invalid")
)
