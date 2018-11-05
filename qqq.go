package main

import (
	"sort"
	"math/rand"
	"strconv"
	"fmt"
)

type BBQ struct {
	v int
	name string
}

type BBQS []*BBQ

func (self BBQS)Len() int {return len(self)}

func (self BBQS)Swap(i,j int){
	self[i],self[j] = self[j],self[i]
}

func (self BBQS)Less(i,j int)bool {
	return self[i].v > self[j].v
}

func (self BBQS)Sort(){sort.Sort(self)}

func (self BBQS)Search(name string)(int,*BBQ){
	for i,v := range self {
		if v.name == name {
			return i,v
		}
	}
	return -1,nil
}

func main(){
	qs := BBQS{}
	for i:=0;i<10;i++{
		m := int(rand.Int63n(20))
		bb := new(BBQ)
		bb.name = strconv.Itoa(m)
		bb.v = m
		qs = append(qs,bb)
		qs.Sort()
	}

	for _,v := range qs {
		fmt.Println(v)
	}
	_,v :=qs.Search("16")
	v.v = 100
	qs.Sort()
	fmt.Println("===============")
	for _,v := range qs {
		fmt.Println(v)
	}
}