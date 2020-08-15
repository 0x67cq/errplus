package errplus
import (
    "fmt"
)

type ErrStack struct {
    stackInfo string
    params    []interface{}
    err error
}

func (e *ErrStack) Error() string {
    return fmt.Sprintf("stackInfo: \n%s \nparams: %v,\nerr: %v", e.stackInfo, e.params, e.err)
}

// Unwrap 用于支持 errors.As API
func (e *ErrStack) Unwrap() error {
	return e.err
}
