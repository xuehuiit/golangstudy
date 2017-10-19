/**
  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

  go语言学习-基础语法

 */
package study

import (
	"fmt"
)

func TestFlow() {

	Ifelseflow()
	Whileflow()

}

// 条件语句

func Ifelseflow() {

	var falg int = -1

	if falg > 0 {
		fmt.Println("print herer \n")
	} else {
		fmt.Printf("  it is > %d  \n", falg)
	}

}

//循环语句

func Whileflow() {

	for a := 0; a < 10; a++ {

		fmt.Printf(" 测试循环 %d  \n", a)
	}

	//遍历数组

	array_int := [6]int{1, 2, 3, 4, 5, 6}

	for i, x := range array_int {

		fmt.Printf(" 本次循环的数组是： %d %d \n", i, x)

	}

	numbers := [6]int{1, 2, 3, 5}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

}
