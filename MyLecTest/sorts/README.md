## 各个排序算法

#### 冒泡排序

&emsp;&emsp;Description：冒泡排序法时间复杂度，平均情况O(n^2)，最好情况O(n)，最坏情况O(n^2)。空间复杂度O(1)。

```go
package  sorts
//冒泡排序
func BubbleSort(args []int){
	N := len(args)
	for i := 0;i<N;i++{
		for j := 0;j<N-1;j++{
			if args[j] > args[j+1]{
				args[j],args[j+1] = args[j+1],args[j]
			}
		}
	}
}
```

&emsp;&emsp;优化，可以加一个标志位，在某一轮遍历中如果没有互换则表示已经是有序的了

#### 选择排序

&emsp;&emsp;Description：选择排序法时间复杂度，平均情况O(n^2)，最好情况O(n^2)，最坏情况O(n^2)。空间复杂度O(1)。

```go
package sorts

func SelectSort(args []int){
	N:= len(args)
	for i := 0;i<N;i++{
		minIndex := i
		for j:=i+1;j<N;j++{
			if args[minIndex]>args[j]{
				minIndex = j
			}
		}
		args[i],args[minIndex] = args[minIndex],args[i]
	}
}
```

#### 插入排序

&emsp;&emsp;Description：插入排序法时间复杂度，平均情况O(n^2)，最好情况O(n)，最坏情况O(n^2)。空间复杂度O(1)。

```go
package sorts

func InsertSort(args []int){
	N := len(args)
	
	for i:= 0;i<N;i++{
		for j:=1;j>0 && j<N;j--{
			if args[j] < args[j-1] {
				args[j],args[j-1] = args[j-1],args[j]
			}
		}
	}
}
```

#### 希尔排序

&emsp;&emsp;Description：希尔排序法时间复杂度，平均情况O(n^1.3)，最好情况O(n)，最坏情况O(n^2)。空间复杂度O(1)。

```go
package sorts

func ShellSort(args []int){
	N := len(args)
	gap := 1
	for (gap>(N/3)){
        gap = gap*3+1
	}
	
	for gap >=1{
		for i:=gap;i<N;i++{
			for j:=i;j>gap&&args[j]<args[j-gap];j-=gap{
				args[j],args[j-gap]=args[j-gap],args[j]
			}
		}
		gap /=3
	}
	
}
```

#### 堆排序

&emsp;&emsp;Description：堆排序法时间复杂度，平均情况O(n*logn)`(底数为2)`，最好情况O(n*logn)`(底数为2)`，最坏情
况O(n*logn)`(底数为2)`。空间复杂度O(1)。

```go
package sorts
```

#### 归并排序

&emsp;&emsp;Description：归并排序法时间复杂度，平均情况O(n*logn)`(底数为2)`，最好情况O(n*logn)`(底数为2)`，最坏情
况O(n*logn)`(底数为2)`。空间复杂度O(1)。


```go
package sorts

func Swap(arr []int,i,j int) {
	arr[i],arr[j] = arr[j],arr[i]
}

func Less(arr []int,i,j int) bool {
	return arr[i] > arr[j]
}

func Merge(arr []int){
	mid := len(arr)/2
	lo := 0
	lr := mid
	N := len(arr)
	merArr := make([]int,N)
	merArr = arr
	
	for i:=0;i<N;i++ {
		if lo >= mid {
			arr[i] = merArr[lr]
			lr++
		}else if lr >= N {
			arr[i] = merArr[lo]
			lo++
		}else if merArr[lo] > merArr[lr]{
			arr[i] = merArr[lr]
			lr++
		}else{
			arr[i] = merArr[lo]
			lo++
		}
	}
}

func MergeSort(args []int){
	if len(args) <=1 {
		return
	}
	mid := len(args) /2
	MergeSort(args[0:mid])
	MergeSort(args[mid:])
	Merge(args)
}
```

#### 快速排序——单路排序

