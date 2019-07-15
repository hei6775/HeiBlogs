package Lecture03

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

//Red Black Tree
//It's chraacter
//1. 节点不是红色就是黑色
//2.根节点是黑色的
//3.空叶子节点都是黑色的
//4.如果一个节点是红色的，那么它的两个子节点都是黑色的（不会出现连续的两个红色节点）
//5.从任意一个节点到这个节点后代的每一个空叶子节点的直接路径都要经过相同个数的黑色节点。
// 这里有几个定义：从根节点到某个节点所经过的黑色节点的数目成为这个节点的 black depth；从根节点到叶子节点的所经过的黑色节点数目成为这个树的 black-height

//note:每次插入的节点初始时都是红色的

const (
	Black = false //黑色
	Red   = true  //红色
)

//节点结构
type RBNode struct {
	n                   int     //子节点数量
	value               int     //value值
	color               bool    //颜色 false 黑 red 红
	right, left, parent *RBNode //右节点
}

//return new RBnode(init coldr is red)
func NewRBNode(value int) *RBNode {
	return &RBNode{color: Red, value: value}
}

//return grandparent RBNode
func (rbnode *RBNode) getGrandParent() *RBNode {
	if rbnode.parent == nil || rbnode.parent.parent == nil {
		return nil
	}
	return rbnode.parent.parent
}

//return brother RBNode
func (rbnode *RBNode) getSibling() *RBNode {
	if rbnode.parent == nil {
		return nil
	}
	pa := rbnode.parent
	if pa.right == rbnode && pa.left != nil {
		return pa.left
	} else if pa.left == rbnode && pa.right != nil {
		return pa.right
	}
	return nil
}

//return Uncle RBnode
func (rbnode *RBNode) getUncle() *RBNode {
	if rbnode.parent == nil || rbnode.parent.parent == nil {
		return nil
	}
	pa := rbnode.parent
	grandpa := rbnode.parent.parent
	if grandpa.right == pa && grandpa.left != nil {
		return grandpa.left
	} else if grandpa.left == pa && grandpa.right != nil {
		return grandpa.right
	}
	return nil
}

//right or left rotate
//-----------A------------------
//--------B-----C---------------
//------A--E--------------------
//----F--G----------------------
//------------------------------
//------------------------------
func (rbnode *RBNode) rotate(isRotateLeft bool) (*RBNode, error) {
	var root *RBNode

	if rbnode == nil {
		return root, nil
	}
	if !isRotateLeft && rbnode.left == nil {
		return root, errors.New("The Right Rotate Left Node can't be nil")
	} else if isRotateLeft && rbnode.right == nil {
		return root, errors.New("The Left Rotate Right Node can't be nil")
	}

	parent := rbnode.parent
	//判断是否左边
	var isleft bool
	if parent != nil {
		isleft = parent.left == rbnode
	}

	//左旋
	if isRotateLeft {
		grandsonleft := rbnode.right.left
		//magic operation
		rbnode.right.left = rbnode
		rbnode.parent = rbnode.right
		rbnode.right = grandsonleft
	} else {
		//右旋
		//记录孙节点
		grandsonright := rbnode.left.right
		//magic operation
		//要移动的孙节点 将其赋值为节点
		//节点的父节点设置为节点的左节点
		//节点的左节点设置为之前保存的孙节点
		rbnode.left.right = rbnode
		rbnode.parent = rbnode.left
		rbnode.left = grandsonright
	}

	//
	if parent == nil {
		rbnode.parent.parent = nil
		root = rbnode.parent
	} else {
		if isleft {
			parent.left = rbnode.parent
		} else {
			parent.right = rbnode.parent
		}
		rbnode.parent.parent = parent
	}
	return root, nil
}

//return the most left rbnode
func (rbnode *RBNode) getMostLeftNode() *RBNode {
	if rbnode.left == nil {
		return rbnode
	}
	return rbnode.left.getMostLeftNode()
}

//红黑树结构
type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{root: nil}
}
func (rbtree *RBTree) Insert(data int) {
	if rbtree.root == nil {
		rootnode := NewRBNode(data)
		rootnode.color = Black
		rbtree.root = rootnode
	} else {
		rbtree.insertNode(rbtree.root, data)
	}
}

func (rbtree *RBTree) Delete(data int) {
	rbtree.delete_child(rbtree.root, data)
}

