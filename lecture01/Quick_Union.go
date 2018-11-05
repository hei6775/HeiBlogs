package lecture01

type NewObj struct{
	intid	[]int
}

func (NO *NewObj)Init(N int){
	for i:=0; i<N;i++{
		NO.intid = append(NO.intid,i)
	}
}

func (NO *NewObj)Root(Y int)(int){
	for {
		if Y == NO.intid[Y]{
			break
		}
		Y =NO.intid[Y]
	}
	return Y
}

func (NO *NewObj)Connect(p,q int)(bool){
	return NO.Root(p)==NO.Root(q)
}

func (NO *NewObj)Union(p,q int){
	i := NO.Root(p)
	j := NO.Root(q)
	NO.intid[i] = j
}

//func main(){
//	N := new(NewObj)
//	N.Init(10)
//	N.Union(5,6)
//	N.Union(4,5)
//	a := N.Root(4)
//	fmt.Println(a)
//	b := N.Connect(1,4)
//	c := N.Connect(4,6)
//	fmt.Println(b,c)
//}