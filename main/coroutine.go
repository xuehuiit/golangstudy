/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

 */

package main

import (
	"fmt"
	"math/rand"
)



func  rand1() chan int {

	out := make( chan int , 10)

	go func() {

		for {

			out <- rand.Int()
		}
	}()

	return out

}


func  rand_m() chan int{

	out := make( chan int)

	rand1_1 := rand1()
	rand1_2 := rand1()

	go func() {

		for{

			out <- (<- rand1_1)/1000

		}

	}()

	go func() {

		for{

			out <- (<- rand1_2)/10000000

		}
	}()


	return out
}

func main(){


	rand_service_hand := rand_m()

	for{

		select{
		case radnvalue := <-rand_service_hand:
			fmt.Println(" this is %d ",radnvalue)
		}

	}




	/*

	for ind:=1 ; ind < 20 ; ind++ {
		fmt.Println(" this is %d ", <-rand_service_hand)
	}*/

	/*for{

		fmt.Println(" this is %d ", 1)

	}*/


}