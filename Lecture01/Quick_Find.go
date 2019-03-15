package Lecture01

//快速查找法

type Obj struct {
	intid []int
}

func (O *Obj) Init(N int) {
	for i := 0; i < N; i++ {
		O.intid = append(O.intid, i)
	}
}

//判断是否连通
func (O *Obj) Connect(p, q int) bool {
	return O.intid[p] == O.intid[q]
}

//连通两个节点
func (O *Obj) Union(p, q int) {
	pid := O.intid[p]
	qid := O.intid[q]
	for i := 0; i < len(O.intid); i++ {
		if O.intid[i] == pid {
			O.intid[i] = qid
		}
	}
}
