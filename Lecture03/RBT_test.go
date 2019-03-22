package Lecture03

import "errors"

func (this *RBNode)Rotate123(isLeftRotate bool)(*RBNode,error){
	root := new(RBNode)

	if this == nil {
		return root,nil
	}

	if !isLeftRotate && this.left == nil {
		return nil,errors.New("right rotate:the node left is nil")
	}else if isLeftRotate && this.right == nil {
		return nil,errors.New("left rotate:the node right is nil")
	}

	parent := this.parent

	var isleft bool

	if parent != nil {
		isleft = parent.left == this
	}
	//
	if isLeftRotate {
		grandson := this.right.left

		this.right.left = this
		this.parent = this.right
		this.right = grandson
	}else{
		grandson := this.left.right

		this.left.right = this
		this.parent = this.left
		this.left = grandson
	}

	if parent == nil{
		this.parent.parent = nil
		root = this.parent
	}else{
		if isleft{
			p
		}else{

		}

	}



}