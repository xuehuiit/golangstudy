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

func PointTest() {
	pointtests()

}

func pointtests() {

	var a int = 20
	var a_point *int

	a_point = &a

	fmt.Printf("变量   a 的值  为：  %d   \n", a)

	fmt.Printf("指针 a_point的地址 %x \n", a_point)

	fmt.Printf("指针 a_point的值   %d \n", *a_point)
}
