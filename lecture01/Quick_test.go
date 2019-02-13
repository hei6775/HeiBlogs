package lecture01

import (
	"fmt"
	"testing"
)

//快速查找法
func TestQuickFind(t *testing.T) {
	obj := new(Obj)
	obj.Init(10)
	fmt.Println("is node 0 connected node 2:", obj.Connect(0, 2))
	obj.Union(0, 2)
	fmt.Println("is node 0 connected node 2:", obj.Connect(0, 2))

}

//快速连通
func TestQuickUnion(t *testing.T) {
	N := new(NewObj)
	N.Init(10)
	N.Union(5, 6)
	N.Union(4, 5)
	fmt.Println("the root of 4:", N.Root(4))
	b := N.Connect(1, 4)
	c := N.Connect(4, 6)
	fmt.Println("is 1 connected 4:", b)
	fmt.Println("is 4 connected 6:", c)
}

func TestQuickUnionImprove(t *testing.T) {

}
