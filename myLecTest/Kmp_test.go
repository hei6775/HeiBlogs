package myLecTest

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	str := "ABAD"
	result := getNext(str)
	fmt.Println(result)
}
