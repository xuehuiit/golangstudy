package study

import (
	"fmt"
)

func FunctionTest() {

	functionPram(12, 13.333333)

	df := functionReturn(13134, 13.88888)

	fmt.Println("  return value %d ", df)

}

func functionPram(int_parm int, float_parm float32) {

	fmt.Printf(" 传入的参数是：   %d ,   %f \n ", int_parm, float_parm)

}

func functionReturn(int_parm int, float_parm float32) int {

	return int_parm

}
