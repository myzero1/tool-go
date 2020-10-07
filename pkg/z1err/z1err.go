package z1err

import (
	"fmt"
	"strings"

	pkgerrors "github.com/pkg/errors"
)

// https://www.cnblogs.com/zhangboyu/p/7911190.html

// Check panic error
func Check(e error) {
	if nil != e {
		panic(e)
	}
}

// Handle catching error to return or assign to recoverFunc.
// eg1:defer z1err.Handle(&err)
// eg2:defer z1err.Handle(nil, func(err error) {...}
func Handle(returnErr *error, recoverFunc ...func(err error)) {
	r := recover()
	if nil != r {
		errMsg := fmt.Sprintf(`%s`, r)
		err := pkgerrors.New(errMsg)
		if returnErr != nil {
			*returnErr = err
		} else {
			if len(recoverFunc) > 0 {
				recoverFunc[0](err)
			} else {
				panic("z1error.Handle has not this logic")
			}
		}
	}
}

// StackSkipPrint print msg with stack and can set skip.
// eg:z1err.StackSkipPrint(err)
// eg:z1err.StackSkipPrint(err, 7)
func StackSkipPrint(err error, skip ...int) (str string) {
	// https://studygolang.com/articles/17430?fr=sidebar
	skipFlag := 7
	if len(skip) > 0 {
		skipFlag = skip[0]
	}
	errCause := fmt.Sprintf(`%v`, pkgerrors.Cause(err))
	errInfo := fmt.Sprintf(`%+v`, err)
	errInfo2 := strings.Replace(errInfo, errCause, "", -1)
	errInfoArr := strings.Split(errInfo2, "\n")
	if len(errInfoArr) > skipFlag {
		str = errCause + "\n" + strings.Join(errInfoArr[skipFlag:], "\n")
	} else {
		str = errCause
	}

	return
}

// CheckErr check and add statck to error,or throw a exception.
// eg:CheckErr(err)
// eg:CheckErr(err, true)
// eg:CheckErr(err, true, "this is addon msg")
func CheckErr(e error, opt ...interface{}) (noerr bool, errWithStack error) {
	if e != nil {
		msg := ``
		if len(opt) > 1 {
			myMsg, ok := opt[1].(string)
			if ok {
				msg = myMsg
			}
		}

		errWithStack = pkgerrors.Wrap(e, msg)

		if len(opt) > 0 {
			isPanic, ok := opt[0].(bool)
			if ok && isPanic {
				panic(errWithStack)
				// panic(e.Error())
			}
		}

		noerr = false
	} else {
		noerr = true
	}

	return
}
