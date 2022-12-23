package error

import "fmt"

type dycErr struct {
	code   int
	msg    string
	oriErr httpErr
}

type httpErr struct {
	code int
	msg  string
}

func (e dycErr) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", e.code, e.msg)
}

func NewErr(code int, msg string, oriCode int, oriMsg string) error {
	return dycErr{
		code: code,
		msg:  msg,
		oriErr: struct {
			code int
			msg  string
		}{code: oriCode, msg: oriMsg},
	}
}

func GetCode(err error) int {
	if e, ok := err.(dycErr); ok {
		return e.code
	}
	return -1
}

func GetMsg(err error) string {
	if e, ok := err.(dycErr); ok {
		return e.msg
	}
	return ""
}
