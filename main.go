package main

import (
	"fmt"
	"time"
)


func Pic(dx, dy int) [][]uint8 {
	resultArray := make([][]uint8, dx, dx)
	for i, ithRow := range resultArray {
		ithRow = make([]uint8, dy, dy)
		for j := range ithRow {
			if (i+j)%30 <= 15 {
				ithRow[j] = uint8(255)
			}
		}
	}
	return resultArray
}
func main(){
	//fmt.Println(Pic(3,2)[0]==nil)
	//add("s",1,0)
	mychan := make(chan int ,1)

	go func() {
		mychan<-5

	}()
	a := <- mychan
	time.Sleep(2*time.Second)
	fmt.Println(a)
	a = <- mychan
	time.Sleep(2*time.Second)
	return
}

func add(s string,a,b int)(int,error){
	fmt.Println(s,a+b)
	defer fmt.Println("world")
	return fmt.Println("hello")
}