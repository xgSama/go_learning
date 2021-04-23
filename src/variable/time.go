package variable

import (
	"fmt"
	"time"
)

/**
 * @Author xgSama
 * @Description:
 * @Date 2021/4/23 10:19
 */

func testTimeFunc() {
	now := time.Now()
	fmt.Printf("now = %v\n", now)
	fmt.Println(now.Year(), "年", int(now.Month()), "月", now.Day(), "日", now.Hour(), "时", now.Minute(), "分", now.Second(), "秒")
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Printf("%T\n", 100*time.Millisecond)

	// 格式化时间
	fmt.Println(now.Format("2006-01-02 03:04:05"))
}