//插入操作
func (rbtree *RBTree) insertNode(pnode *RBNode, data int) {
	if pnode.value == data {
		fmt.Println("insert the data:", data, "equal the value", pnode.value)
		return
	}
	if pnode.value > data {
		if pnode.left != nil {
			rbtree.insertNode(pnode.left, data)
		} else {
			tmpnode := NewRBNode(data)
			tmpnode.parent = pnode
			pnode.left = tmpnode
			rbtree.insertCheck(tmpnode)
		}
	} else {
		if pnode.right != nil {
			rbtree.insertNode(pnode.right, data)
		} else {
			tmpnode := NewRBNode(data)
			tmpnode.parent = pnode
			pnode.right = tmpnode
			rbtree.insertCheck(tmpnode)
		}
	}
}

//左旋
func (rbtree *RBTree) rotateLeft(node *RBNode) {
	if tmproot, err := node.rotate(true); err == nil {
		if tmproot != nil {
			rbtree.root = tmproot
		}
	} else {
		log.Printf(err.Error())
	}
}

//右旋
func (rbtree *RBTree) rotateRight(node *RBNode) {
	if tmproot, err := node.rotate(false); err == nil {
		if tmproot != nil {
			rbtree.root = tmproot
		}
	} else {
		log.Printf(err.Error())
	}
}

//插入节点情况分析：
//① 被插入的节点是根节点，直接将根节点颜色改黑
//② 被插入的节点的父节点是黑色的，无需操作
//③ 被插入的节点的父节点是红色的
//---<1>被插入的节点的父节点为红色，且叔叔节点也是红色，then对祖父节点进行递归处理
//---<2>被插入的节点的父节点为红色，且叔叔节点是黑色，当前节点为父节点的右孩子，then进入旋转平衡
//---<3>被插入的节点的父节点为红色，且叔叔节点是黑色，当前节点为父节点的左孩子，then进入旋转平衡

//插入检查
func (rbtree *RBTree) insertCheck(node *RBNode) {
	if node.parent == nil {
		rbtree.root = node
		rbtree.root.color = Black
		return
	}

	if node.parent.color == Red {
		if node.getUncle() != nil && node.getUncle().color == Red {
			node.parent.color = Black
			node.getUncle().color = Black
			node.getGrandParent().color = Red
			rbtree.insertCheck(node.getGrandParent())
		} else {
			isleft := node == node.parent.left
			isparentleft := node.parent == node.getGrandParent().left
			if isleft && isparentleft {
				rbtree.rotateRight(node.getGrandParent())
				node.parent.color = Black
				node.parent.right.color = Red
			} else if !isleft && isparentleft {
				rbtree.rotateLeft(node.parent)
				rbtree.rotateRight(node.parent)
				node.color = Black
				node.right.color = Red
				//node.left.color =Red //??? what
			} else if !isleft && !isparentleft {
				rbtree.rotateLeft(node.getGrandParent())
				node.parent.color = Black
				node.getSibling().color = Red
			} else if isleft && !isparentleft {
				rbtree.rotateRight(node.parent)
				rbtree.rotateLeft(node.parent)
				node.color = Black
				node.left.color = Red
			}
		}
	}
}

//节点的删除
//==========================
//1、节点没有子节点，可以直接删除
//2、有一个子节点，直接将子节点顶替该节点
//3、有两个子节点
//      a) 选择右子树中的节点作替换（左右无所谓，道理是一样的，我们以替换右子树中的节点为例）
//      b) 找到节点N右子树中最靠左边的非空节点M（不一定是叶子节点），将M的值复制到N，然后把M节点改为要删除的节点的值，则转换成了删除节点中
//         包含有一个子节点（或没有子节点）的问题，例如下图中我们如果要删除5，则寻找节点7的最左侧的非空节点6，将6复制到5的位置，然后排序
//         树仍然成立（原来的6要删除了，不考虑在内）

//以上操作之后都变成删除单个子节点或者没有子节点
//		1、删除的是根节点，且节点没有子节点，则直接删除，root置空
//		2、删除的是根节点，且节点只有一个子节点，直接删除，将根节点指向子节点，并设置颜色为黑色
//		3、删除的不是根节点
//			a)要删除的节点是红色，则它的父节点与子节点（如果有的话）都是黑色，直接把子节点或空节点替换要删除节点的位置即可，不会影响黑节点的平衡
//			b) 要删除的是黑色节点，但它的子节点存在且是红色，我们要做到只是用子节点替换它，然后改变子节点为黑色，使黑色平衡
//			c) 要删除的是黑色节点且子节点也是黑色，这就需要检验平衡了，检查替换后的孩子节点

