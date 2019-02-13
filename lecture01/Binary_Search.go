package lecture01

import (
	"fmt"
	"sort"
)

//二分法

func Binary_Search(a []int, tar int) int {
	left := 0
	right := len(a)
	for {
		n := (left + right) / 2
		if tar > a[n] {
			left = n + 1
		} else if tar < a[n] {
			right = n
		} else {
			return n
		}
	}
	return -1
}

//第二种实现方法
type BBQ struct {
	v int
}

type BBQS []*BBQ

func (self BBQS) Len() int           { return len(self) }
func (self BBQS) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }
func (self BBQS) Less(i, j int) bool { return self[i].v < self[j].v }
func (self BBQS) Sort()              { sort.Sort(self) }

//适合从大到小排序，寻找符合条件的最小的index
func (self BBQS) Search(v int) int {
	return sort.Search(len(self), func(i int) bool {
		fmt.Printf("self index[%v] data [%v] \n", i, self[i])
		return self[i].v <= v
	})
}
func (self BBQS) Search2(v int) int {
	left := 0
	right := len(self)
	for n := (left + right) / 2; left < right; n = (left + right) / 2 {
		fmt.Printf("index [%v] data [%v] \n", n, self[n])
		if self[n].v < v {
			left = n + 1
		} else if self[n].v > v {
			right = n
		} else {
			return n
		}
	}
	return -1
}
