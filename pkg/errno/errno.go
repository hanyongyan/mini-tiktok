package errno

import (
	"errors"
	"fmt"
)

const (
	ErrCode_SuccessCode                int64 = 0
	ErrCode_ServiceErrCode             int64 = 10001
	ErrCode_ParamErrCode               int64 = 10002
	ErrCode_UserAlreadyExistErrCode    int64 = 10003
	ErrCode_AuthorizationFailedErrCode int64 = 10004
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(ErrCode_SuccessCode, "Success")
	ServiceErr             = NewErrNo(ErrCode_ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ErrCode_ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(ErrCode_UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(ErrCode_AuthorizationFailedErrCode, "Authorization failed")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
