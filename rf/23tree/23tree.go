//package _3tree
package main
import (
	"fmt"
	"log"
)

//红黑树的特点character
//1.节点不是红色就是黑色
//2.根节点是黑色
//3.空叶节点都是黑色
//4.如果一个节点是红色的，那么它的两个子节点都是黑色的（不会出现连续的两个红色节点）
//5.从任意一个节点到该节点后代的每一个空叶节点的直接路径都要经过相同的个数的黑色节点


//平衡查找树的特点
//1.红链接均为左链接
//2.不会出现两个连续的红色节点
//3.从任意一个节点到该节点后代的每一个空叶节点的直接路径都要经过相同的个数的黑色节点


const (
	Red = true
	Black = false
)
//节点
//-----------------------------------
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

func (this *Node)isRed(n *Node)(bool){
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
	if parent  != nil {
		isleft = this == parent.left
	}

	//左旋
	if isLeftRotate {
		grandleft := this.right.left
		this.color,this.right.color = this.right.color,this.color
		this.right.left = this
		this.parent = this.right
		this.right = grandleft

	}else {
		grandright := this.left.right
		this.color,this.left.color = this.left.color,this.color
		this.left.right = this
		this.parent = this.left
		this.left = grandright


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
		this.parent.parent = parent
	}
	return this.parent,nil
}
//-------------------------------------------------------------------

//23树具体实现
//插入、删除--------------------------------------
type Tree_23 struct {
	root *Node
}

func New23Tree()(*Tree_23){
	tree := new(Tree_23)
	tree.root = nil
	return tree
}
//左旋
func (this *Tree_23)RotateLeft(node *Node)(*Node){
	if h, err := node.rotate(true);err == nil {
		return h
	}
	return nil
}
//右旋
func (this *Tree_23)RotateRight(node *Node)(*Node){
	if h, err := node.rotate(false);err == nil {
		return h
	}
	return nil
}
//找到最右节点
func (this *Tree_23)findRightst(node *Node)(*Node){
	if node.left != nil {
		this.findRightst(node.left)
	}
	return node
}
//flipColors操作，父节点变红，子节点变红
func (this *Tree_23)flipColors(node *Node){
		node.color = Red
		node.right.color = Black
		node.left.color = Black
}
//节点是否为红色节点
func (this *Tree_23)isRed(node *Node)(bool){
	if node != nil && node.color ==Red{
		return Red
	}
	return Black
}

//插入
func (this *Tree_23)Insert(value int){
	if this.root == nil {
		root := NewNode(value)
		this.root = root
		this.root.color = Black
	}else {
		this.insertOne(this.root,value)
	}
}
//插入节点内函数
func (this *Tree_23)insertOne(node *Node,value int)(){
	if node.value > value {
		if node.left != nil {
			this.insertOne(node.left,value)
		}else {
			leftNode := NewNode(value)
			node.left = leftNode
			leftNode.parent = node
			this.insertCheck(leftNode)
			return
		}
	}

	if node.value < value{
		if node.right != nil {
			this.insertOne(node.right,value)
		}else {
			rightNode := NewNode(value)
			node.right = rightNode
			rightNode.parent = node
			this.insertCheck(rightNode)
			return
		}
	}

	if node.value == value{
		fmt.Printf("The value had exist in your heart,扑通 \n")
		return
	}
}
//插入检查
func (this *Tree_23)insertCheck(node *Node)(){
	if node.parent == nil {
		node.color = Black
		this.root = node
		return
	}

	if this.isRed(node) && this.isRed(node.parent) && node == node.parent.right {
		 h := this.RotateLeft(node.parent)
		 node = h.left

	}

	if this.isRed(node) && this.isRed(node.parent){
		node = this.RotateRight(node.parent.parent)
	}

	if this.isRed(node.left) && this.isRed(node.right){
		this.flipColors(node)
		this.insertCheck(node)
	}

	if this.isRed(node) && node.parent.right == node{
		node = this.RotateLeft(node.parent)
		this.insertCheck(node)
	}
}
//删除对外函数
func (this *Tree_23)DelMin(){
	if this.root == nil {
		return
	}
	if !this.isRed(this.root.left) && !this.isRed(this.root.right) {
		//todo I don't know why??
		this.root.color = Red
	}
	this.deleteMin(this.root)
	qq := this.root
	for qq.parent!= nil{
		qq = qq.parent
	}
	this.root = qq
	if this.root != nil {
		this.root.color = Black
	}
}
//删除的flipCplors
//不同的是删除操作中的是将父节点变黑，两个子节点变红
func (this *Tree_23)flashfordel(node *Node){
	node.color = Black
	node.right.color = Red
	node.left.color = Red
}
//for delete
func (this *Tree_23)moveRedLeft(node *Node)(*Node){
	this.flashfordel(node)
	if  this.isRed(node.right.left) {

	}
	if node.right.left != nil && node.right.left.color {
		this.RotateRight(node.right)
		this.RotateLeft(node)
	}
	return node
}

func (this *Tree_23)deleteMin(node *Node)(){
	if (node.left == nil){
		if node == node.parent.left {
			node.parent.left = nil
			node = nil
		}else {
			node.parent.right = nil
			node = nil
		}
		return
	}
	//修复
	if !this.isRed(node.left) && !this.isRed(node.left.left) {
		this.moveRedLeft(node)
	}
	this.deleteMin(node.left)
	this.balance(node)
}

func (this *Tree_23)DelMax(){

}

func (this *Tree_23)deleteMax(node *Node){

}
//平衡
func(this *Tree_23)balance(node *Node)(){
	if this.isRed(node.right) {
		this.RotateLeft(node)
	}
	if this.isRed(node) && this.isRed(node.parent) && node == node.parent.right {
		this.RotateLeft(node.parent)
	}

	if this.isRed(node) && this.isRed(node.parent){
		node = this.RotateRight(node.parent.parent)
	}

	if this.isRed(node.left) && this.isRed(node.right){
		this.flipColors(node)

	}

}


//----------------main函数
func main(){
	tree := New23Tree()
	a := []int{1,7,3,5,4}
	for _,v := range a {
		tree.Insert(v)
	}
	tree.DelMin()
	Print23Tree(tree)
}
//打印树结构
func Print23Tree(tree *Tree_23) {
	if tree.root == nil {
		fmt.Printf("The tree is nil \n")
		return
	}
	printTree(tree.root,"root")
}

func printTree(node *Node,front string){
	if node != nil {
		var colorstr string
		if node.color {
			colorstr = "红"
		}else{
			colorstr = "黑"
		}
		log.Printf(front+"[%d],%s\n", node.value, colorstr)
		printTree(node.left, front+"-(l)|")

		printTree(node.right, front+"-(r)|")
	}
}


