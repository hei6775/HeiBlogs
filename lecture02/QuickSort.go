package lecture02

import (
	"math/rand"
	"time"
)

//Quick Sort Algothrims
func QuickSort(inputs []int)([]int){
	inputs = Shuffle(inputs)
	hi:=len(inputs)-1
	sortforquick(0,hi,inputs)
	return inputs
}

func sortforquick(lo,hi int,inputs []int){
	if hi<=lo{
		return
	}
	j := partition(lo,hi,inputs)
	sortforquick(lo,j-1,inputs)
	sortforquick(j+1,hi,inputs)
}
//divide
func partition(lo,hi int,inputs []int)(j int){
	lefP,rigP := lo+1,hi
	v := inputs[lo]
	for {
		for less(inputs[lefP],v){
			lefP++
			if (lefP>=hi){
				break
			}
		}
		for less(v,inputs[rigP]){
			rigP--
			if (rigP<=lo){
				break
			}
		}
		if (lefP>=rigP){
			break
		}
		exch(lefP,rigP,inputs)
	}
	exch(lo,rigP,inputs)
	return rigP
}

//exchange inputs[i], inputs[j]
func exch(i,j int, inputs []int){
	inputs[i],inputs[j] = inputs[j],inputs[i]
}
//随机Shuffle
//return pseudo-random array from vals array
func Shuffle(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]int, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}
