package lecture02

//Bubble Sort Algorithms
func BubbleSort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		for j := i+1; j < len(v); j++ {
			if  v[i]>v[j]{
				v[i],v[j] = v[j],v[i]
			}
		}
	}
}


//Selection Sort Algorithms
func SelSort(inputs []int){
	for i:=0;i<len(inputs);i++{
		minindex := i
		for j:= i+1;j<len(inputs);j++{
			if inputs[minindex]>inputs[j]{
				minindex = j
			}
		}
		inputs[i],inputs[minindex]=inputs[minindex],inputs[i]
	}
}
