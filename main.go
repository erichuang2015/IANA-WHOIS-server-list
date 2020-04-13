package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var f *os.File

const (
	latestf = "latest.log" //记录程序运行以来最后抓取的TLD
)

func init() {
	tempf, err := os.OpenFile("serverlist.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	checkErr(err)
	f = tempf
}

func main() {

	http.HandleFunc("/whois/latestTLD", func(w http.ResponseWriter, r *http.Request) {
		setAccessControlAllowOrigin(w)
		if http.MethodGet != r.Method {
			w.WriteHeader(http.StatusBadRequest)
		}
		buff, err := ioutil.ReadFile(latestf)
		checkErr(err)
		io.WriteString(w, string(buff))
	})

	http.HandleFunc("/whois", func(w http.ResponseWriter, r *http.Request) {
		setAccessControlAllowOrigin(w)

		//不是POST不接受
		if http.MethodPost != r.Method {
			w.WriteHeader(http.StatusBadRequest)
		}

		body, err := ioutil.ReadAll(r.Body)
		checkErr(err) //没提交内容不接受

		o := struct {
			TLD    string `json:"TLD"`
			Server string `json:"server"`
		}{}

		fmt.Println(string(body))

		err = json.Unmarshal(body, &o)
		checkErr(err) //提交的内容转换失败不接受

		//抓取的内容持久化到文件中
		f.WriteString(fmt.Sprintf("%v:%v\n", o.TLD, o.Server))

		//记录上一次的TLD，如果重启程序接着来
		ioutil.WriteFile(latestf, []byte(o.TLD), os.ModePerm)

		//OK
		w.WriteHeader(http.StatusNoContent)
	})
	http.ListenAndServe(":8081", nil)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func setAccessControlAllowOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "https://www.iana.org")
}
