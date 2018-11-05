package lecture03

import "fmt"

//二叉查找树
//平均查找比较1.39lgN
//查找的的成本约比二分查找高了百分之39，插入是线性级别

type BST struct {
	root *Node
}

type Node struct {
	Key int  //key值
	Value string //value值
	Left *Node //左节点
	Right *Node //右节点
	N int //节点数量 本节点也算一个节点
}

//数量
func (b *BST)size()int{
	return b.root.N
}

func (b *BST)initBST(key int,value string){
	root := new(Node)
	root.Key = key
	root.N = 1
	root.Value = value
	b.root = root
}
//初始化一个节点
func initNode(key,n int,value string)(Node){
	return Node{Key:key,Value:value,N:n}
}

//查询一个节点
func (b BST)Search(key int)(string){
	return b.root.Search(key)
}
//查找
func (n Node)Search(key int)(value string){
	if n.N == 1 {
		return value
	}
	if key<n.Key {
		return n.Left.Search(key)
	}else if key == n.Key {
		return n.Value
	}else {
		return n.Right.Search(key)
	}
}
//插入或者更新值
func (b *BST)Put(key int,value string){
	b.root.Put(key,value)
}
//删除
func (b *BST)Delete(key int){
	b.root.delete(key)
}

func (n *Node)delete(key int)(*Node){
	if key < n.Key{
		n.Left = n.Left.delete(key)
		n.N = size(n.Right)+size(n.Left)+1
	}else if key > n.Key {
		n.Right = n.Right.delete(key)
		n.N = size(n.Right)+size(n.Left)+1
	}else {
		if n.Right == nil {
			return n.Left
		}
		if n.Left == nil {
			return n.Right
		}
		temp := new(Node)
		temp.Right = n.Right
		temp.Left = n.Left
		temp.Key = n.Key
		temp.Value = n.Value
		temp.N = n.N
		n = temp.Min()
		n.Right = temp.DeleteMin()
		n.Left = temp.Left
		n.N = size(n.Right)+size(n.Left)+1
	}
	return n
}

//最小
func (b *BST)Min(){
	b.root.Min()
}

//最小
func (n *Node)Min()(*Node){
	if n.Left == nil {
		return n
	}
	return n.Left.Min()
}
//删除最小
func (n *Node)DeleteMin()(*Node){
	if n.Left == nil {
		return n.Right
	}
	goon := n.DeleteMin()
	goon.N = size(n.Right)+size(n.Left)+1
	return goon
}
//插入or 更新
func (n *Node)Put(key int,value string)(*Node){
	if n == nil {
		s := initNode(key,1,value)
		return &s
	}

	if key < n.Key{
		n.Left = n.Left.Put(key,value)

	} else if key > n.Key {
		n.Right = n.Right.Put(key,value)
	}else{
		n.Value = value
	}
	n.N = size(n.Left)+size(n.Right)+1
	return n
}

func size(n *Node)int{
	if n == nil {
		return 0
	}else{
		return n.N
	}
}

func (b *BST)SearchRange(lo,xo int,Que []Node){

	b.root.SearchRange(lo,xo,Que)
}

func (n *Node)SearchRange(lo,xo int,Que []Node){
	if lo> xo {
		fmt.Println("The range is error.")
		return
	}
	return
}


func main(){
	bst := BST{}
	bst.initBST(5,"E")
	bst.Put(1,"A")
	bst.Put(2,"B")
	fmt.Println(bst.root)
	fmt.Println(bst.root.Left)
	bst.Delete(2)
	fmt.Println(bst.root)
	fmt.Println(bst.root.Left)
}


