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
