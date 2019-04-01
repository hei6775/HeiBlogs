package main

import (
	"os"
	"fmt"
	"time"
	"path/filepath"
	"io/ioutil"
	"strconv"
	"flag"
	"path"
	"strings"
	"bufio"
	"io"
	"encoding/json"
)

type WsJson struct {
	WsAddr string `json:"WSAddr"`
}

var WsAddr1 = flag.Int("ws",30101,"game WSAddr")

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		time.Sleep(5*time.Second)
	}
}
func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)
	return dir
}

func main(){
	var (
		listfile []os.FileInfo
		childfiles map[string]int
		dir string
		err error
	)
	childfiles = make(map[string]int)

	flag.Parse()
	dir = getCurrentPath()
	fmt.Println("当前目录：",dir)

	listfile,err = ioutil.ReadDir(dir)
	checkErr(err)
	time.Sleep(3*time.Second)
	for _,v := range listfile {
		var (
			id int
			isgame bool
			idStr,childfile string
			err1 error
		)
		isgame,_ = path.Match("game*",v.Name())
		if v.IsDir() && isgame{
			idStr = strings.Trim(v.Name(),"game")
			id,err1 = strconv.Atoi(idStr)
			checkErr(err1)
			childfile = path.Join(dir,v.Name(),"gameconf","server.json")
			childfiles[childfile] = id
		}
	}
	fmt.Println("所有game配置目录：",childfiles)

	fmt.Println("============================")

	for k,v := range childfiles {
		Readfile(k,*WsAddr1,v)
	}
	time.Sleep(3*time.Second)
}

func Readfile(filename string,ws,n int){
	var (
		f,newFile *os.File
		r *bufio.Reader
		w *bufio.Writer
		data,result []byte
		name string
		err error
	)
	n = ws+n
	name = path.Base(filename)
	f,err = os.Open(filename)
	if err != nil {
		fmt.Println("读取文件：",name,"错误：",err)
		f.Close()
		return
	}

	r=bufio.NewReader(f)

	for {
		data,err = r.ReadBytes('\n')
		if err == io.EOF {
			result = append(result,data...)
			break
		}
		data = toWs(data,n)
		result = append(result,data...)

	}
	f.Close()
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("删除原文件:",name,"错误：",err)
		return
	}
	newFile, err = os.Create(filename)
	if err != nil {
		fmt.Println("创建文件",name,"错误：",err)
		return
	}
	w = bufio.NewWriter(newFile)
	w.Write(result)
	w.Flush()
}


func toWs(input []byte,n int)[]byte{
	var (
		a []byte
		wsData *WsJson
		err error
	)
	a = append(a,'{')
	a = append(a,input...)
	a = append(a,'}')
	if err = json.Unmarshal(a,wsData);err != nil{
		return input
	}
	return []byte("\"WsAddr\": \":"+ strconv.Itoa(n) +"\",\n")

}
