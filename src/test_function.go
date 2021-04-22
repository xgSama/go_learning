package main

import "fmt"

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/21 14:05
 */

func main() {
	add, mul := SumAndProduct(11, 44)
	fmt.Println("add=", add, "mul=", mul)
}



func SumAndProduct(A, B int) (add, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}
