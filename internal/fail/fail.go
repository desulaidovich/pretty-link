package fail

import (
	"fmt"
)

type Fail struct {
	code int
	data *UserErrorData
}

func New(code int) *Fail {
	return &Fail{
		code: code,
		data: MessageByID(code),
	}
}

func (f *Fail) Error() string {
	return fmt.Sprintf("code: %d: status code: %d message: %s", f.code, f.data.httpStatusCode, f.data.message)
}

func (f *Fail) Code() int {
	return f.code
}

func (f *Fail) HTTPStatusCode() int {
	return f.data.httpStatusCode
}

func (f *Fail) Message() string {
	return f.data.message
}
