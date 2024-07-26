package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

// 实现 error 接口
func (e MyError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func main() {
	// 模拟一个可能失败的操作
	err := MyError{Code: 404, Message: "Not Found"}

	// 使用 errors.As 检查错误是否是 MyError 类型
	var myErr MyError
	if errors.As(err, &myErr) {
		fmt.Println("Caught a MyError:", myErr.Error())
		// 可以访问 myErr 的字段进行进一步处理
	} else {
		fmt.Println("Caught a different error:", err)
	}
}
