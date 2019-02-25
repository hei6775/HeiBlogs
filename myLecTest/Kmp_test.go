package myLecTest

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	str := "ababa"
	result := getNext(str)
	fmt.Println(result)
}
