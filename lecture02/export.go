package lecture02

type IBinaryHeap interface {
	Less(i, j int) bool
	Swap(i,j int)
	Insert(a int)
	Dele()int
}
