package main

import (
	"math/rand"
	"time"
	"fmt"
)
func Shuffle(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]int, len(vals))
	perm := r.Perm(len(vals))
	fmt.Println(perm)
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}


func main() {
	vals := []int{10, 12, 14, 16, 18, 20}
	vals=Shuffle(vals)
	fmt.Println(vals)
}
