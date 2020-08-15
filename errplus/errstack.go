package errplus

import (
	"runtime/debug"
    "os"
)

func Wrap(err error, params ...interface{}) error {
    flag := os.Getenv("GOLANG_ERROR_STACK_DEBUG")
    if flag != "GOLANG_ERROR_STACK_DEBUG" {
        return err
    }

    e := ErrStack {
        stackInfo: string(debug.Stack()),
        params: params,
        err: err,
    }
    return &e
}
