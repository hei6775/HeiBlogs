package mdb

import (
	"testing"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type Address struct {
	Address string
}
type Location struct {
	Longitude float64
	Latitude  float64
}

type Person struct {
	Name     string
	Age_Int  int
	Address  []Address
	Location Location
}


func TestMongoDB(t *testing.T) {
	c, err := Dial("mongodb://admin:123456@127.0.0.1:27017/admin", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// session
	s := c.Ref()
	defer c.UnRef(s)
	err = s.DB("test").C("counters").RemoveId("test")

	if err != nil && err != mgo.ErrNotFound {
		fmt.Println(err)
		return
	}
	// auto increment
	//err = c.EnsureCounter("test", "counters", "test")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for i := 0; i < 3; i++ {
	//	id, err := c.NextSeq("test", "counters", "test")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(id)
	//}
	person := Person{
		Name:    "逍遥",
		Age_Int: 25,
		Address: []Address{
			Address{
				Address: "仙岛",
			},
		},
		Location: Location{
			Longitude: 1,
			Latitude:  1,
		},
	}
	err1 := c.SendData("test", "counters", "1", person)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	query := bson.M{"age_int": 24}
	result, errfind := c.Find("test", "counters", query)
	if errfind != nil {
		fmt.Println("1111", err)
		return
	}
	fmt.Println(result)
	// index
	//c.EnsureUniqueIndex("test", "counters", []string{"key1"})
	//indexs, _ := c.GetIndexs("test", "counters")
	//for _, v := range indexs {
	//	fmt.Println(v)
	//}
}
var (
	//DB对象
	TestDBObj *DialContext
	TestDBUrl string
	TestDBMinConnNum int
	TestDB string
	TestTbl string
)
func InitTestDB(){
	db,err := Dial(TestDBUrl,TestDBMinConnNum)
	if err != nil {
		fmt.Println("Dial mingodb error:",err)
	}

	TestDBObj = db

	//建立索引
	err = db.EnsureIndex(TestDB,TestTbl,[]string{"userid"})
	if err != nil {
		fmt.Println("ensure index error:",err)
	}
}

func DestroyTestDB(){
	TestDBObj.Close()
	TestDBObj = nil
}

func NewObjectId()(id string){
	objid := bson.NewObjectId()
	id = fmt.Sprintf(`%x`,objid)
	return
}

type Data struct {
	Id string


}

func SendTestDB(data *Data){
	db := TestDBObj.Ref()
	defer TestDBObj.UnRef(db)

	tbl := db.DB(TestDB).C(TestTbl)

	//入库
	_, err := tbl.UpsertId(data.Id, data)
	if err != nil {
		return
	}
}

func GetDataTestDB(id string)*Data{
	db := TestDBObj.Ref()
	defer TestDBObj.UnRef(db)

	tbl := db.DB(TestDB).C(TestTbl)

	data := new(Data)
	err := tbl.FindId(id).One(data)

	if err != nil {
		return nil
	}
	return data
}

func GeyDatasTestDB(uid string)[]*Data{
	db := TestDBObj.Ref()
	defer TestDBObj.UnRef(db)

	tbl := db.DB(TestDB).C(TestTbl)

	var datas []*Data

	err := tbl.Find(bson.M{"uid": uid, "state": 1}).All(&datas)

	if len(datas) == 0 {
		return nil
	}

	if err != nil {
		return nil
	}

	bulk := tbl.Bulk()

	//遍历标记
	for _, data := range datas {
		bulk.UpdateAll(bson.M{"_id": data.Id}, bson.M{"$set": bson.M{"state": 1}})
	}

	_, err = bulk.Run()
	if err != nil {
		return nil
	}
	return datas
}

func GetSumMoneyTestDB(id string)(money int){
	db := TestDBObj.Ref()
	defer TestDBObj.UnRef(db)

	tbl := db.DB(TestDB).C(TestTbl)


	resp := []bson.M{}

	err := tbl.Pipe([]bson.M{
		{"$match": bson.M{"userid": id}},
		{"$group": bson.M{"_id": "$userid", "money": bson.M{"$sum": "$money"}}},
	}).All(&resp)
	if err != nil {
		return
	}
	if len(resp) <= 0 {
		return
	}

	money = resp[0]["money"].(int)
	return
}

