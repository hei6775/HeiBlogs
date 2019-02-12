package lecture01

type BagNode struct {
	v    int
	next *BagNode
}

type Bag struct {
	root *BagNode
}

func NewBag(v int) *Bag {
	bag := new(Bag)
	node := new(BagNode)
	node.v = v
	bag.root = node
	return bag
}

func (m *Bag) Add(v int) {
	n := m.root
	for n != nil {
		if n.next == nil {
			break
		}
	}
	bagnode := new(BagNode)
	bagnode.v = v
	n.next = bagnode
}
