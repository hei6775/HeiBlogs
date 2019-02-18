package Tools

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"regexp"
	"testing"
)

//渠道名称
var QuDaosName = []string{
	"紫云（5432u）",
	"1758",
	"4177",
	"7723",
	"KUKU",
	"X游网",
	"8090",
	"妖",
	"萝卜玩",
	"啪啪游",
	"龙川奇点	",
	"一牛",
	"欢聚游HJY（盛世游戏）",
	"游民星空",
	"333游戏",
	"4399",
	"乐嗨嗨(aes勾选去除)",
	"游戏fan（新）",
	"511wan",
	"杭州掌盟",
	"易乐玩",
}

//渠道ID
var QuDaosId = []string{
	"10733",
	"303",
	"10215",
	"10726",
	"10802",
	"406",
	"10454",
	"10884",
	"10795",
	"10604",
	"10620",
	"10396",
	"10386",
	"10724",
	"10214",
	"10531",
	"10465",
	"10916",
	"10255",
	"315",
	"yilewan",
}

//写入渠道的数据结构
type qudao struct {
	Id       string  //Id
	Money    float64 //充值金额
	OrderNum float64 //订单数量
}

//渠道数据map
var qudaos = map[string]*qudao{}

//数据头 用来拼接数据
var infos = []byte(`{"datas":[`)

//行号
var Cols = []string{"A", "B", "C", "D"}
var Rows = 2

//渠道总充值，根据渠道区分 附带订单数
//db.getCollection('recharge').aggregate([{$match:{"state":{$gte:1}}},{$group:{_id:"$realchan",ordernum:{$sum:1},total:{$sum:"$money"}}}])
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

	//初始化map
	for i, _ := range QuDaosName {
		oneQuDao := new(qudao)
		oneQuDao.Id = QuDaosId[i]
		qudaos[QuDaosName[i]] = oneQuDao
	}
	// 正则将/* 2 */之类的替换成,
	pat := "/[\\*] [0-9]{1,2} [\\*]/"
	recharge_str := string(recharge)
	re, _ := regexp.Compile(pat)
	re_result := re.ReplaceAllString(recharge_str, ",")
	//fmt.Print(re_result)
	re_result = fmt.Sprintf("%v%v%v", string(infos), re_result, "]}")
	//json反序列化
	result := Munmarshal([]byte(re_result))
	if result == nil {
		fmt.Println("result is nil")
		return
	}
	//赋值
	for _, one := range result.Datas {
		for _, oneQuDao := range qudaos {
			if oneQuDao.Id == one.Id {
				//fmt.Println(oneQuDao.Id)
				oneQuDao.Money = one.Total
				oneQuDao.OrderNum = one.Ordernum
			}
		}
	}

	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//dir = filepath.Join(dir, "recharge.xlsx")
	//fmt.Println()
	//生成excel
	xlsx := excelize.NewFile()

	_ = xlsx.NewSheet("Sheet1")
	title := []string{"渠道名称", "渠道ID", "充值金额", "渠道成功订单数"}
	for k, v := range title {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[k], 1), v)
	}

	for k, v := range qudaos {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[0], Rows), k)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[1], Rows), v.Id)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[2], Rows), v.Money)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("%v%v", Cols[3], Rows), v.OrderNum)
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
