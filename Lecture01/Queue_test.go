package Lecture01

import (
	"fmt"
	"testing"
)

func TestQueue01(t *testing.T) {
	queue := new(Queue)
	fmt.Println(queue.Isempty())
	queue.Push(5)
	queue.Push(3)
	fmt.Println("after push 5:", queue.Isempty())
	fmt.Println("after push 5 the size:", queue.Size())
	fmt.Println("after pop the data:", queue.Pop())
	fmt.Println("after pop the data:", queue.Pop())
}
