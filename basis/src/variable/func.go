package variable

import "fmt"

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/22 11:06
 */
func printFbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		res := printFbn(n-1) + printFbn(n-2)
		return res
	}
}

// f(1) = 3 f(n)=2*f(n-1)+1
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

// 匿名函数
func testNi() {
	res := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	println(res)

	a := func(n1, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 23)
	println(res2)

	fmt.Printf("a的类型是 %T", a)

}
