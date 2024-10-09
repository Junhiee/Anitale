package errx

import "fmt"

type CustomError struct {
	code uint32
	msg  string
}

func (c *CustomError) GetErrCode() uint32 {
	return c.code
}

func (c *CustomError) GetErrMsg() string {
	return c.msg
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", c.code, c.msg)
}

func (c *CustomError) NewCustomCode(code uint32) *CustomError {
	return &CustomError{
		code: code,
		msg:  GetMessage(code),
	}
}

func (c *CustomError) NewCustomMsg(msg string) *CustomError {
	return &CustomError{
		code: c.code,
		msg:  msg,
	}
}

func NewCustomError(code uint32, msg string) *CustomError {
	return &CustomError{
		code: code,
		msg:  msg,
	}
}
