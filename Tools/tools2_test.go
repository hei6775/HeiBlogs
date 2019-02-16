package Tools

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"regexp"
	"testing"
)

type qudao struct {
	Id    string
	money float64
}

var qudaos = map[string]*qudao{
	"紫云（5432u）": {"10733", 0},
	"1758":      {"303", 0},
	"4177":      {"10215", 0},
	"7723":      {"10726", 0},
	"KUKU":      {"10802", 0},
	"X游网":       {"406", 0},
	"8090":      {"10454", 0},
	"妖":         {"10884", 0},
	"萝卜玩":       {"10795", 0},
	"啪啪游":       {"10604", 0},
	"龙川奇点	": {"10620", 0},
	"一牛": {"10396", 0},
	"欢聚游HJY（盛世游戏）": {"10386", 0},
	"游民星空":         {"10724", 0},
	"333游戏":        {"10214", 0},
	"4399":         {"10531", 0},
	"乐嗨嗨(aes勾选去除)": {"10465", 0},
	"游戏fan（新）":     {"10916", 0},
	"511wan":       {"10255", 0},
	"杭州掌盟":         {"315", 0},
	"易乐玩":          {"yilewan", 0},
}

var infos = []byte(`{"datas":[`)

var Cols = []string{"A", "B", "C"}
var Rows = 2

//黏贴复制，手动删掉/* 1 */即可
var recharge = []byte(`
`)

//like this
//ss := []byte(`{
//    	"_id" : "10620",
//    	"total" : 6.0
//		}
//
//		/* 2 */
//		{
//    	"_id" : "10795",
//    	"total" : 640.0
//		}
//	`)

func TestMmarshal(t *testing.T) {

	pat := "/[\\*] [0-9]{1,2} [\\*]/"
	recharge_str := string(recharge)
	re, _ := regexp.Compile(pat)
	re_result := re.ReplaceAllString(recharge_str, ",")
	fmt.Print(re_result)
	re_result = fmt.Sprintf("%v%v%v", string(infos), re_result, "]}")

	result := Munmarshal([]byte(re_result))
	if result == nil {
		fmt.Println("result is nil")
		return
	}

	for _, one := range result.Datas {
		for _, oneQuDao := range qudaos {
			if oneQuDao.Id == one.Id {
				fmt.Println(oneQuDao.Id)
				oneQuDao.money = one.Total
			}
		}
	}

	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//dir = filepath.Join(dir, "recharge.xlsx")
	//fmt.Println()
	xlsx := excelize.NewFile()

	_ = xlsx.NewSheet("Sheet1")
	title := []string{"渠道名称", "渠道ID", "充值金额"}
	for k, v := range title {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[k], 1), v)
	}

	for k, v := range qudaos {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[0], Rows), k)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[1], Rows), v.Id)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[2], Rows), v.money)
		Rows += 1
	}
	err := xlsx.SaveAs("./recharge.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func TestMmarshal2(t *testing.T) {
	fmt.Println(string(Mmarshal()))
}
