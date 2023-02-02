package main

import (
	"errors"
	"fmt"
)

var a int = 2020

// 遮蔽预定义标识符
type new int

func getYear() (new, error) {
	var b int16 = 2021
	return new(b), nil
}

func checkYear() error {
	err := errors.New("wrong year")
	// 遮蔽了外层 包代码块 中的包级变量 a
	// 遮蔽了上一行的 err
	switch a, err := getYear(); a {
	case 2020:
		fmt.Println("it is", a, err)
	case 2021:
		fmt.Println("it is", a)
		// 置空的是 getYear 返回的 nil，不是 err := errors.New("wrong year")
		err = nil
	}
	fmt.Println("after check, it is", a)

	return err
}

// it is 2021
// after check, it is 2020
// call checkYear error: wrong year
func main() {
	err := checkYear()

	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}
	fmt.Println("call checkYear ok")
}
