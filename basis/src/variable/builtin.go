package variable

import "fmt"

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/23 10:54
 */

func testBuiltin() {

	num1 := 100
	fmt.Printf("num1的类型%T,num1的值=%v,num1的地址%v\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2的类型%T,num2的值=%v,num2的地址%v,\n指向的值是%v\n", num2, num2, &num2, *num2)
}
