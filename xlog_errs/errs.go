package xlog_errs

import "github.com/pubgo/xerror"

var (
	Err              = xerror.New("xlog")
	ErrParamsInValid = Err.New("params is invalid")
)
