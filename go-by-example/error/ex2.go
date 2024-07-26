package main

import (
	"errors"
	"fmt"
	"log"
)

type MyError struct {
	Code int
	Msg  string
}

// 第一优先级 【实现了error 接口】
func (m *MyError) Error() string {
	return fmt.Sprintf(" Error : %s %d", m.Msg, m.Code)
	//return fmt.Errorf(" errors() : %s %d", m.Msg, m.Code)
}

//第二优先级 【实现了fmt.Stringer接口】

func (m MyError) String() string {
	return fmt.Sprintf(" 错误情况 : %s %d", m.Msg, m.Code)
	//fmt.Stringer()
}

func test() error {
	return &MyError{
		100, "ccc",
	}
}
func main() {
	//打印本身；第三优先级
	var err = &MyError{
		Code: 100,
		Msg:  "some nonsense error !",
	}

	var cc = &MyError{}

	var _ error = cc
	var _ error = err

	log.Println(err)
	//fmt.Println(err)

	// 注意 err 和cc类型要一样； 第二个参数必须去地址；类似于反序列化；否则结果不符合预期
	if errors.As(err, &cc) {
		fmt.Println(cc.Code, cc.Msg)
	}

	anoErr := test()
	t, ok := anoErr.(*MyError)
	if ok {
		log.Println("ok", t)
	} else {
		log.Println("......")
	}
}