&emsp;&emsp;Description：快速排序法时间复杂度，平均情况O(n*logn)`(底数为2)`，最好情况O(n*logn)`(底数为2)`，最坏情
况O(n^2)。空间复杂度O(nlogn)`(底数为2)`。


```go
package sorts

func QuickSortOne(args []int){
	N := len(args)
	if N <=1 {
		return
	}
	j := quickSort1(args)
	QuickSortOne(args[0:j])
	QuickSortOne(args[j+1:])
}
// 2 3 1
//  
func quickSort1(args []int)int{
	pivot := args[0]
	N := len(args)
	minEnd := 1
	maxStart := N-1
	for i:=1;i <maxStart;i++{
		if args[i]<pivot{
			minEnd += 1
		}else{
			args[i],args[maxStart] = args[maxStart],args[i]
			i--
			maxStart -= 1
		}
	}
	args[minEnd-1],args[0] = pivot,args[minEnd-1]
	return pivot
}
```
#### 快速排序——双路排序
```go
package sorts

func QuickSortTwo(args []int){
	N := len(args)
	if N <= 1 {
		return
	}
	j := quickSort2(args)
	QuickSortTwo(args[0:j])
	QuickSortTwo(args[j:])
}

func quickSort2(args []int)int{
	N := len(args)
	pivot := 0
	right := N-1
	left := 0
	
	for right > left {
		if args[left]>args[pivot] || args[right]<args[pivot] {
			if args[left]>args[pivot] && args[right]<args[pivot] {
				args[left],args[right] = args[right],args[left]
				left++
				right--
				continue
			}
			if args[left] > args[pivot] {
				right--
				continue
			}else{
				left++
				continue 
			}
		}
		right--
		left++
	}
	args[pivot],args[left] = args[left],args[pivot]
	return left
}
```
#### 快速排序——三路排序
```go
package sorts

func QuickSortThree(head,tail int,args []int){
	if tail < head {
		return
	}
	
	var p,q,i,j int
	i = head      //从左向右扫描
	p = head      
	q = tail - 1  //从右向左扫描
	j = tail - 1
	//比较值
	pivot := args[tail]
	
	for {
		// 工作指针i从右向左不断扫描，找小于或者等于锚点元素的元素
		for (i<tail && args[i]<= pivot){
			//找到与锚点元素相等的元素将其交换到p所指示的位置
			if args[i] == pivot {
				Swap(i,p,args)
				p++
			}
			i++
		}
		//工作指针j从左向右不断扫描，找大于或者等于锚点元素的元素
		for (j >= head && args[j] >= pivot) {
			//找到与锚点元素相等的元素将其交换到q所指示的位置
			if args[j] == pivot {
				Swap(j,q,args)
				q--
			}
			j--
		}
		//如果两个工作指针i j相遇则一趟遍历结束
		if (i>=j) {
			break
		}
		//将左边大于pivot的元素与右边小于pivot元素进行交换
		Swap(i,j,args)
		i++
		j++
	}
	
	//因为工作指针i指向的是当前需要处理元素的下一个元素
	//故而需要退回到当前元素的实际位置，然后将等于pivot元素交换到序列中间
	i--
	p--
	for p >=head {
		Swap(i,p,args)
		i--
		p--
	}
	//因为工作指针j指向的是当前需要处理元素的上一个元素
    //故而需要退回到当前元素的实际位置，然后将等于pivot元素交换到序列中间
	j++
	q++
	for q <= tail {
		Swap(j,q,args)
		j++
		q++
	}
	//递归遍历左右子序列
	QuickSortThree(head,i,args)
	QuickSortThree(j,tail,args)
}

func Swap(i,j int,arr []int){
	arr[i],arr[j] = arr[j],arr[i]
}

```

#### 双基准排序

&emsp;&emsp;1、对于长度小于17时选择归并排序

&emsp;&emsp;2、两个Pivot，数组开头与结尾

&emsp;&emsp;3、使得`Pivot1<Pivot2` 

&emsp;&emsp;4、排成这样：`****Pivot1*****Pivot2****` 单向扫描

&emsp;&emsp;5、递归

