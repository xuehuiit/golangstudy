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

func StrutssTest() {

}

func strutstest() {

	type person struct {
		name    string
		age     int
		address string
	}

	var person_tim person

	person_tim.name = "robert"
	person_tim.age = 123
	person_tim.address = "长临路1318弄32号1202室内"

	fmt.Println(person_tim.name)
}