//删除n节点下值为data的节点
func (rbtree *RBTree) delete_child(n *RBNode, data int) bool {
	//左节点
	if data < n.value {
		if n.left == nil {
			return false
		}
		return rbtree.delete_child(n.left, data)
	}
	//右节点
	if data > n.value {
		if n.right == nil {
			return false
		}
		return rbtree.delete_child(n.right, data)
	}
	//查找到该节点
	if n.right == nil || n.left == nil {
		rbtree.delete_one(n)
		return true
	}

	//寻找n节点的右节点的最左后代节点
	tmpmostleft := n.left.getMostLeftNode()
	tmpvalue := tmpmostleft.value
	tmpmostleft.value = n.value
	n.value = tmpvalue

	rbtree.delete_one(tmpmostleft)
	return true
}

//删除单个节点
func (rbtree *RBTree) delete_one(node *RBNode) {
	//todo why
	//可能导致child 为 nil
	var child *RBNode
	isadded := false
	if node.left == nil {
		child = node.right
	} else {
		child = node.left
	}

	//如果是根节点的情况一
	if node.parent == nil && node.left == nil && node.right == nil {
		node = nil
		rbtree.root = node
		return
	}
	//如果是根节点的情况二
	if node.parent == nil {
		node = nil
		child.parent = nil
		rbtree.root = child
		rbtree.root.color = Black
		return
	}
	//todo 红色可以直接删
	if node.color == Red {
		if node == node.parent.left {
			node.parent.left = child
		} else {
			node.parent.right = child
		}
		if child != nil {
			child.parent = node.parent
		}
		//删除
		node = nil
		return
	}
	if child != nil && child.color == Red && node.color == Black {
		if node == node.parent.left {
			node.parent.left = child

		} else {
			node.parent.right = child
		}
		child.parent = node.parent
		child.color = Black
		node = nil
		return
	}
	//删除节点是黑色
	if child == nil {
		child = NewRBNode(0)
		child.parent = node
		isadded = true
	}
	if node.parent.left == node {
		node.parent.left = child
	} else {
		node.parent.right = child
	}
	child.parent = node.parent
	//todo 节点是黑色的话
	if node.color == Black {
		if !isadded && child.color == Red {
			child.color = Black
		} else {
			rbtree.deleteCheck(child)
		}
	}

	if isadded {
		if child.parent.left == child {
			child.parent.left = nil
		} else {
			child.parent.right = nil
		}
		child = nil
	}
	node = nil
}

//删除检查
func (rbtree *RBTree) deleteCheck(n *RBNode) {

	if n.parent == nil {
		n.color = Black
		return
	}
	//兄弟节点为红色
	if n.getSibling().color == Red {
		n.parent.color = Red
		n.getSibling().color = Black
		if n == n.parent.left {
			rbtree.rotateLeft(n.parent)
		} else {
			rbtree.rotateRight(n.parent)
		}
	}

	//注意：这里n的兄弟节点发生了变化，不再是原来的兄弟节点
	is_parent_red := n.parent.color
	is_sib_red := n.getSibling().color
	is_sib_left_red := Black
	is_sib_right_red := Black

	//兄弟节点的左右子节点 颜色判断
	if n.getSibling().left != nil {
		is_sib_left_red = n.getSibling().left.color
	}
	if n.getSibling().right != nil {
		is_sib_right_red = n.getSibling().right.color
	}

	if !is_parent_red && !is_sib_red && !is_sib_right_red && !is_sib_left_red {
		n.getSibling().color = Red
		rbtree.deleteCheck(n.parent)
		return
	}

	if is_parent_red && !is_sib_red && !is_sib_left_red && !is_sib_right_red {
		n.getSibling().color = Red
		n.parent.color = Black
		return
	}

	if n.getSibling().color == Black {
		if n.parent.left == n && is_sib_left_red && !is_sib_right_red {
			n.getSibling().color = Red
			n.getSibling().left.color = Black
			rbtree.rotateRight(n.getSibling())
		} else if n.parent.right == n && !is_sib_left_red && is_sib_right_red {
			n.getSibling().color = Red
			n.getSibling().right.color = Black
			rbtree.rotateLeft(n.getSibling())
		}
	}
	n.getSibling().color = n.parent.color
	n.parent.color = Black
	if n.parent.left == n {
		n.getSibling().right.color = Black
		rbtree.rotateLeft(n.parent)
	} else {
		n.getSibling().left.color = Black
		rbtree.rotateRight(n.parent)
	}
}

// log输出树
func printTreeInLog(n *RBNode, front string) {
	if n != nil {
		var colorstr string
		if n.color == Red {
			colorstr = "红"
		} else {
			colorstr = "黑"
		}
		log.Printf(front+"%d,%s\n", n.value, colorstr)
		printTreeInLog(n.left, front+"-(l)|")

		printTreeInLog(n.right, front+"-(r)|")
	}
}
