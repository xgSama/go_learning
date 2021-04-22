package variable

import "fmt"

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/22 10:44
 */

func printShape() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", i, i, i*j)
		}
		fmt.Println()
	}
}

func testGoto() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	goto Here //跳转到Here去
}
