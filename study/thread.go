package study

import (
	"fmt"
	"time"
)

func TestThread() {

	//thread()
	threadchan()

}

var channel chan int = make(chan int, 3)

func threadchan() {

	go threadcall("线程1")
	go threadcall("线程2")

	for i := 1; i < 10; i++ {

		channel <- i

		//time.Sleep(1000*1000*1000*2)
	}

	time.Sleep(1000 * 1000 * 1000 * 2)

}

func threadcall(threadname string) {

	for true {
		queuev := <-channel
		fmt.Printf(" 线程  %s 从队列中获取的值   %d  \n", threadname, queuev)
	}

}

//============

/**
 * 简单的线程
 */
func thread() {

	go loop()
	loop()

	time.Sleep(time.Second)
}

func loop() {

	for i := 0; i < 10; i++ {
		fmt.Printf("loop %d  \n", i)
	}
}
