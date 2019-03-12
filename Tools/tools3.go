package Tools

import (
	"os"
	"fmt"
	"time"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"flag"
	"path"
	"strings"
)
var WsAddr = flag.Int("ws",30101,"game WSAddr")

type GameConf struct {
	LogLevel string `json:"LogLevel"`
	LogPath string `json:"LogPath"`
	WSAddr string `json:"WsAddr"`
	DBUrl string `json:"DbUrl"`
	AccDB string `json:"AccDB"`
	GameDB string `json:"GameDB"`
	ServerID int `json:"ServerID"`
	RecordDBUrl string `json:"RecordDBUrl"`
	RecordDB string `json:"RecordDB"`
	RechargeAddr string `json:"RechargeAddr"`
	LoginAddr string `json:"LoginAddr"`
}

func confEditFile(data []byte,port int)(result []byte){
	conf := new(GameConf)
	err := json.Unmarshal(data,conf)
	checkErr(err)
	portStr := strconv.Itoa(port)
	conf.WSAddr = fmt.Sprintf(":%v",portStr)
	result,err = json.MarshalIndent(conf,"","\t")
	checkErr(err)
	return result
}

func readConfFile(confpath string,port int,add int)(){
	data,err :=ioutil.ReadFile(confpath)
	checkErr(err)
	port = port +add-1
	result := confEditFile(data,port)
	ioutil.WriteFile(confpath,result,0644)
}


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
	flag.Parse()
	dir := getCurrentPath()
	fmt.Println("当前目录：",dir)
	var childfiles map[string]int
	childfiles = make(map[string]int)
	listfile,err := ioutil.ReadDir(dir)
	checkErr(err)
	time.Sleep(3*time.Second)
	for _,v := range listfile {
		isgame,_ := path.Match("game*",v.Name())
		if v.IsDir() && isgame{
			idStr := strings.Trim(v.Name(),"game")
			id,err := strconv.Atoi(idStr)
			checkErr(err)
			childfile := path.Join(dir,v.Name(),"gameconf","server.json")
			childfiles[childfile] = id
		}
	}
	fmt.Println("所有game配置目录：",childfiles)

	fmt.Println("============================")

	for k,v := range childfiles {
		readConfFile(k,*WsAddr,v)
	}
	time.Sleep(3*time.Second)
}