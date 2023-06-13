package error

import "fmt"

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("错误代码：%d，错误信息：%s", e.Code, e.Msg)
}
