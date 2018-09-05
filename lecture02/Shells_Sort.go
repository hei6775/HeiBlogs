package main

import "fmt"

func Shell_Sort(args []int)([]int){
	gap := len(args)/2
	time := 0
	for gap >=1 {
		right := 0+gap
		for i:=0;right<len(args);right=i+gap{
			if args[i] > args[right]{
				temp := args[i]
				args[i] = args[right]
				args[right]= temp
				time += 1
			}
			i++
		}
		gap /= 2
	}
	fmt.Println("time :",time)
	return args
}
func BubbleAsort(v []int) {
	time := 0
	for i := 0; i < len(v)-1; i++ {
		for j := i+1; j < len(v); j++ {
			if  v[i]>v[j]{
				v[i],v[j] = v[j],v[i]
				time +=1
			}
		}
	}
	fmt.Println("time :",time)
}
func main(){
	a := []int{5,8,4,7,9,3,2,1,22,6}
	c := []int{5,8,4,7,9,3,2,1,22,6}
	b := Shell_Sort(a)
	BubbleAsort(c)
	fmt.Println(b)
}