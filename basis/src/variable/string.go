package variable

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/23 9:38
 */

func testStringFunc() {
	str := "i love china，哈哈"
	// go统一使用utf8编码，数字字母占1字节，中文及符号占3字节
	strLen := len(str) // 21
	fmt.Println(str, strLen)

	// 中文被截断产生乱码
	for i := 0; i < strLen; i++ {
		fmt.Printf("%c\t", str[i])
	}
	println()
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c\t", str1[i])
	}
	println()

	for _, s := range str {
		fmt.Printf("%c\t", s)
	}
	println()

	// string 转 字符串
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果是", n)
	}

	// 整数转字符串
	i := strconv.Itoa(123)
	fmt.Printf("%v, %T \n", i, i)

	bytes := []byte("hello go")
	fmt.Printf("bytes = %v \n", bytes)

	str = string(bytes)
	fmt.Println("bytes => ", str)

	// 进制转换
	fmt.Println("123对应的二进制是：", strconv.FormatInt(123, 2))
	fmt.Println("123对应的八进制是：", strconv.FormatInt(123, 8))
	fmt.Println("123对应的十六进制是：", strconv.FormatInt(123, 16))

	println(strings.Contains("seafood", "foo"))
	println(strings.Count("seafood", "o"))
	println(strings.EqualFold("asd", "ASD"))
}
