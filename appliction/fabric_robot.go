/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

  fabric 0.6 客户端机器人，模拟客户端的相关操作，增加测试数据

 */
package main



import (
	//"encoding/json"
	"fmt"
	//"github.com/17golang/golang/goenv"
	"github.com/17golang/golang/gonetwork"
	//gologging "github.com/golang/glog"
	"time"
	//"flag"
	"math/rand"
)


var chaincodeid = "ffb96aba7f7d2a90dcbd55ceb63b8df280b31596e5728b06c85de19aa1d4c83c7a1b30c74263c2ab6b7c39b7165eff09e773fbdd1bd525b8aa39482783d425d8"
var chaincode_url = "http://121.43.167.246:8050/chaincode"

var chain_client_json_tel = `{
  "jsonrpc": "2.0",
  "method": "invoke",
  "params": {
    "type": 1,
    "chaincodeID": {
      "name": "%s"
    },
    "ctorMsg": {
      "function": "write",
      "args": [
        "testkey","psend###value####"
      ]
    },
    "secureContext": "jim"
  },
  "id": 0
}`



var sendtranChal1 chan string = make(chan string, 10)


func main()  {


	for ind := 1; ind < 10; ind++ {
		go sendFabricTranThread(fmt.Sprintf("发送线程  %d", ind))
	}


	for true {

		counts := rand.Intn(10);

		fmt.Printf("本次发送 %d 条记录  ",counts)
		for intt:=0 ; intt<counts; intt++  {
			fmt.Println("send to queue")
			sendtranChal1 <- "ddddd"
		}
		time.Sleep(2 * 1000 * 1000 * 1000)

	}



}


/**
调用fabric chaincode 发送交易的过程
 */
func sendFabricTranThread(threadinfo string) {

	fmt.Println(threadinfo)

	for true {

		sendJson := <-sendtranChal1

		chain_cmd:= fmt.Sprintf(chain_client_json_tel,chaincodeid)

		//fmt.Println(chain_cmd)


		_, contesnts, err := gonetwork.HttpPost4String(chaincode_url, nil, chain_cmd)

		if err == nil {

			fmt.Println("------" + time.Now().Format("2006-01-02 15:04:05") + " -----------------------------------  " + contesnts+sendJson)

		}
	}
}
