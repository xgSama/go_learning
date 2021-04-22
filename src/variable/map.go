package variable

import "fmt"

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/21 14:39
 */

func createMap() {

	// 方式一
	var map1 map[string]string
	// 在使用map前，需要先使用make分配空间
	map1 = make(map[string]string, 1)
	map1["one"] = "java"
	map1["two"] = "c++"
	map1["three"] = "go"
	fmt.Println(map1)

	// 方式二
	map2 := make(map[int]string)
	map2[1] = "java"
	map2[2] = "c++"
	map2[3] = "go"
	fmt.Println(map2)

	// 方式三
	map3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "go",
	}
	fmt.Println(map3)

	// 遍历
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	// 删除
	delete(map1, "one")
	fmt.Println(map1)

}
