package myLecTest

import (
	"testing"
	"fmt"
)

func TestReverseP(t *testing.T) {
	head := new(Point)
	head.Val = 0
	s := head
	for i:=1;i<10;i++{
		point := new(Point)
		point.Val = i
		s.Next = point
		s = point
	}
	fmt.Println("1111")
	PrintP(head)
	fmt.Println(" \n")
	c := ReverseP(head)
	PrintP(c)
}