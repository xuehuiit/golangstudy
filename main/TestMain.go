package main

import (
	//"fmt"
	//"goutils"
	"study"
	/*"encoding/hex"
	"time"*/
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func main() {

	study.GetCurrMem()
	/*study.SysLogTest()
	study.SysLogTest4File()*/
	study.Getos()
	//study.GoLogger()
	study.Golog()

	fmt.Println("当前路径" + getCurrentPath())

	//输入启动参数
	arg_num := len(os.Args)
	fmt.Printf("当前启动参数的数目是 ：   %d\n", arg_num)

	fmt.Printf(" 参数是 :\n")
	for i := 0; i < arg_num; i++ {
		fmt.Println(os.Args[i])
	}

	//环境变量
	/*environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
	fmt.Println("------------------------------------------------------------\n")
	logname := os.Getenv("LOGNAME")
	fmt.Printf("logname is %s\n", logname)
	*/
	study.Const_string_change = " 我已经被改变啦！"

	//study.Vardefin()

	//study.TestFlow()

	//study.FunctionTest()

	//study.PointTest()

	//study.CollectionTest()

	//study.TestDate()

	//study.TestString()

	//study.TestThread()

	//IMamain()

	/*sha1 := study.CareatePrivateKey()

	for ind := 1; ind < 20; ind++ {

		sha2 := study.CareatePrivateKey()

		fmt.Println("hex ", hex.EncodeToString(sha1[:]))//16进
		fmt.Println("hex ", hex.EncodeToString(sha2[:]))//16进制

		n := study.Logdist(sha2,sha1 ) //距离
		fmt.Println(n)

		time.Sleep(10 * 1000 * 1000 * 1000)

		fmt.Println("===================================================================")

	}
	*/

	//fmt.Println(sha1)
	//fmt.Println(sha2)

	//mt.Printf("在另外一个包中调用全局变量 调用全局变量  %s "  ,  study.Const_string )

}
