package study

import (
	"fmt"
	"log"
	"os"
	//"github.com/op/go-logging"
	"flag"
	"github.com/golang/glog"
	//"github.com/op/go-logging"
)

/**
测试系统日志
*/
func SysLogTest() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[info] ")
	log.Println("======================proscesss======================")

}

func SysLogTest4File() {

	logfile, err := os.OpenFile("source.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("hello")

}

/*
var logging = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
*/

/**
测试   "github.com/op/go-logging" 的日志
*/
/*func GoLogger() {

	logFile, err := os.OpenFile("log.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)

	logging.INFO("INFO");
	logging.NOTICE("NOTIC");

	logging.WARNING("warning")
	logging.ERROR("err")
	logging.CRITICAL("crit")


}*/

func Golog() {

	flag.Parse()
	glog.Info("========hello, glog============")
	glog.Info("hello, glog")
	glog.Info("hello, glog")
	glog.Info("hello, glog")
	glog.Info("hello, glog")
	glog.Info("hello, glog")
	//glog.Error("ddddddddddddd error")
	//glog.Flush()

}
