package Lecture02

import (
	"math"
)

//Merge Sort Algorithms
//from top to bottom
func MergeSortT2B(inputs []int){
	lo := 0
	hi := len(inputs)-1
	aux := make([]int,hi+1)
	sort(lo,hi,inputs,aux)
}

//Merge Sort Algorithms
//from bottom to top
func MergeSortB2T(inputs []int){
	hi := len(inputs)-1
	aux := make([]int,hi+1)
	for sz:=1;sz<hi+1;sz = sz+sz{
		for lo:=0;lo<hi+1-sz;lo+=sz+sz{
			min := int(math.Min(float64(lo+sz+sz-1),float64(hi)))
			merge(lo,lo+sz-1,min,inputs,aux)
		}
	}
}

func sort(lo,hi int,inputs,aux []int){
	if hi<=lo {
		return
	}
	mid := lo+(hi-lo)/2
	sort(lo,mid,inputs,aux) //0-3  0-1 2-3
	sort(mid+1,hi,inputs,aux)
	merge(lo,mid,hi,inputs,aux)
}

func merge(lo,mid,hi int,inputs,aux []int){
	i:=lo
	j:=mid+1

	for k:=lo;k<=hi;k++{
		aux[k] = inputs[k]
	}

	for k:=lo;k<=hi;k++{
		if (i>mid){
			inputs[k]=aux[j]
			j++
		}else if (j>hi){
			inputs[k]=aux[i]
			i++
		}else if less(aux[i],aux[j]) {
			inputs[k] = aux[i]
			i++
		}else{
			inputs[k]= aux[j]
			j++
		}
	}

}

func less(a,b int)(bool){
	return a<=b
}
