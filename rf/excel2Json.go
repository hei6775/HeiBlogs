package rf

import (
	"github.com/Luxurioust/excelize"
	"fmt"
	"os"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
	"path/filepath"
)

var ActiveSheet = 1
var Directory = "jsonConfigs"
var RegexpPattern = "^#.*"
var JsonDir = "Json"

const (
	IntType = iota
	FloatType
	Other
)

func readExcel(filaname string){
	xlsx, err := excelize.OpenFile(filaname)
	if err != nil {
		fmt.Printf("The filename [%v] [%v] \n",filaname,err)
		return
	}

	xlsx.SetActiveSheet(ActiveSheet)
	sheetName := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())
	// Get all the rows in the Sheet1.
	//if len(xlsx.Sheet) != 1 {
	//	fmt.Println("Your Config sheet conut over 1,must 1")
	//	return
	//}
	fieldName := []string{}
	//get the rows of sheetName xlsx
	rows := xlsx.GetRows(sheetName)

	josnDir,jsonF := path.Split(filaname)
	jsonF = strings.TrimSuffix(jsonF,".xlsx")+".json"
	jsonF = path.Join(josnDir,JsonDir,jsonF)
	jsonFile,err := os.Create(jsonF)
	if err != nil {
		fmt.Printf("create json file [%v] \n",err)
		return
	}
	defer jsonFile.Close()
	jsonFile.WriteString("[\n")
	for rowindex, row := range rows {
		switch rowindex {
		case 0:
			continue
		case 1:
			continue
		case 2:

		default:
			jsonFile.WriteString("\t{\n")
		}

		for k, colCell := range row {
			switch rowindex {
			case 1:
				if k == 0 {
					isReg,errReg := regexp.MatchString(RegexpPattern,colCell)
					if errReg != nil {
						fmt.Printf("match string pattern[%v] string[%v]",RegexpPattern,colCell)
						return
					}
					if !isReg {
						fmt.Printf("the first row must be comment \n")
						return
					}
				}
			case 2:
				if colCell == "" ||colCell == " "{
					continue
				}
				fieldName = append(fieldName,colCell)
			default:
				if k>= len(fieldName){
					continue
				}
				strType,strResult := ReflectValue(colCell)
				switch strType {
				case 0:
					strResult = fmt.Sprintf("\t\"%s\":%s",fieldName[k],colCell)
				case 1:
					strResult = fmt.Sprintf("\t\"%s\":%s",fieldName[k],colCell)
				case 2:
					strResult = fmt.Sprintf("\t\"%s\":\"%s\"",fieldName[k],colCell)
				default:
					strResult = fmt.Sprintf("\t\"%s\":\"%s\"",fieldName[k],colCell)
				}
				if k+1 == len(fieldName) {
				}else{
					strResult += ",\n"
				}

				jsonFile.WriteString(strResult)
			}

		}
		switch rowindex {
		case 0:
		case 1:
		case 2:
		case len(rows)-1:
			jsonFile.WriteString("\n\t}\n")
		default:
			jsonFile.WriteString("\n\t},\n")
		}
	}
	jsonFile.WriteString("]")

}

//convert excel to json
func excel2json(name string)error{
	fileInfos,err :=ioutil.ReadDir(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _,v := range fileInfos {
		if v.IsDir() {
			continue
		}
		fileName := v.Name()
		ext := path.Ext(fileName)
		if ext != ".xlsx"{
			continue
		}
		readExcel(fileName)
	}
	return nil
}

func createDir(name string){
	flag := RemoveAll(path.Join(name,JsonDir))
	if !flag {
		fmt.Println("Remove the configs has some problems")
		return
	}
	fmt.Println("创建Json文件夹")
	err := os.Mkdir(JsonDir,os.ModeDir)
	if err != nil {
		fmt.Println("创建json文件夹失败",err)
		return
	}
	excel2json(name)
}


//check the directory and remove the all files
func RemoveAll(name string)(bool){
	_,err := os.Stat(name)
	if err == nil {
		errRe := os.RemoveAll(name)
		if errRe != nil {
			fmt.Println(errRe)
			return false
		}
		return false
	}
	if os.IsExist(err) {
		errRe := os.RemoveAll(name)
		if errRe != nil {
			fmt.Println(errRe)
			return false
		}
		return true
	}
	return true
}

//check the value type
func ReflectValue(str string)(int,string){
	result := strings.Split(str,".")
	if len(result) != 2 && len(result) != 1 {
		return 2,str
	}

	if len(result) == 1 {
		ok2,_ := regexp.MatchString("^([0-9])+",result[0])
		if!ok2{
			return 2,str
		}
		return 0,str
	}

	if len(result) == 2 {
		if ok2,_ := regexp.MatchString("^[0-9]+",result[0]);!ok2{
			return 2,str
		}

		if ok,_:=regexp.MatchString("^[0-9]+",result[1]);!ok {
			return 2,str
		}
	}

	return 1,str

}

func main(){
	nowDir :=getCurrentPath()
	fmt.Println("当前地址：",nowDir)
	createDir(nowDir)
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return dir
}