package lecture01

import (
	"fmt"
	"testing"
)

func TestBag01(t *testing.T) {
	b := NewBag(5)
	fmt.Println(b.root.v)
	b.Add(3)
	fmt.Println(b.root.v, b.root.next.v)
}
