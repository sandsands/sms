package main

import (
	"fmt"
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
	fmt.Printf("%s", result)
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