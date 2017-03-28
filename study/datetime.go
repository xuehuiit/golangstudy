//时间日期类型的
/**
 *
 *  http://studygolang.com/articles/240  相关参考方法
 *
 */

package study

import (
	"fmt"
	"time"
)

func TestDate() {
	dates()
}

func dates() {

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	str_time := time.Unix(1389058332, 0).Format("2006-01-02")
	fmt.Println(str_time)

	//字符串转化成时间
	the_time := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)
	unix_time := the_time.Unix()
	fmt.Println(unix_time)

	the_time, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println(unix_time)
	}

	//时间戳
	t := time.Now().Unix()
	fmt.Println(t)

	//时间戳到具体显示的转化
	fmt.Println(time.Unix(t, 0).String())

	//带纳秒的时间戳
	t = time.Now().UnixNano()
	fmt.Println(t)
	fmt.Println("------------------")

	//基本格式化的时间表示
	fmt.Println(time.Now().String())

	fmt.Println(time.Now().Format("2006year 01month 02day"))

	//时间加减

	baseTime := time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
	date := baseTime.Add(1722*7*24*time.Hour + 24*time.Hour + 66355*time.Second)
	fmt.Println(date)

}
