package Lecture01

//先进后出
//栈
type Stack struct {
	First *StackNode
	Len   int
}

//节点
type StackNode struct {
	Item interface{}
	Next *StackNode
}

//为空
func (S *Stack) Isempty() bool {
	if S.Len == 0 || S.First == nil {
		return true
	}
	return false
}

//入栈
func (S *Stack) Push(item interface{}) {
	newNode := new(StackNode)
	newNode.Item = item
	if S.First == nil {
		S.First = newNode
		S.Len++
		return
	}
	oldNode := S.First
	newNode.Next = oldNode
	S.First = newNode
	S.Len++
}

//出栈
func (S *Stack) Pop() (item interface{}) {
	if S.Isempty() {
		return
	}
	item = S.First.Item
	S.First = S.First.Next
	S.Len--
	return
}

//大小
func (S *Stack) Size() (size int) {
	return S.Len
}
