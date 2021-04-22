package main

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/21 10:25
 */

import (
	"errors"
	"fmt"
	"go_learning/src/lib"
	_ "go_learning/src/lib"
)

func main() {


	var g float32 = 3.14
	fmt.Printf("%T", g)

	lib.Lib1Test()
	lib.Lib2Test()

	a := 'a'
	b := 'b'
	fmt.Println(a + b)

	var i int = 10
	fmt.Println("i的地址", &i)

	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v\n", &ptr)
	fmt.Printf("ptr 指向的值=%v\n", *ptr)

	// variables()
	// test_iota()
	// testArray()
	// testSlice()

	// myfunc("aa", "nn", "dd", `kk`)

}

func variables() {
	// 定义变量
	var variableName int8 = 99
	var v1, v2, v3 string = "1", "2", "3"
	var enabled, disabled = true, false
	// := 是一个声明语句，左边无新的变量会报错
	// v1, v2 := "11", "21" 这种会报错
	v1, v2, v6 := "11", "21", 9
	v4, v5 := "abc", "aaa"

	var (
		vv = 1
		jj = true
	)
	println(vv, jj)

	const Pi float32 = 3.1415926

	println(variableName, v1, v2, v3, v4, v5, v6, enabled, disabled, Pi)

	// 数值类型
	// `rune`, `int8`, `int16`, `int32`, `int64`和`byte`, `uint8`, `uint16`, `uint32`, `uint64`。
	// 其中`rune`是`int32`的别称，`byte`是`uint8`的别称
	var num1 int = 7
	println(num1)
	// 尽管int的长度是32 bit, 但int 与 int32并不可以互用
	/* var num2 int32 = num1 */

	// 复数 `complex128`（64位实数+64位虚数） 也有`complex64`(32位实数+32位虚数)
	var c complex64 = 5 + 1i
	var c1 complex64 = 5 - 1i
	fmt.Println(real(c), imag(c))
	fmt.Println(c * c1)

	// 字符串
	var str string = `sss`
	var str2 = "ssss"
	println(str, str2)
	// 修改字符串
	s := "hello"
	ss := []byte(s) // 将字符串 s 转换为 []byte 类型
	ss[0] = 'c'
	s2 := string(ss) // 再转换回 string 类型
	fmt.Printf("%s\n", s2)

	sw := "hello"
	sw = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", sw)

	// error
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

}

func testArray() {
	var arr [10]int                                 // 声明了一个int类型的数组
	arr[0] = 42                                     // 数组下标是从0开始的
	arr[1] = 13                                     // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])  //返回未赋值的最后一个元素，默认返回0

	a := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b := [10]int{1, 2, 3}  // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	fmt.Println(a, b, c)

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	// ...只能存在于第一个[]内
	easyArray1 := [...][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray, easyArray, easyArray1)
}

func testSlice() {
	// `slice`是引用类型
	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g

	fmt.Println(aSlice, bSlice)
}

func myfunc(args ...string) {
	for i, arg := range args {
		fmt.Println(i, arg)
	}
}
