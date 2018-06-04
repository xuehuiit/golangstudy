/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

  以太坊客户端模拟器，模拟客户端发起交易

 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/17golang/golang/goenv"
	"github.com/17golang/golang/gonetwork"
	//gologging "github.com/golang/glog"
	"time"
	//"flag"
)

type result_accounts struct {
	Jsonrpc string   `json:"jsonrpc"`
	Id      int32    `json:"id"`
	Result  []string `json:"result"`
}

var mock_blank string = "0xd56ed21460299b0d442678cdfa419cc80194c62f"
var eth_server = "http://172.16.10.163:8543"

var sendtranChal chan string = make(chan string, 10)

func main() {

	//goenv.InitApp()

	//flag.Parse()

	for key, value := range goenv.AppConstant {

		fmt.Println(" the key is %s and  value is %s ", key, value)

	}

	sendtans := `{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}`

	//glogging.Info("ddddddddddddddddd")

	for ind := 1; ind < 10; ind++ {
		go sendTransThread(fmt.Sprintf("线程  %d", ind))
	}

	for true {

		_, contesnts, err := gonetwork.HttpPost4String(eth_server, nil, sendtans)

		if err == nil {

			//println(contesnts)
			result_accountsstruct := &result_accounts{}
			json.Unmarshal([]byte(contesnts), &result_accountsstruct)

			result := result_accountsstruct.Result

			for key := range result {

				sendaccount := result[key]

				if sendaccount == mock_blank {
					continue

				}

				tpl := `{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from":"%s","to":"%s","gasPrice":"0","value":"2","data":"DDDDDDDDDDD"}],"id":"67"}`
				vdd := fmt.Sprintf(tpl, mock_blank, sendaccount)

				sendtranChal <- vdd

				//glogging.Info(vdd)

			}

		}

		time.Sleep(2 * 1000 * 1000 * 1000)
	}
	//sendtans := `{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from":"0xd56ed21460299b0d442678cdfa419cc80194c62f","to":"0xa0c6dd260cafac2ad2919a27009432f34e1aa1d5","gasPrice":"0","value":"2","data":"DDDDDDDDDDD"}],"id":"67"}`

}

/**

 */
func sendTransThread(threadinfo string) {

	fmt.Println(threadinfo)

	for true {

		sendJson := <-sendtranChal

		_, contesnts, err := gonetwork.HttpPost4String(eth_server, nil, sendJson)

		if err == nil {

			fmt.Println("------" + time.Now().Format("2006-01-02 15:04:05") + " -----------------------------------  " + contesnts)

		}
	}
}
