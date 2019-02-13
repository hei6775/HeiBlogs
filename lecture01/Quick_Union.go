package lecture01

//快速连通
type NewObj struct {
	intid []int
}

func (NO *NewObj) Init(N int) {
	for i := 0; i < N; i++ {
		NO.intid = append(NO.intid, i)
	}
}

//寻找根节点
func (NO *NewObj) Root(Y int) int {
	for {
		if Y == NO.intid[Y] {
			break
		}
		Y = NO.intid[Y]
	}
	return Y
}

//判断是否连通
func (NO *NewObj) Connect(p, q int) bool {
	return NO.Root(p) == NO.Root(q)
}

//连通两个节点
func (NO *NewObj) Union(p, q int) {
	i := NO.Root(p)
	j := NO.Root(q)
	NO.intid[i] = j
}
