package lecture01

import (
	"fmt"
	"testing"
)

func TestStack01(t *testing.T) {
	stack := new(Stack)
	fmt.Println("isempty:", stack.Isempty(), "size:", stack.Size())

	stack.Push(5)
	fmt.Println("isempty:", stack.Isempty(), "size:", stack.Size())
	stack.Push(3)
	fmt.Println("isempty:", stack.Isempty(), "size:", stack.Size())
	fmt.Println(stack.Pop(), stack.Pop())
}
