package study

import (
	"fmt"
)

const const_int int = 12341234
const Const_string string = "dddd"

var Const_string_change string = " my can be change "

func Vardefin() {

	//整数
	var int_a = 1

	var int_b int
	int_b = 13

	int_c := 15

	fmt.Printf("数据的值为： %d  %d \n", int_a+int_b, int_c)

	//浮点数

	var float_a float32
	float_a = 123.33333

	fmt.Printf(" 浮点数： %f  \n", float_a)

	//字符串和字符串
	var string_a = "我是字符串"
	var string_b string
	string_b = "我也是字符串"

	fmt.Printf(" 字符串  %s 和字符串  %s \n", string_a, string_b)

	//布尔类型变量

	var boolean_a bool
	boolean_a = true

	fmt.Printf(" 布尔值   %b  \n", boolean_a)

	//常量
	fmt.Printf(" 常量   %s \n ", Const_string)
	fmt.Printf("  全局变量    %s \n ", Const_string_change)

	//指针

}
