package variable

import (
	"errors"
	"fmt"
)

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/23 11:10
 */

func testError() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("x的类型是%T\n", err)
			fmt.Println("error =", err)
		}
	}()

	num1 := 100
	num2 := 0
	res := num1 / num2
	fmt.Printf("res = %v", res)
}

func testCustomError(name string) (res string, err error) {

	if len(name) < 3 {
		return "yes", nil
	} else {
		return name, errors.New("too long")
	}

}
