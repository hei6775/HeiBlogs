package Lecture01

//队列  数据结构

//节点
type QueueNode struct {
	Item interface{}
	Next *QueueNode
}

//队列结构体
type Queue struct {
	First *QueueNode
	Last  *QueueNode
	Len   int
}

//为空
func (Q *Queue) Isempty() bool {
	if Q.First == nil {
		return true
	}
	return false
}

//入队
func (Q *Queue) Push(item interface{}) {
	oldlast := Q.Last
	newNode := new(QueueNode)
	newNode.Item = item
	Q.Last = newNode
	//magic operation
	if Q.Isempty() {
		Q.First = Q.Last
	} else {
		oldlast.Next = newNode
	}
	Q.Len++
}

//出队
func (Q *Queue) Pop() (item interface{}) {
	if Q.First == nil {
		return nil
	} else {
		item = Q.First.Item
		Q.First = Q.First.Next
		if Q.Isempty() {
			Q.Last = nil
		}
		Q.Len--
	}
	return
}

//大小
func (Q *Queue) Size() (len int) {
	return Q.Len
}
