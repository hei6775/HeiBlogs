package lecture02

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := new(BinaryHeap)
	heap.Insert(1)
	heap.Insert(2)
	fmt.Println(heap)
}
