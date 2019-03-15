## 红黑树的插入

&emsp;&emsp;插入节点，二叉查找树插入方法一样，只是在插入操作后做了插入修复


### 红黑树插入修复

1、插入的是根节点，涂黑

2、插入的节点的父节点是黑色，不影响红黑树，不做任何操作

3、插入的节点的父节点是红色

&emsp;&emsp;&emsp;&emsp; - 当前节点的父节点是红色且叔父节点是红色<br/>
````bash
直接涂黑父节点和叔父节点，祖父节点涂红
````
&emsp;&emsp;&emsp;&emsp; - 当前节点的父节点是红色且叔父节点是黑色，当前节点是其父节点的右孩子<br/>
```bash
将父节点作为新的当前节点，再以新的当前节点进行左旋
```
&emsp;&emsp;&emsp;&emsp; - 当前节点的父节点是红色且叔父节点是红色，当前节点是其父节点的左孩子<br/>
```bash
父节点涂黑，祖父节点涂红，以祖父节点为支点进行右旋
```

## 红黑树删除

1、被删除的节点没有子节点，直接删除即可

2、被删除节点只有一个子节点，用子节点替换它即可

3、被删除的节点有两个子节点，找到该节点的右子树的最小节点（非空最左节点），交换这两个节点，然后删除最左节点，（转化为情况1，情况2）

### 红黑树删除修复

&emsp;&emsp;删除节点y之后，x占据了原来节点y的位置。 既然删除y(y是黑色，因为被删除的节点是红色的话，树中黑高度没变，不存在两个相邻的红色节
点，如果删除的节点是红的，就不可能是根)，意味着减少一个黑色节点；那么，再在该位置上增加一个黑色即可。这样，当我们假设"x包含一个额外的黑色"
，就正好弥补了"删除y所丢失的黑色节点"，也就不会违反"特性(5)"。 因此，假设"x包含一个额外的黑色"(x原本的颜色还存在)，这
样就不会违反"特性(5)"。

```go
package lecture03

type TreeNode struct{
    value int
	right,left *TreeNode
}

type  RBTree struct{
	root *TreeNode
}

func(this *RBTree)Delete(data int){
	var (
		deleteNode func(node *TreeNode)
		node *TreeNode = this.Search(data)
		parent *TreeNode
		revise string
	)
	
	if node == nil {
		return
	}else{
		parent = node.parent
	}
	//判断子节点
	if node.left == nil && node.right == nil {
		revise = "none"
	}else if node.parent == nil {
		revise = "parent"
	}else if node == parent.left {
		revise = "left"
	}else if node == parent.right{
		revise = "right"
	}
	//匿名函数
	deleteNode = func(node *TreeNode){
		if node == nil {
			return
		}
		//如果没有子节点
		if node.right == nil && node.left == nil {
			if node == this.root {
				this.root= nil
			}else{
				if node.parent.right == node {
					node.parent.right = nil
				}else{
					node.parent.left = nil
				}
			}
		}
		
		//如果只有一个子节点
		if node.right == nil && node.left != nil {
			//如果节点是根节点
			if node == this.root {
				this.root = node.left
				node.left.parent = nil
			}else{
				if node.parent.left == node {
					node.parent.left = node.left
				}else{
					node.parent.right = node.left
				}
				node.left.parent = node.parent
			}
		}else if node.right != nil && node.left == nil{
			//如果节点是根节点
			if node == this.root {
				this.root = node.right
				node.right.parent = nil
			}else{
				if node.parent.left == node {
					node.parent.left = node.right
				}else{
					node.parent.right = node.right
				}
				node.right.parent = node.parent
				}
		}
		
		//如果有两个节点
		if node.right != nil && node.left != nil {
			minNode := this.SearchMin(node.right)
            node.data = minNode.data
            node.color = minNode.color
            deleteNode(minNode)
		}
		
	}
	deleteNode(node)
	if node.color == "black"{
		if revise == "root" {
			this.deleteBalanceFixup(this.root)
		} else if revise == "left" {
			this.deleteBalanceFixup(parent.left)
		} else if revise == "right" {
			this.deleteBalanceFixup(parent.right)
	}
	
}
```
