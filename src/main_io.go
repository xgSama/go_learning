package main

import (
	"fmt"
	_ "go_learning/src/lib"
	"strings"
)

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/22 10:39
 */

var (
	Fun1 = func(n1, n2 int) int {
		println("SSS3333")
		return n1 * n2
	}
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	//var name string
	//_, _ = fmt.Scanln(&name)
	//fmt.Println(name)
	//var age int
	//_, _ = fmt.Scanf("%s %ds", &name, &age)
	//fmt.Println(name, age)

	f:= makeSuffix(".jpg")
	fmt.Println(f("winter"))

	println(Fun1(2, 2))

}
