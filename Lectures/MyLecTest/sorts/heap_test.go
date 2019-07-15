package sorts

import (
	"testing"
	"fmt"
)

func TestTopN(t *testing.T) {
	a := []int{5,3,8,4,9,1,0,6}
	buildheap(6,a)
	fmt.Println(a)
	adjust(7,6,a)
	fmt.Println(a)
}

func BenchmarkFibonacci(b *testing.B) {
	n := 30
	b.ResetTimer()
	for i := 0;i<b.N;i++{
		Fibonacci(n)
	}
}
//但是通过基准测试发现，优化前版本的比优化后版本的更耗时
func BenchmarkFibonacci2(b *testing.B) {
	n := 30
	b.ResetTimer()
	for i := 0;i<b.N;i++{
		Fibonacci2(n)
	}
}