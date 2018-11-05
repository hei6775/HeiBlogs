package lecture01

type NewOBJ2 struct {
	intid []int
}

func (Nb *NewOBJ2)Init(N int){
	for i := 0;i<N;i++{
		Nb.intid = append(Nb.intid,i)
	}
}

func (Nb *NewOBJ2)Root(n int)(int){
	for {
		if n == Nb.intid[n] {
			break
		}
		n = Nb.intid[n]
	}
	return n
}

func (Nb *NewOBJ2)Connect(p,q int)(bool){
	return Nb.Root(p) == Nb.Root(q)
}

func (Nb *NewOBJ2)Union(p,q int){

}
