package errors

import (
	"fmt"
)

type SDKError struct {
	Code int
	Msg  string
}

func (err *SDKError) Error() string {
	return fmt.Sprintf("[SDKError] code=%d, msg=%s", err.Code, err.Msg)
}

func NewSDKError(code int, msg string) error {
	return &SDKError{
		Code: code,
		Msg:  msg,
	}
}

func (err *SDKError) GetCode() int {
	return err.Code
}

func (err *SDKError) GetMsg() string {
	return err.Msg
}
