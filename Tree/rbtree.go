package Tree

import "fmt"

type RBtree struct {
	 root *RBnode
}

func NewRBtree()(*RBtree){
	return &RBtree{nil}
}
//左旋
func (this *RBtree)rotateLeft(node *RBnode){
	if _,err := node.rotate(true);err!= nil {
		fmt.Println(err)
	}
}
//右旋
func (this *RBtree)rotateRight(node *RBnode){
	if _,err := node.rotate(false);err!= nil {
		fmt.Println(err)
	}
}
//找寻node节点下最小的节点
func (this *RBtree)findLeft(node *RBnode)(*RBnode){
	if node.left != nil {
		this.findLeft(node.left)
	}
	return node
}

//插入
func (this *RBtree)Insert(value int){
	if this.root == nil {
		rootnode := NewRBnode(value)
		rootnode.color = Black
		this.root = rootnode
	}else {
		this.InsertNode(this.root,value)
	}

}
//插入
func (this *RBtree)InsertNode(node *RBnode,value int){

	if node.value < value {
		if node.right != nil {
			this.InsertNode(node.right,value)
		}else{
			newnode := NewRBnode(value)
			node.right = newnode
			newnode.parent = node
			this.InserCheck(newnode )
			return
		}
	}
	if node.value > value {
		if node.left != nil {
			this.InsertNode(node.left,value)
		}else {
			newnode := NewRBnode(value)
			node.left = newnode
			newnode.parent = node
			this.InserCheck(newnode)
			return
		}
	}
	if node.value == value {
		fmt.Printf("The value had exist in your heart,扑通 \n")
		return
	}
}
//插入检查
func (this *RBtree)InserCheck(node *RBnode){
	if node.parent == nil {
		node.color = Black
		return
	}

	//父节点为红色的状态
	if node.parent.color == Red {
		if node.getUncle() != nil && node.getUncle().color==Red {
			node.getUncle().color = Black
			node.parent.color =Black
			node.getGrandParent().color = Red
			this.InserCheck(node.getGrandParent())
		}else {
			isleft := node == node.parent.left
			isgrandleft := node.parent == node.getGrandParent().left

			if isleft && isgrandleft {
				this.rotateRight(node.getGrandParent())
				node.parent.color = Black
				node.getUncle().color =Red
			}else if !isleft && !isgrandleft {
				this.rotateLeft(node.getGrandParent())
				node.parent.color = Black
				node.getUncle().color = Red
			}else if !isleft && isgrandleft {
				this.rotateLeft(node.parent)
				this.rotateRight(node.parent)
				node.color = Black
				node.right.color = Red

			}else if isleft && !isgrandleft {
				this.rotateRight(node.parent)
				this.rotateLeft(node.parent)
				node.color = Black
				node.left.color = Red
			}
		}
	}
}


//删除
func (this *RBtree)Delete(value int)(){
	if this.root == nil {
		return
	}
	this.deletenode(value,this.root)
}
func (this *RBtree)deletenode(value int,node *RBnode){

	if node.value < value{
		if node.right == nil {
			return
		}
		this.deletenode(value,node.right)
	}else if node.value>value{
		if node.left == nil {
			return
		}
		this.deletenode(value,node.left)
	}

	if node.left == nil || node.right == nil {
		this.delete_child(node)
		return
	}

	tmpMin := this.findLeft(node.left)
	tmpvalue := tmpMin.value
	tmpMin.value = node.value
	node.value = tmpvalue
	this.delete_child(tmpMin)
	return
}
func (this *RBtree)delete_child(node *RBnode){
	//无子节点
	if node.parent == nil && node.left == nil &&node.right == nil{
		node = nil
		this.root = node
		return
	}
	//子节点
	var child *RBnode
	isadded := false
	if node.left == nil {
		child = node.right
	}else {
		child = node.left
	}
	//有子节点
	if node.parent == nil {
		node = nil
		child.parent = nil
		this.root = child
		this.root.color = Black
		return
	}
	//非根节点
	if node.color == Red{
		if node == node.parent.left {
			node.parent.left = child
		}else {
			node.parent.right = child
		}
		if child != nil {
			child.parent = node.parent
		}
		node = nil
		return
	}
	if child != nil && child.color == Red && node.color ==Black {
		child.parent = node.parent
		if node == node.parent.right {
			node.parent.right = child
		}else {
			node.parent.left = child
		}
		node = nil
		return
	}

	if child == nil {
		child = NewRBnode(0)
		child.parent = node
		isadded =true
	}

	if node.parent.left == node {
		node.parent.left = child
	}else {
		node.parent.right = child
	}
	child.parent = node.parent

	//todo 节点黑色
	if node.color == Black {
		if !isadded && child.color ==Red {
			child.color = Black
		}else {
			this.deleteCheck(child)
		}
	}
	if isadded {
		if child.parent.left == child {
			child.parent.left = nil
		}else {
			child.parent.right = nil
		}
		child = nil
	}
	node = nil
}

func (this *RBtree)deleteCheck(node *RBnode){
	if node.parent == nil {
		node.color = Black
		return
	}

	if  node.getSibing().color == Red{
		node.parent.color = Red
		node.getSibing().color = Black
		if node == node.parent.left{
			this.rotateLeft(node.parent)
		}else {
			this.rotateRight(node.parent)
		}
	}

	is_parent_red := node.parent.color
	is_sib_red := node.getSibing().color
	is_sib_left := Black
	is_sib_right := Black

	if node.getSibing().left != nil {
		is_sib_left = node.getSibing().left.color
	}
	if node.getSibing().right != nil {
		is_sib_right = node.getSibing().right.color
	}
	if !is_parent_red && !is_sib_red &&!is_sib_left && !is_sib_right {
		node.getSibing().color = Red
		this.deleteCheck(node.parent)
		return
	}
	if is_parent_red && !is_sib_red && !is_sib_right && !is_sib_left {
		node.parent.color = Black
		node.getSibing().color =Red
		this.deleteCheck(node.parent)
		return
	}

	if node.getSibing().color ==Black {
		if node == node.parent.left && is_sib_left && !is_sib_right{
			this.rotateRight(node.getSibing())
			node.getSibing().color = Black
			node.getSibing().right.color = Red
		}else if node == node.parent.right && !is_sib_left && is_sib_right{
			node.getSibing().color = Red
			node.getSibing().right.color = Black
			this.rotateLeft(node.getSibing())
		}
	}
	node.getSibing().color = node.parent.color
	node.parent.color = Black
	if node.parent.left == node {
		node.getSibing().right.color = Black
		this.rotateLeft(node.parent)
	}else {
		node.getSibing().left.color = Black
		this.rotateRight(node.parent)
	}

}