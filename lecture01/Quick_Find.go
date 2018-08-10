package lecture01

type Obj struct {
	intid []int
}

func (O *Obj)Init(N int){
	for i:=0; i<N; i++{
		O.intid= append(O.intid,i)
	}
}
func (O *Obj)Connect(p,q int)(bool){
	return O.intid[p] == O.intid[q]
}

func (O *Obj)Union(p,q int){
	pid := O.intid[p]
	qid := O.intid[q]
	for i:=0;i<len(O.intid);i++{
		if (O.intid[i] == pid){
			O.intid[i] = qid
		}
	}
}

//func main(){
//	O := new(Obj)
//	O.Init(10)
//	O.Connect(0,2)
//	O.Union(0,2)
//	fmt.Println(O)
//}
