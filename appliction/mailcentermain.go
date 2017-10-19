/**
  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

  一个模拟邮件发送和http请求的程序

 */
package appliction

//github.com/alecthomas/log4go
import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	//"log"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	"time"
)

var printflag int = 1

type mailsendresult struct {
	Result               string `json:"result"`
	ResultFlag           string `json:"resultFlag"`
	ResultExceptionfInfo string `json:"resultExceptionfInfo"`
}

//{"cmd":"","resultflag":"","mail":"mailinfo. dddddddddd","remoteproxy":"http:\/\/share.eshowpro.com\/\/share\/remotecurls","testflag":"pquX0QCdUaI="}
type mail_send_result_cmd struct {
	Cmd           string        `json:"cmd"`
	Resultflag    string        `json:"resultflag"`
	Sendorginfo   string        `json:"sendorginfo"`
	Hgpurchaserid string        `json:"hgpurchaserid"`
	Sendcontent   string        `json:"sendcontent"`
	Testflag      string        `json:"testflag"`
	Sendmails     []sendmailinf `json:"sendmails"`
}

type sendmailinf struct {
	Smtpaccount  string `json:"smtpaccount"`
	Smtppasswd   string `json:"smtppasswd"`
	Smtpurl      string `json:"smtpurl"`
	Smtpport     string `json:"smtpport"`
	Sendmailitem string `json:"sendmailitem"`
	Sendmailid   int32  `json:"sendmailid"`
}

func update(rw http.ResponseWriter, req *http.Request) {

	fmt.Println(" 收到了一个http 请求    ")

}

func main() {

	/*str := "welcome to sharejs.com"
	str = strings.Replace(str, " ", ",", -1)
	fmt.Println(str)

	*/

	//GethTest()

	if printflag == 0 {
		fmt.Println("I am herer")
	}

	//http.HandleFunc( "/build" , update )

	//log.Fatal(http.ListenAndServe(":8002", nil))

	//httpcode,contesnts := httpGet("http://api.eshowpro.com/table/remotenotify", nil)

	//prms := &postparm {
	//"asdf",
	//"3",
	//"3",
	//"ddddd",
	//}

	prms := make(map[string]string)

	prms["checkflag"] = "dddd"
	prms["cmd"] = "3"
	prms["splids"] = "3"
	prms["extdata"] = "dddddddddd"
	prms["version"] = "1.0"
	//var prms postparm
	//prms.checkflag = "dddd"
	// prms.cmd = "3"
	// prms.splids = "d"
	//prms.extdata = "dddddddddd";

	res, _ := json.Marshal(prms)
	//fmt.Println(res)
	if printflag == 0 {
		fmt.Println(string(res))
	}
	parmbase64 := base64.StdEncoding.EncodeToString(res)

	//fmt.Println(string(parmbase64))
	parmmap := map[string]string{"contents": parmbase64}

	//发送邮件
	//sendmailtest()

	for true {

		_, contesnts, err := HttpPost("http://api4.eshowpro.com/table/remotenotify", nil, parmmap)
		//httpcode,contesnts := HttpPost4String("http://api.eshowpro.com/table/remotenotify", nil,"contents="+parmbase64)
		//fmt.Printf(" 返回编码 %d  返回值  %s  \n" , httpcode,contesnts )
		//HttpPost1();
		//httpDo()
		if err == nil {

			mailsendresultstruct := &mailsendresult{}
			json.Unmarshal([]byte(contesnts), &mailsendresultstruct)

			result := mailsendresultstruct.Result
			//fmt.Println(result)
			base64code := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

			result_un64, _ := base64code.DecodeString(result)

			//fmt.Println(string(result_un64)+"\n")

			mailsendresultcmdstr := &mail_send_result_cmd{}
			json.Unmarshal([]byte(result_un64), &mailsendresultcmdstr)

			resultflag := mailsendresultcmdstr.Resultflag

			if resultflag == "1" {

				sendmailcontent := mailsendresultcmdstr.Sendcontent
				sendmailcontent_un64, _ := base64code.DecodeString(sendmailcontent)

				sendmails := mailsendresultcmdstr.Sendmails

				for sendmailtiemindex := range sendmails {

					smtpaccount := sendmails[sendmailtiemindex].Smtpaccount
					smtppasswd := sendmails[sendmailtiemindex].Smtppasswd
					smtpport := sendmails[sendmailtiemindex].Smtpport
					sendmailid := sendmails[sendmailtiemindex].Sendmailid
					smtpurl := sendmails[sendmailtiemindex].Smtpurl
					sendmailitem := sendmails[sendmailtiemindex].Sendmailitem

					if printflag == 0 {
						fmt.Printf(" sendmailid :   %d  \n", sendmailid)
					}

					sendmailidstr := fmt.Sprintf("%d", sendmailid)

					if printflag == 0 {
						fmt.Printf(" sendmailidstr :   %s \n", sendmailidstr)
					}

					sendmailcontents := strings.Replace(string(sendmailcontent_un64), "###sendid###", sendmailidstr, -1)

					//sendmailitem = "fengxiang@eshowpro.com"

					//fmt.Printf( " %s %s %s %s %s %s  " ,smtpaccount,smtppasswd,smtpport  )

					sendmail(smtpaccount, smtppasswd, smtpurl+":"+smtpport, sendmailitem, "hello", sendmailcontents)

				}

			}

			if printflag == 0 {
				fmt.Println("================== " + time.Now().Format("2006-01-02 15:04:05") + " =============================")
			}

			time.Sleep(10 * 1000 * 1000 * 1000)

		} else {
			//fmt.Println(err.Error())
		}

	}

}

