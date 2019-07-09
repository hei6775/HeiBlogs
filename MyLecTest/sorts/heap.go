package sorts

import "fmt"

func TopN(n int,data []int){
	// 先构建n个数的小顶堆
	buildheap(n, data)
	// n往后的数进行调整
	for i := n; i < len(data); i++{
		adjust(i, n, data)
	}
}

//      0
//   1      2
//3    4  5    6
//父节点
func parent(n int)int{
	return (n-1)/2
}
//左节点
func left(n int)int{
	return 2*n+1
}

//右节点
func right(n int)int{
	return 2*n+2
}

//构建堆
func buildheap(n int, args []int){
	for i:=1;i<n;i++ {
		t := i
		for t != 0 && args[parent(t)]>args[t] {
			temp := args[t]
			args[t] = args[parent(t)]
			args[parent(t)] = temp
			t = parent(t)
		}
	}
}
//调整
func adjust(i,n int,args []int){
	if args[i]<= args[0]{
		return
	}
	temp := args[i]
	fmt.Println(temp)
	args[i] = args[0]
	args[0] = temp

	t := 0
	for (left(t)<n && args[t]> args[left(t)]) || (right(t)<n && args[t] > args[right(t)]){
		if right(t)<n && args[t] > args[right(t)]{
			temp = args[t]
			args[t] = args[right(t)]
			args[right(t)] = temp
			t = right(t)
		}else{
			temp = args[t]
			args[t] = args[left(t)]
			args[left(t)] = temp
			t = left(t)
		}
	}
}

//斐波那契数列
func Fibonacci(n int)int{
	startOne,startTwo := 0,1
	for i:= 2;i<n ;i++ {
		startOne,startTwo = startTwo,startOne+startTwo
	}
	return startTwo
}
//斐波那契数列递归法
func Fibonacci2(n int )int{
	if n ==0 || n==1 {
		return n
	}
	return Fibonacci2(n-1)+Fibonacci2(n-2)
}