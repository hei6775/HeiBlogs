package Lecture01

import (
	"fmt"
	"testing"
)

func TestBS(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("len n :", len(a))
	result := Binary_Search(a, 8)
	fmt.Println("result 8, index:", result)
	result = Binary_Search(a, 3)
	fmt.Println("result 3, index:", result)
	result = Binary_Search(a, 7)
	fmt.Println("result 7, index:", result)
	result = Binary_Search(a, 1)
	fmt.Println("result 1, index:", result)
	result = Binary_Search(a, 9)
	fmt.Println("result 9, index:", result)
	fmt.Println("-----sort包函数实现--------")
	bbqs := BBQS{}
	for i := 1; i < 10; i++ {
		bbq := new(BBQ)
		bbq.v = i
		bbqs = append(bbqs, bbq)
	}
	bbqs.Sort()
	printBBQS(bbqs)
	fmt.Println("the search", bbqs.Search(3))
	fmt.Println("自定义实现二分查找法")
	fmt.Println("the search", bbqs.Search2(3))
}

func printBBQS(bbqs BBQS) (v string) {
	for _, i := range bbqs {
		v = fmt.Sprintf("%s %v", v, i)
	}
	fmt.Println(v)
	return
}
