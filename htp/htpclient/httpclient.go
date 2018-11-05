package main

import (
	"net/http"
	"log"
	"strings"
	"fmt"
	"io/ioutil"
)
//http客户端
//发起请求
func main(){
	resp,err := http.Post("http://127.0.0.1:9630/eoo",		"application/x-www-form-urlencoded",
		strings.NewReader("mobile=xxxxxxxxxx&isRemberPwd=1"))
	if err != nil {
		log.Println(err)
		return
	}
	p,_:=ioutil.ReadAll(resp.Body)
	fmt.Println("response:",string(p))
	defer resp.Body.Close()
}
