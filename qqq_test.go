package test

import (
	"testing"


	"sort"
	"fmt"
	"ss/sssj/src/github.com/name5566/leaf/log"
)

type testRank struct {
	rank int
}
type RankDataSlice [] *testRank

func (p RankDataSlice) Len() int           { return len(p) }
func (p RankDataSlice) Less(i, j int) bool { return p[i].rank < p[j].rank }
func (p RankDataSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p RankDataSlice) Sort() { sort.Sort(p) }

func (p RankDataSlice) Search(x int) int {
	return sort.Search(len(p), func(i int) bool { return p[i].rank < x  })
}


type RankDataSliceb [] int

func (p RankDataSliceb) Len() int           { return len(p) }
func (p RankDataSliceb) Less(i, j int) bool { return p[i] <= p[j] }
func (p RankDataSliceb) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p RankDataSliceb) Sort() { sort.Sort(p) }

func (p RankDataSliceb) Search(x int) int {
	return sort.Search(len(p), func(i int) bool { return p[i] == x  })
}


func TestPhoneCertify(test *testing.T) {


	var a RankDataSlice
	var b RankDataSliceb
	for i:=0;i<20;i++{
		temp  := testRank{
			rank:i,
		}
		//log.Debug("a is:%v",temp)
		a = append(a, &temp)
		b = append(b,i)
	}
	a.Sort()
	b.Sort()
	//d := sort.Search(len(c), func(i int) bool { return c[i] == 1 })

	e := SearchMy(b,4,0)
	log.Debug("b.SearchMy(4) is:%d",e)

	//log.Debug("a is:%v", a, a.Search(5))
	//log.Debug("b is:%v", b, a.Search(3))
	//log.Debug("d is:%v", c, d)


}

func  SearchMy(p []int ,i int,indexInt int) (index int) {
	if indexInt == 0{
		indexInt = len(p)/2
	}
	fmt.Println("indexInt",p,indexInt)
	fmt.Println("index",index)
	temp := p[indexInt]
	if temp == i {
		index = indexInt
		fmt.Println("等于",index)
		return index
	}else  {
		if temp > i{
			indexInt  -= indexInt/2
			fmt.Println("大于",indexInt)
			index = SearchMy(p,i,indexInt)
			return
		}else {
			indexInt += indexInt/2
			fmt.Println("小于",indexInt)
			index = SearchMy(p,i,indexInt)
			return
		}
	}
	log.Debug("返回 ",index)
	return index
}