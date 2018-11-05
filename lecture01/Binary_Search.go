package lecture01

//二分法

func Binary_Search(a []int,tar int)(int){
	left := 0
	right := len(a)
	for {
		n := (left+right)/2
		if tar > a[n]{
			left = n+1
		}else if tar < a[n]{
			right = n
		}else{
			return n
		}
	}
	return -1
}


//func main()  {
//	a := []int{1,2,3,4,5,6,7,8,9}
//	fmt.Println("len n :",len(a))
//	result := Binary_Search(a,5)
//	fmt.Println("result:",result)
//	result = Binary_Search(a,3)
//	fmt.Println("result:",result)
//	result = Binary_Search(a,7)
//	fmt.Println("result:",result)
//	result = Binary_Search(a,1)
//	fmt.Println("result:",result)
//	result = Binary_Search(a,9)
//	fmt.Println("result:",result)
//}