func sendmail(user string, password string, host string, to string, subject string, body string) {

	if printflag == 0 {
		fmt.Printf("%s-%s-%s-%s-%s    \n", user, password, host, to, subject)
	}
	/*user = "ruogou763864@sohu.com"
	password = "poliangkao"
	host = "smtp.sohu.com:25"
	to = "fx19800215@163.com"*/

	//subject = "使用Golang发送邮件"
	/*
		body := `
			<html>
			<body>
			<h3>
			"Test send to email"
			</h3>
			</body>
			</html>
			`*/
	if printflag == 0 {
		fmt.Println("send email")
	}

	err := SendToMail(user, password, host, to, subject, body, "html")

	if err != nil {
		if printflag == 0 {
			fmt.Println("Send mail error!")
			fmt.Println(err)
		}
	} else {
		if printflag == 0 {
			fmt.Println("Send mail success!")
		}
	}
}

func sendmailtest() {

	user := "fx19800215@163.com"
	password := "FengXiang1980"
	host := "smtp.163.com:25"
	to := "fengxiang@eshowpro.com"

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

/**
 * http get 方法
 */
func httpGet(url string, httphead map[string]string) (int, string) {

	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)

	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	//reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	if httphead != nil {

		for key := range httphead {

			value := httphead[key]

			reqest.Header.Set(key, value)

		}

	}

	var bodystr string
	response, _ := client.Do(reqest)

	if response.StatusCode == 200 {

		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 16)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				bodystr += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(response.Body)
			bodystr = string(bodyByte)
		}
	}

	defer response.Body.Close()
	return response.StatusCode, bodystr

}

/**
 * http post 基本方法
 */
func httpPostBase(paths string, httphead map[string]string, bodycontent io.Reader) (int, string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", paths, bodycontent)
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//req.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//req.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	//req.Header.Set("Accept-Language","zh-CN,zh;q=0.8")

	if httphead != nil {

		for key := range httphead {

			value := httphead[key]

			req.Header.Set(key, value)

		}

	}

	resp, err := client.Do(req)

	if err == nil {

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		return resp.StatusCode, string(body), err
	} else {
		return -100, "", err
	}

	//fmt.Println(string(body))

	////////////////////////////////

}

/**
 *
 */
func HttpPost4String(paths string, httphead map[string]string, body string) (int, string, error) {

	return httpPostBase(paths, httphead, strings.NewReader(body))

}

/**
 * 传入参数的HTTPpost
 */
func HttpPost(paths string, httphead map[string]string, parmmap map[string]string) (int, string, error) {

	postValues := url.Values{}

	if parmmap != nil {

		for key := range parmmap {

			value := parmmap[key]
			postValues.Set(key, value)

		}
	}

	//postbody := ioutil.NopCloser(strings.NewReader(postValues.Encode()))

	postDataStr := postValues.Encode()
	postDataBytes := []byte(postDataStr)

	postBytesReader := bytes.NewReader(postDataBytes)

	return httpPostBase(paths, httphead, postBytesReader)

}

func HttpPost1() {

	resp, err := http.Post("http://api.eshowpro.com/table/remotenotify",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://api.eshowpro.com/table/remotenotify", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func SendToMail(user, password, host, to, subject, body, mailtype string) error {

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err

}
