package lecture01

//先进后出
//栈
type Stack struct {
	First *Node
	Len int
}
//节点
type Node struct {
	Item interface{}
	Next *Node
}
//为空
func (S *Stack)Isempty()(bool){
	if S.Len == 0 || S.First==nil{
		return true
	}
	return false
}
//入栈
func (S *Stack)Push(item interface{}){
	NewNode := new(Node)
	NewNode.Item = item
	if S.First == nil {
		S.First = NewNode
		S.Len++
		return
	}
	OldNode := Node{}
	OldNode.Item = S.First.Item
	NewNode.Next = &OldNode
	S.First = NewNode
	S.Len++
}
//出栈
func (S *Stack)Pop()(item interface{}){
	item = S.First.Item
	S.First = S.First.Next
	S.Len--
	return
}
//大小
func (S *Stack)Size()(size int){
	return S.Len
}