package sys

import (
	"github.com/greenblat17/platform-common/pkg/sys/codes"
	"github.com/pkg/errors"
)

type commonError struct {
	msg  string
	code codes.Code
}

func NewCommonError(msg string, code codes.Code) *commonError {
	return &commonError{msg: msg, code: code}
}

func (e *commonError) Error() string {
	return e.msg
}

func (e *commonError) Code() codes.Code {
	return e.code
}

func IsCommonError(err error) bool {
	var ce *commonError
	return errors.As(err, &ce)
}

func GetCommonError(err error) *commonError {
	var ce *commonError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
