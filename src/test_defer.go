package main

import "fmt"

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/21 14:26
 */

func main() {
	returnAndDefer()
}

func returnAndDefer() int {
	defer deferFunc()

	return returnFunc()
}

func returnFunc() int {
	fmt.Println("return function")
	return 0
}

func deferFunc() int {
	fmt.Println("defer function")
	return 0
}
