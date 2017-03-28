package study

import (
	"fmt"
	//"container/list"
	"strconv"
)

type person struct {
	name string
	age  int
	addr string
}

var arr_int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var arr_dy_int []int

var person_arr []person

func CollectionTest() {

	arrtest()
	teststructarr()
	maptest(1)

}

/**
 * 数组测试
 */
func arrtest() {

	for a := 0; a < len(arr_int); a++ {

		fmt.Printf("固定长度数组循环  %d ", a)

	}

	arr_dy_int = make([]int, 150)

	for index := 0; index < 150; index++ {
		arr_dy_int[index] = index + 100
	}

	for index1 := 0; index1 < len(arr_dy_int); index1++ {
		fmt.Printf(" 动态数组 - 切片的 循环  %d  \n", arr_dy_int[index1])
	}

}

func teststructarr() {

	person_arr = make([]person, 20)

	for index := 1; index < len(person_arr); index++ {

		var personitem person

		int_str := strconv.Itoa(index)

		personitem.name = "djing " + fmt.Sprintf("%d", index) + "  " + int_str
		personitem.age = index
		person_arr[index] = personitem

	}

	fmt.Println(person_arr)

	for index1 := 1; index1 < len(person_arr); index1++ {

		personitem := person_arr[index1]

		fmt.Println("   " + personitem.name)

	}

}

func listest() {

	//lists := list.New()

}

func maptest(parm int) {

	var keyValueMap map[string]string

	keyValueMap = make(map[string]string)

	keyValueMap["fffffffff"] = "ddddddddddddd"
	keyValueMap["fffffffff1"] = "ddddddddddddd"
	keyValueMap["fffffffff2"] = "ddddddddddddd"
	keyValueMap["fffffffff3"] = "ddddddddddddd"
	keyValueMap["fffffffff3"] = "ddddddddddddd"
	keyValueMap["fffffffff5"] = "ddddddddddddd"
	keyValueMap["fffffffff6"] = "ddddddddddddd"
	keyValueMap["fffffffff6"] = "ddddddddddddd"
	keyValueMap["fffffffff"] = "ddddddddddddd"
	keyValueMap["fffffffff"] = "ddddddddddddd"
	keyValueMap["fffffffff"] = "ddddddddddddd"
	keyValueMap["fffffffff"] = "ddddddddddddd"

	for key := range keyValueMap {

		fmt.Printf(" the %s    and value is %s  \n", key, keyValueMap[key])

	}

	var strutsmap map[string]person

	strutsmap = make(map[string]person)

	var personitem1 person
	personitem1.name = "source"

	strutsmap["ddd"] = personitem1

	for key := range strutsmap {

		fmt.Printf(" ############the %s    and value is %s  \n", key, strutsmap[key])

	}

}
