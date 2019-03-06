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

func quickSort1(args []int)int{
	pivot := args[0]
	N := len(args)
	minEnd := 1
	maxStart := N-1
	for i:=1;i<maxStart;i++{
		if args[i]<pivot{
			minEnd += 1
		}else{
			args[i],args[maxStart] = args[maxStart],args[i]
			i--
			maxStart -= 1
		}
	}
	args[minEnd-1],args[0] = pivot,args[minEnd-1]
	fmt.Println("result:",args)
	return pivot
}

func TestPrintP1(t *testing.T) {
	a := []int{5,1,3,4,8,1,5,2}
	quickSort1(a)
	fmt.Println(a)
}