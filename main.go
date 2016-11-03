package main

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"log"
	"crypto/md5"
	"encoding/hex"
	"time"
	"strconv"
	"strings"
	"os"
)


func sendsms( mobile string, content string) {
	timestamp := time.Now().Unix()
	ntime := strconv.FormatInt(timestamp, 10)
	opass := ""
	s := []string{"key", "=", opass, "&", "timestamp", "=", ntime}
	newpass := strings.Join(s, "")
	npass := md5.New()
	npass.Write([]byte(newpass))
	fpass := hex.EncodeToString(npass.Sum(nil))
	u, _ := url.Parse("")
	q := u.Query()
	q.Set("user_name", "")
	q.Set("password", fpass)
	q.Set("timestamp", ntime)
	q.Set("mobile", mobile)
	q.Set("content", content)
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String());
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", result)
        //fmt.Printf(string(result))

	fileName := "/tmp/a.log"
	logFile,err  := os.Create(fileName)
        defer logFile.Close()
        if err != nil {
         log.Fatalln("open file error !")
        }
         //创建一个日志对象
        debugLog := log.New(logFile,"[Debug]",log.LstdFlags)
        //配置一个日志格式的前缀
        debugLog.SetPrefix("[Info]")
        //配置log的Flag参数
        debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
        debugLog.Println(string(result))

}

func main(){
	args :=os.Args
	if args ==nil ||len(args) < 2 {
		return
	}
	mobile := args[1]
	content := args[2]

	sendsms(mobile, content)
}

//func main()  {
//	sendsms("11", "测试")
//}