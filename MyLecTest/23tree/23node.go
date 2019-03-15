package _3tree
//package main
const (
	Red = true
	Black = false
)

type Node struct {
	value int
	color bool
	parent,left,right *Node
}

func NewNode(value int)(*Node){
	newnode := new(Node)
	newnode.value = value
	newnode.color = Red
	return newnode
}

func (this *Node)isRed (n *Node)(bool){
	if n == nil {
		return false
	}
	return n.color == Red
}

func (this *Node)getSibling()(*Node){
	if this.parent == nil {
		return nil
	}
	if this == this.parent.left  && this.parent.right != nil{
		return this.parent.right
	}else if this == this.parent.right && this.parent.left !=nil{
		return this.parent.left
	}
	return nil
}

func (this *Node)getUncle()(*Node){
	if this.parent == nil || this.parent.parent == nil {
		return nil
	}
	if this.parent == this.parent.parent.left  && this.parent.parent.right != nil{
		return this.parent.parent.right
	}else if this.parent == this.parent.parent.right && this.parent.parent.left !=nil{
		return this.parent.parent.left
	}
	return nil

}

func (this *Node)rotate(isLeftRotate bool)(*Node,error){
	var root *Node

	if this == nil {
		return root,nil
	}

	parent := this.parent
	isleft := false
	if parent  == nil {
		isleft = this == parent.left
	}

	//左旋
	if isLeftRotate {
		grandleft := this.right.left

		this.right.left = this
		this.right = grandleft
		this.parent = this.right
	}else {
		grandright := this.left.right

		this.left.right = this
		this.left = grandright
		this.parent = this.left

	}

	if parent == nil {
		this.parent.parent = nil
		root = this.parent
	}else {
		if isleft {
			parent.left = this.parent
		}else {
			parent.right =this.parent
		}
		this.parent.parent =parent
	}
	return root,nil
}