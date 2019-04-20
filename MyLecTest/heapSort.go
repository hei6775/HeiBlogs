package MyLecTest

import "container/heap"

type Openll []int

func (self Openll) Len() int            { return len(self) }
func (self Openll) Swap(i, j int)       { self[i], self[j] = self[j], self[i] }
func (self Openll) Less(i, j int) bool  { return self[i] > self[j] }
func (self *Openll) Push(x interface{}) { *self = append(*self, x.(int)) }
func (self *Openll) Pop() interface{} {
	old := *self
	n := self.Len()
	x := old[n-1]
	*self = old[0 : n-1]
	return x
}

//堆排序
func TestHeapSort() {
	temp := Openll{9,8,7,15,13,17,19}
	heap.Init(&temp)
	heap.Push(&temp,5)
	heap.Pop(&temp)
}
