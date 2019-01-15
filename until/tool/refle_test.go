package tool

import (
	"fmt"
	"testing"
)

var (
	a int      = 1
	b bool     = true
	c float64  = 64
	d string   = "i'm string"
	e Myint8   = 8
	f Mystruct = Mystruct{1, false}
)

type Myint8 int8

type Mystruct struct {
	A int
	B bool
}

func TestRefle(t *testing.T) {
	aname := GetTypeName(a)
	bname := GetTypeName(b)
	cname := GetTypeName(c)
	dname := GetTypeName(d)
	ename := GetTypeName(e)
	fname := GetTypeName(f)

	fmt.Printf("aname[%v] bname[%v] cname[%v] \n dname[%v] ename[%v] fname[%v] \n", aname, bname, cname,
		dname, ename, fname)
}
