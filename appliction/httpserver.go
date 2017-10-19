package main


import(

  	"fmt"
  	"log"
  	"net/http"
	"time"
	"os/exec"
	"io"
)


//启动服务器
func startuphttpd(rw http.ResponseWriter, req *http.Request){


	req.ParseForm() //解析参数，默认是不会解析的


	//fmt.Println( time.Now().Format("2006-01-02 15:04:05")+"   id的值:  " + "  " +req.PostFormValue("id"));


	query := req.URL.Query()
	get_act := query["act"][0]
	//fmt.Println(get_act)

	if( get_act == "ylfremovteshoidnd" ){

		fmt.Println( time.Now().Format("2006-01-02 15:04:05") + "   cat的值:  " + "  " +get_act);

		cmd := exec.Command("/bin/sh", "-c", "service httpd  restart")
		//cmd := exec.Command("/etc/init.d/httpd ", " restart ")

		//exec.
		out, err := cmd.CombinedOutput()


		if err != nil {
			fmt.Println(" error !  ")
			io.WriteString(rw, " run error ")

		}else{
			fmt.Println(string(out))
			io.WriteString(rw, " run  success  "+string(out))
		}


		}else{
			fmt.Println( time.Now().Format("2006-01-02 15:04:05") + " 非法的值 " )
		}



}


//更新信息
func updateRoom(rw http.ResponseWriter, req *http.Request) {
       req.ParseForm();
}


func main() {

  http.HandleFunc("/restarthttp", startuphttpd)

  s := &http.Server{
    Addr: ":11002",
    ReadTimeout: 60 * time.Second,
    WriteTimeout: 60 * time.Second,
  }

  //.HandleFunc("/updateRoom", updateRoom)
  log.Fatal(s.ListenAndServe())

}