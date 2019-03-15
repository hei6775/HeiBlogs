package Tree

const (
	Red = true
	Black = false
)
type RBnode struct {
	color bool
	parent,left,right *RBnode
	value int
}

func NewRBnode (data int)(*RBnode){
	return &RBnode{color:Red,value:data}
}
//return 兄弟节点
func (node *RBnode)getSibing()(*RBnode){
	if node.parent == nil {
		return nil
	}

	if node.parent.left == node && node.parent.right != nil {
		return node.parent.right
	}else if node.parent.right == node && node.parent.left != nil {
		return node.parent.left
	}
	return nil
}
//return 祖父节点
func (node *RBnode)getGrandParent()(*RBnode){
	if node.parent== nil || node.parent.parent == nil {
		return nil
	}else {
		return node.parent.parent
	}
}
//return 叔叔节点
func (node *RBnode)getUncle()(*RBnode){
	if node.parent == nil || node.parent.parent == nil {
		return nil
	}
	if node.parent == node.getGrandParent().left && node.getGrandParent().right != nil {
		return node.getGrandParent().right
	}else if node.parent == node.getGrandParent().right && node.getGrandParent().left!= nil {
		return node.getGrandParent().left
	}
	return nil
}
//旋转
func (node *RBnode)rotate(isrotateleft bool)(*RBnode,error){
	var root *RBnode
	//空节点情况
	if node == nil {
		return root,nil
	}

	parent := node.parent
	//判断节点是父节点的左节点
	isleft := false
	if parent != nil {
		isleft = parent.left == node
	}

	//左旋
	if isrotateleft {
		grandsonleft := node.right.left
		//power operating
		node.right.left = node
		node.parent = node.right

		node.right = grandsonleft
	}else {
		//右旋
		grandsonright := node.left.right
		//power operating

		node.left.right = node
		node.parent = node.left

		node.left = grandsonright
	}

	//修复节点与原父节点的关系
	if parent == nil {
		node.parent.parent = nil
		root = node.parent
	}else {
		if isleft {
			parent.left = node.parent
		}else{
			parent.right = node.parent
		}
		node.parent.parent = parent
	}
	return root,nil
}
