/**
  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

  socket 通信模拟测试程序

 */package main

import (
	"fmt"
	"net"
	"os"
	//"time"
	"bufio"
	"log"
	"strconv"
	"time"
)

type sendMsg struct {
	msgid int
	msg   string
}

var chanMap map[string]chan sendMsg

func main() {

	fmt.Printf(" 这是服务器端 \n ")

	listener, err := net.Listen("tcp", "192.168.23.36:8888")

	if err != nil {
		log.Println("listen error:", err)
		return
	}

	flag := 1

	for {

		conn, err := listener.Accept()

		var channel chan sendMsg = make(chan sendMsg, 3)
		chanMap[strconv.Itoa(flag)] = channel

		if err != nil {
			continue
		}

		go sendmsg(conn, channel)
		go receivemsg(conn, channel)

		flag++
		/*fmt.Printf(" 收到一个客户端的请求 \n ")

		data, err := bufio.NewReader(conn).ReadString('\n')

		fmt.Printf(" 客户端发送的数据是：  %s  \n ", data)

		if err != nil {
			continue
		}

		conn.Write([]byte( "  ================  ** " + time.Now().Format("2006-01-02 15:04:05") + "  this is server sened \n"))
		*/
	}

}

/**
 * 接收消息
 */
func receivemsg(conn net.Conn, channel chan sendMsg) {

	defer conn.Close()

	for {

		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			continue
		}
		fmt.Printf(" 客户端发送的数据是：  %s  \n ", data)
	}

}

/**
 * 发送消息
 */
func sendmsg(conn net.Conn, channel chan sendMsg) {

	defer conn.Close()

	for {

		conn.Write([]byte("  ================  ** " + time.Now().Format("2006-01-02 15:04:05") + "  this is server sened \n"))
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err)
		os.Exit(1)
	}
}
