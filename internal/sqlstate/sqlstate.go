package sqlstate

import (
	"regexp"

	"github.com/desulaidovich/pretty-link/internal/fail"
)

func ErrNo(sqlSate error) *fail.Fail {
	re, _ := regexp.Compile(`.+(SQLSTATE\s(\d+))`)

	errno := re.FindStringSubmatch(sqlSate.Error())[2]

	list := map[string]*fail.Fail{
		"23505": fail.New(fail.AccountIsExists),
		"23514": fail.New(fail.InvalidRequestJSON),
	}

	if key, ok := list[errno]; ok {
		return key
	}

	return fail.New(fail.Undefined)
}
