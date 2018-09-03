package lecture01
//队列结构体
type Queue struct{
	First *Node
	Last *Node
	Len int
}
//为空
func (Q *Queue)Isempty()(bool){
	if Q.First == nil{
		return true
	}
	return false
}
//入队
func (Q *Queue)Push(item interface{})(){
	if Q.First == nil {
		NewNode := new(Node)
		NewNode.Item= item
		Q.First = NewNode
		Q.Last = NewNode
	}else{
		Newnode := new(Node)
		Newnode.Item = Q.Last.Item
		Newnode.Next = Q.Last.Next
		Q.Last.Item = item
		Q.Last.Next =Newnode
	}
	Q.Len++
}
//出队
func (Q *Queue)Pop()(item interface{}){
	if Q.First == nil {
		return nil
	}else {
		NewNode := new(Node)
		item = NewNode.Item
		Q.First.Item = item
		Q.First.Next = nil
		Q.Len--
	}
	return
}
//大小
func (Q *Queue)Size()(len int){
	return Q.Len
}