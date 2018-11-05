package lecture03

import (
	"log"
	"testing"
)


func addSon(value int, parent *RBNode, isleft bool) *RBNode {
	son := new(RBNode)
	son.value = value
	son.parent = parent
	if isleft {
		parent.left = son
	} else {
		parent.right = son
	}
	return son
}

func Test_rbnode_rotate(test *testing.T) {
	root := new(RBNode)
	root.value = 1
	l := addSon(2, root, true)
	r := addSon(3, root, false)
	addSon(4, l, true)
	addSon(5, l, false)
	addSon(6, r, true)
	addSon(7, r, false)
	log.Printf("输入数据")
	printTreeInLog(root, "(root)")
	log.Printf("父节点")
	log.Printf("%d", root.left.right.parent.value)
	log.Printf("父节点的兄弟节点")
	log.Printf("%d", root.left.right.getUncle().value)
	log.Printf("祖父节点")
	log.Printf("%d", root.left.right.getGrandParent().value)
	log.Printf("左旋")
	if tmproot, err := root.right.rotate(true); err == nil {
		if tmproot != nil {
			root = tmproot
		}
		printTreeInLog(root, "(root)")
	} else {
		log.Printf(err.Error())
	}

	log.Print("右旋")
	if tmproot, err := root.left.rotate(false); err == nil {
		if tmproot != nil {
			root = tmproot
		}
		printTreeInLog(root, "(r)")
	} else {
		log.Printf(err.Error())
	}

}
