package Lecture02

import (
	"fmt"
	"math"
)

//priority queue
//use array achieve Binary Heap
type BinaryHeap []int

//insert one element to Binary Heap
//Then it start to swim
func (b *BinaryHeap) Insert(a int) {
	*b = append(*b, a)
	b.Swim(len(*b) - 1)
}

//Delete and Return One Element
//And Binary Heap Sink
func (b *BinaryHeap) Dele() int {
	b.Swap(0, len(*b)-1)
	old := *b
	n := len(*b)
	x := old[n-1]
	*b = old[0 : n-1]
	fmt.Println("sink:", b)
	b.Sink(0)
	fmt.Println("del:", b)
	return x
}

//上浮
func (b *BinaryHeap) Swim(k int) {
	for k > 0 && b.Less((k-1)/2, k) {
		b.Swap(k, (k-1)/2)
		k = (k - 1) / 2
	}
}




type MinStack struct {
    data []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}    
}

//     1
//   2    3
// 4  5 6   7
func (this *MinStack)up(curIndex){
    topIndex = (curIndex / 2) -1
    if this.data[topIndex] > this.data[curIndex]{
        this.data[topIndex],this.data[curIndex] = this.data[curIndex],this.data[topIdex]
        if topIndex == 0{
            return
        }else{
            this.up(topIndex)
        }
    }
    return
}


func (this *MinStack) Push(x int)  {
    this.data = append(this.data,x)
    this.up(len(this.data)-1)
}

//     1
//   2    3
// 4  5 6   7
func (this *MinStack)down(curIndex int){
    index := 2*curIndex + 1
    rightIndex := index +1
    if index > len(this.data)-1 {
        return
    }
    //比较左右节点谁小
    if rightIndex >len(this.data)-1 && this.data[index]<this.data[rightIndex]{
        index = rightIndex
    }
    
    this.data[curIndex],this.data[index] = this.data[index],this.data[curIndex]
    
    this.down(index)
}

func (this *MinStack) Pop()  {
    this.data[0],this.data[len(this.data)-1]= this.data[len(this.data)-1],this.data[0]
    x := this.data[len(this.data)-1]
    this.data = this.data[:len(this.data)-1]
    this.down(0)
    return x
}


//下沉
func (b *BinaryHeap) Sink(k int) {
	for k*2 <= len(*b)-1 {
		j := 2*k + 1 //left child
		if j >= len(*b)-1 || j < 0 {
			break
		}
		if j2 := j + 1; j2 < len(*b)-1 && b.Less(j, j2) {
			j = j2 //right child
		}
		if !(b.Less(k, j)) {
			break
		}
		b.Swap(k, j)
		k = j
	}
}

//交换
func (b BinaryHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//比较
func (b BinaryHeap) Less(i, j int) bool {
	return b[i] > b[j]
}

//打印输出
func PrintFormat(arg interface{}) {
	a := arg.(*BinaryHeap)
	if len((*a)) == 0 {
		fmt.Println("no element")
	}
	f := 0
	c := *a
	fla := false
	for i := 0; i < len(*a); i++ {
		conut := math.Exp2(float64(i))
		for v := 0; v < int(conut); v++ {
			g := int(f + v)
			if g >= len(c) {
				fla = true
				break
			}
			fmt.Printf("**%v**", c[g])
		}
		fmt.Printf("\n")
		if fla {
			break
		}
		f += int(conut)
	}
}
