package lecture02

import (
	"fmt"
	"testing"
)

//1,3,5,7,9,11,13,15
var arr = []int{11, 3, 15, 5, 1, 7, 13, 9}
var arrtest = make([]int, len(arr))

//冒泡排序
func TestBubble(t *testing.T) {
	copy(arrtest, arr)
	BubbleSort(arrtest)
	fmt.Println("Bubble sort result:", arrtest)
}

//选择排序
func TestSelect(t *testing.T) {
	copy(arrtest, arr)
	SelSort(arrtest)
	fmt.Println("Select sort result:", arrtest)
}

//插入排序
func TestInsert(t *testing.T) {
	copy(arrtest, arr)
	InsertSort(arrtest)
	fmt.Println("Insert sort result:", arrtest)
}

//堆排序
//有错
func TestHeap(t *testing.T) {
	copy(arrtest, arr)
	arrtest = HeapSort(arrtest)
	fmt.Println("Heap sort result:", arrtest)
}

//希尔排序
func TestShell(t *testing.T) {
	copy(arrtest, arr)
	arrtest = Shell_Sort(arrtest)
	fmt.Println("Shell sort result:", arrtest)
}

//归并排序
func TestMerge(t *testing.T) {
	copy(arrtest, arr)
	MergeSortT2B(arrtest)
	fmt.Println("MergeSortT2B result:", arrtest)
	copy(arrtest, arr)
	MergeSortB2T(arrtest)
	fmt.Println("MergeSortB2T result:", arrtest)
}

//快速排序
func TestQuick(t *testing.T) {
	copy(arrtest, arr)
	arrtest = QuickSort(arrtest)
	fmt.Println("QuickSort result:", arrtest)
}

//
func TestZuoyi(t *testing.T) {
	fmt.Println("1右移2位:", 1>>2) //0001 >> 0000
	fmt.Println("1左移2位", 1<<2)  //0001 << 0100
}
