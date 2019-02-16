package Tools

import (
	"encoding/json"
	"fmt"
)

type DataInfo struct {
	Id    string  `json:"_id"`
	Total float64 `json:"total"`
}
type Info struct {
	Datas []DataInfo `json:"datas"`
}

func Munmarshal(args []byte) *Info {
	info := new(Info)
	err := json.Unmarshal(args, info)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return info
}

func Mmarshal() []byte {
	info := new(Info)
	ss := []string{"1", "2"}
	sss := []float64{3, 4}
	for k, v := range ss {
		ii := new(DataInfo)
		ii.Id = v
		ii.Total = sss[k]
		info.Datas = append(info.Datas, *ii)
	}
	result, _ := json.Marshal(info)
	return result
}
