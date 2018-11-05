package main

import (
	"net/http"
	"log"
	"time"
	"fmt"
	"io/ioutil"
	"io"
	"bufio"
)
//http服务端
//1、mux.Handle
//2、http.HandlerFunc,mux.Handle
//3、mux.HandleFunc
//4、闭包
type timeHandler struct {
	format string
}

func (th *timeHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: "+tm))
}

func timeHH(w http.ResponseWriter,r *http.Request){
	newtime := time.Now().Format(time.RFC3339)
	w.Write([]byte(newtime))
}

func coo(w http.ResponseWriter,r *http.Request){
	fmt.Println("r.Host: ",r.Host)
	fmt.Println("r.Body: ",r.Body)
	fmt.Println("r.Header",r.Header)
	p,_ := ioutil.ReadAll(r.Body)
	w.Write(p)
}

func eoo(w http.ResponseWriter,r *http.Request){
	p,_ := ioutil.ReadAll(r.Body)
	fmt.Println("receive body:",string(p))
	w.Write([]byte("success"))

	io.WriteString(w,"success \n")
	f :=bufio.NewWriter(w)
	f.WriteString("success")
	f.Flush()
}

func main(){
	mux := http.NewServeMux()
	//1
	rh := http.RedirectHandler("http://www.baidu.com",307)
	mux.Handle("/foo",rh)
	//2
	a := &timeHandler{format:time.RFC1123}
	mux.Handle("/aoo",a)

	//3
	b := http.HandlerFunc(timeHH)
	mux.Handle("/boo",b)

	//4
	mux.HandleFunc("/coo",coo)

	mux.HandleFunc("/eoo",eoo)

	log.Println("Listening...")
	http.ListenAndServe(":9630",mux)
}
