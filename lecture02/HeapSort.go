package lecture02

//Heap Sort Algothrims
func HeapSort(inputs []int)(outputs []int){
	heap := new(BinaryHeap)
	for _,v:= range inputs{
		heap.Insert(v)
	}
	for i:=0;i<len(inputs);i++{
		k := heap.Dele()
		outputs = append(outputs,k)
	}
	return
}
