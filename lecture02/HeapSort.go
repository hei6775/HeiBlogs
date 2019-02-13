package lecture02

import "fmt"

//Heap Sort Algothrims
func HeapSort(inputs []int) (outputs []int) {
	heap := new(BinaryHeap)
	for _, v := range inputs {
		heap.Insert(v)
	}
	fmt.Println(heap)
	for i := 0; i < len(inputs); i++ {
		k := heap.Dele()
		outputs = append(outputs, k)
	}
	return
}
