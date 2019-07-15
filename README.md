# Hei。 Blogs

&emsp;&emsp;Principe University Algorithm Course And Blogs

## 简介

&emsp;&emsp;数据结构、算法的学习之路，后续发展成随笔记录

## 目录说明

##### [Golang文章翻译](https://github.com/hei6775/HeiBlogs/tree/master/GoVersion)
> 英文文章翻译以及不错的文章记录

##### [beego源码分析](https://github.com/hei6775/HeiBlogs/tree/master/Beego)
> beego 小部分源码，主要为 log 模块和 tool 模块

##### [A 星寻路算法](https://github.com/hei6775/HeiBlogs/tree/master/Astar)
> 轻量简易版的A*算法golang实现

##### [数据库部分](https://github.com/hei6775/HeiBlogs/tree/master/DB)
> mysql mongodb redis

##### [网络部分](https://github.com/hei6775/HeiBlogs/tree/master/Protocol)

##### [ZooKeeper](https://github.com/hei6775/HeiBlogs/tree/master/Zk)

##### [个人随笔](https://github.com/hei6775/HeiBlogs/tree/master/Recoder)

##### [源码分析](https://github.com/hei6775/HeiBlogs/tree/master/GoSources)

## 记录

&emsp;&emsp;golang 的调度系统简介：

因为`golang`想要实现高并发所以采取的 M:N 的内核线程与用户线程的印射，即一个用户线程可以对应多个内核线程，
一个内核线程也可以对应多个用户线程，而且为了便于管理 golang 的 GC，需要 golang 自己实现 goroutine 的调度。
golang 的调度系统主要基于 M-P-G 的结构，M 是内核线程，P 是上下文管理器，G 是 goroutine 也就是需要被调度的任务，
G 需要绑定 P 才能被 M 执行，当使用一个 go 关键字调起一个 goroutine 时，底层通过调用 newproc 生成一个新的调度任务，放入 GlobalRunningQueue 或者 LocalRunningQueue 中等待调度。当某个 P 阻塞的时候，该 P 下的 G 就会被放到其他的 P 中执行，如果
P 中的 G 执行完成，那么这个 P 会从 GlobalRunningQueue 中获取 G 或者从其他的 P 中偷一半的 G 来执行，当然 P 也会定期检测 GlobalRunningQueue
，防止 G 不被调用，P 不会饿死。

&emsp;&emsp;golang 中赋值都是复制，如果赋值了一个指针，那我们就复制了一个指针副本。
如果赋值了一个结构体，那我们就复制了一个结构体副本。往函数里传参也是同样的情况。

&emsp;&emsp;但是有一点点不同的是，函数传参：

1、指针传递，传递的是指针的地址，但是形参的地址是另外一个，存储的是实参的地址，修改形参会直接修改实参

2、数组传递，传递的是数组的“值拷贝”，对形参进行操作并不会影响到实参

3、数组名传递，和 2 相同

4、Slice 传递，地址拷贝，传递的是底层数组的内存地址，修改形参实际上会修改实参

5、函数传递

&emsp;&emsp;Golang 反射三大定律

1、反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”。

2、反射第二定律：反射可以将“反射类型对象”转换为“接口类型变量”。

3、反射第三定律：如果要修改“反射类型对象”，其值必须是“可写的”（settable）

Golang 中 byte、string、rune 的关系

&emsp;&emsp;首先我们要知道 golang 的默认编码是 utf-8，中文 unicode 下是占两个字节，在 utf-8 下占三个字节，而在 string 底层使用 byte 数组
存，并且不可改变。直接对中文字符串`len()`操作得出的不一定是真实的长度，这是因为 byte 等同于 int8，常用来处理 ascii 字符，而 rune 等于 int32，常用来处理 unicode 和 utf-8 字符。
想要获取中文的话需要使用 rune 转换

&emsp;&emsp;表达式 go f(x, y, z)会启动一个新的 goroutine 运行函数 f(x, y, z)。函数 f，变量 x、y、z 的值是在原 goroutine 计算的，只有函数 f 的执行是在新的 goroutine 中的

&emsp;&emsp;defer 关键字的实现跟 go 关键字很类似，不同的是它调用的是 runtime.deferproc 而不是 runtime.newproc

&emsp;&emsp;runtime.newproc 函数接受的参数分别是：参数大小，新的 goroutine 是要运行的函数，函数的 n 个参数。首先，让我们看一下如果是 C 代码新建一条线程的实
现会是什么样子的。大概会先建一个结构体，结构体里存 f、x、y 和 z 的值。然后写一个 help 函数，将这个结构体指针作为输入，函数体内调用 f(x, y, z)。
接下来，先填充结构体，然后调用 newThread(help, structptr)。其中 help 是刚刚那个函数，它会调用 f(x, y, z)。help 函数将作为所有新建线程的入口函数。

&emsp;&emsp;逃逸分析，从栈中逃逸到堆中

&emsp;&emsp;连续栈技术

&emsp;&emsp;`defer`是在`return`之后，函数真正返回之前执行，且`defer`的特
性有两个，先进后出和定义时参数就确定了。例子：
```golang
package main

import "fmt"

func main(){
	a,b := 0,1
	defer add("m",a,add("p",a,b))
	fmt.Println("stop start")
	a = 2
	defer add("n",a,add("k",a,b))
	return
}

func add(s string,a,b int)int{
	fmt.Println(s,a+b)
	return a+b
}
//outputs:
//  p 1
//  stop start
//  k 3
//  n 5
//  m 1
```

## Golang 的 HTTP

1.目前我知道 http.Request.Body 可以获取上传之后的文件具体内容，并且它是一个 io.ReadCloser 类型

2.再问一下 http.Request.Body 是服务器这边全部接收到了内容之后才能读取到吗？

3.http.Request.ContentLength 是如何得到的，是服务器这边全部接收到了内容之后算的？还是客户端那边带过来的

答：

如果文件很大,分片上传即可,在客户端分片,最后在服务器端组合成原文件,可以看看某些 CDN api 的实现,比如阿里云 CDN 等都有分片上传的实现.因为大文件直接上传时,碰到网络中断后就要重新开始, 那就要吐血了.

http.Request.ContentLength 是客户端实现的,有兴趣,就翻翻源码.

http.Request.Body 无需全部接收到了内容之后才能读取,但读取过程是阻塞的.

上传文件即是将文件编码为 multipart/form-data 后再放到 http.body 里进行上传,当 http.body 过大时,底层的 tcp 会分段进行传输.因此上传一个大文件,用 for 和小的 buffer 进行循环读取即可验证该过程.而且我说的"读取过程是阻塞的"是指客户端在传输过程中突然停止,服务端没数据可读时也会阻塞住.

## Golang 的内存分配

Golang 运行时的内存分配算法主要源自 Google 为 C 语言开发的`TCMalloc`算法，全称`Thread-CachingMalloc`。
核心思想就是把内存分为多级管理，从而降低锁的粒度。
它将可用的堆内存采用二级分配的方式进行管理：每个线程都会自行维护一个独立的内存池，进行内存分配时优先从该内存池中分配，
当内存池不足时才会向全局内存池申请，以避免不同线程对全局内存池的频繁竞争。

Go 在程序启动的时候，会先向操作系统申请一块内存（注意这时还只是一段虚拟的地址空间，并不会真正地分配内存），切成小块后自己进行管理。

- spans 区域
- bitmap 区域
- arena 区域
- mspan
- mcache
- mcentral
- mheap

&emsp;&emsp;内存池，垃圾回收。

> #### 内存池

- 动态分配内存大小

&emsp;&emsp;

- 每条线程都会有自己的本地的内存，然后还有一个全局的分配链

&emsp;&emsp;

- 两级内存管理结构，MHeap 和 MCache

&emsp;&emsp;MHeap 用于分配大对象，每次分配都是若干连续的页，也就是若干个 4KB 的大小。使用的数据结构是 MHeap 和 MSpan，用 BestFit 算法做分配，
用位示图做回收。

&emsp;&emsp;MCache 用于管理的基本单位是不同类型的固定大小的对象，更像是一个对象池而不是内存池，用引用计数做回收。

> #### 垃圾回收

## Golang 小陷阱

golang 解析 json 时把所有的 int,float,double 等数字，向 interface{}解析时都当成 float64（当然被双信号包围的数字除外，任何被双引号包围的，都是字符串。

```golang
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var v map[string]interface{}
    jsonstr := `{"id":13,"name":"胖胖","weight":216.5,"dd":"123"}`
    json.Unmarshal([]byte(jsonstr), &v)
    for k, v1 := range v {
        fmt.Print(k, " = ")
        switch v1.(type) {
        case int:
            fmt.Println(v1, "is an int value.")
        case string:
            fmt.Println(v1, "is a string value.")
        case int64:
            fmt.Println(v1, "is an int64 value.")
        case float64:
            fmt.Println(v1, "is an float64 value.")
        default:
            fmt.Println(v1, "is an unknown type.")
        }
    }
}
//output:
// weight = 216.5 is an float64 value.
// dd = 123 is a string value.
// id = 13 is an float64 value.
// name = 胖胖 is a string value.
// Process exiting with code: 0

```

### Golang死锁的原因

形成死锁的四个必要条件：

* 互斥条件：一个资源每次只能被一个进程所使用
* 请求与保持条件：一个进程因请求资源而阻塞时，对已获得的资源保持不放
* 不剥夺条件：进程已获得的资源，在未使用完之前，不能强行剥夺
* 循环等待条件：若干个进程之间形成一种头尾相接的循环等待资源关系

## Others

> #### 树

- 前序遍历： 根结点 ---> 左子树 ---> 右子树

```go
//1、创建一个栈对象，将根节点入栈；

//2、当栈为非空时，将栈顶结点弹出栈并访问该结点；

//3、对当前访问的非空左孩子结点相继依次访问（不需要入栈），并将访问结点的非空右孩子结点入栈

//4、重复执行步骤②和步骤③，直到栈为空为止
func PreStackTraverse(t *TreeNode){//先根遍历，非递归
	if t != nil {
		S := CreateStack()
		S.Push(t)
		for !S.IsEmpty() {
			T,_ := S.Pop()
			fmt.Printf("%d ",T.Value)
			for T != nil {
				if T.Left != nil {
					fmt.Printf("%d ",T.Left.Value)
				}
				if T.Right != nil {
					S.Push(T.Right)
				}
				T = T.Left
			}
		}
	}
	fmt.Println()
}
```

- 中序遍历：左子树---> 根结点 ---> 右子树

- 后序遍历：左子树 ---> 右子树 ---> 根结点

> #### 进制转换

十进制转其他进制

- 原则 1：整数部分与小数部分分别转换；

- 原则 2：整数部分采用除基数(转换为 2 进制则每次除 2，转换为 8 进制每次除 8，以此类推)取余法，直到商为 0，而余数作为转换的结果，第一次除后的余数为最低位，最后一次的余数为最高位；

- 原则 3：小数部分采用乘基数(转换为 2 进制则每次乘 2，转换为 8 进制每次乘 8，以此类推)取整法，直至乘积为整数或达到控制精度。

```python
725.625D=1011010101.101B
725/2 余数1 商362  0.625*2=1.25 整数部分1 小数部分0.25
362/2 余数0 商181  0.25*2=0.5 整数部分0 小数部分0.5
....               ......
2/2 余数0 商1       0.5*2=1.0 整数部分1 小数部分0
1/2 余数1 商0
```

其他进制转十进制

- 按权展开法，即把各数位乘权的 i 次方后相加

```python
#二进制转十进制
01011010.01B=0×2^7+1×2^6+0×2^5+1×2^4+1×2^3+0×2^2  +1×2^1+0×2^0+0×2^-1+1×2^-2 = 90.25
```

### 跨域

|                             URL                              | Description                       |             是否允许通信              |
| :----------------------------------------------------------: | --------------------------------- | :-----------------------------------: |
|      `http://www.d.com/d.js`<br>`http://www.d.com/w.js`      | 同一域名下                        |                 允许                  |
| `http://www.d.com/lab/a.js`<br>`http://www.d.com//src/b.js`  | 同一域名下不同文件夹              |                 允许                  |
| `http://www.d.com:3333/a.js`<br>`http://www.d.com:4444/b.js` | 同一域名不同端口                  |                不允许                 |
|     `http://www.d.com/a.js`<br>`http://46.33.22.44/b.js`     | 域名和域名对应 IP                 |                不允许                 |
|    `http://www.d.com/a.js`<br>`http://script.d.com/b.js`     | 主域相同，子域不同                |                不允许                 |
|        `http://www.d.com/a.js`<br>`http://d.com/w.js`        | 同一域名,不同二级域名<br>（同上） | 不允许(cookie 这种情况下也不允许访问) |
|      `http://www.d.com/d.js`<br>`http://www.v.com/w.js`      | 不同域名                          |                不允许                 |

js

```javascript
res.header('Access-Control-Allow-Origin', '*')
res.header('Access-Control-Allow-Headers', 'X-Requested-With,Content-Type')
res.header('Access-Control-Allow-Methods', 'PUT,POST,GET,DELETE,OPTIONS')
```

golang

```go
	//	Origin := r.Header.Get("Origin")
	//	if Origin != "" {
	//		w.Header().Add("Access-Control-Allow-Origin", "*")
	//		w.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	//		w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type")
	//		w.Header().Add("Access-Control-Allow-Credentials", "true")
	//	}
```

## LeetCode

### 189.Rotate Array

如果 k>n

```golang
//author:abbycoding
//comments:genius
//such as
// [1,2,3,4,5]  k=2
// output [4,5,1,2,3]
// n=5, k = k % n = 2,
//reverse(nums,0,4)  [1,2,3,4,5] ==> [5,4,3,2,1]
//reverse(nums,0,1)  [5,4,3,2,1] ==> [4,5,3,2,1]
//reverse(nums,2,4)  [4,5,3,2,1] ==> [4,5,1,2,3]

func rotate(nums []int, k int) {
    n := len(nums) //数组长度
	k %= n //如果k>n的情况，则取k/n的余数
    reverse(nums, 0, n - 1)
    reverse(nums, 0, k - 1)
    reverse(nums, k, n - 1)
}

//反转 交换
func reverse(nums []int, start int, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
```

### 193. Valid Phone Numbers

file.txt 是每行都是一串的数字，
从 file.txt 选出合法的电话号码，电话号码的格式为：

(xxx) xxx-xxxx or xxx-xxx-xxxx.

```shell
grep -P '^(\d{3}-|\(\d{3}\) )\d{3}-\d{4}$' file.txt
#-r参数开启扩展正则模式，-n只打印被sed处理的行
sed -n -r '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/p' file.txt

awk '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/' file.txt
```

### 198. House Robber

在这道题了解到动态规划（`dynamic-programming`）

```golang
func rob(nums []int) int {
    prevMax := 0
    currMax := 0

    for i:=0; i < len(nums); i++ {
        temp:=currMax
        if prevMax + nums[i] > currMax {
            currMax = prevMax + nums[i]
        }
        prevMax = temp
    }
    return currMax
}
// input [1,2,3,1]
// output 4

// temp = currMax = 0
// 0+1 = 1 > 0
// currMax = 1
// prevMax = 0

// temp = currMax = 1
// 0 + 2 > 1
// currMax = 2
// prevMax = 1

// temp = currMax = 2
// 1 + 3 = 4 > 2
// currMax = 4
// prevMax = 2

// temp = currMax = 4
// 2 + 1 > 4
// prevMax = 4
// return 4

```

## 等比数列 等差数列

![等差数列](Asset/等差数列求和公式.png)

![等比数列](Asset/等比数列求和公式.png)

## Links

[网络计算机基础篇](https://hit-alibaba.github.io/interview/basic/network/HTTP.html)

[深入理解 Go Slice](https://mp.weixin.qq.com/s?__biz=MjM5OTcxMzE0MQ==&mid=2653371806&idx=1&sn=37cdffa7b5ec5bfb901455bb3997a040&chksm=bce4dd848b9354929adecd85d1b502f381a60ae36ecdd69c320942a367fbeb417a48a87c60c8&scene=21##)

[图解 Go 语言内存分配](https://mp.weixin.qq.com/s/Hm8egXrdFr5c4-v--VFOtg)

[Go 语言实战笔记（二十七）| Go unsafe Pointer](https://www.flysnow.org/2017/07/06/go-in-action-unsafe-pointer.html)

[Go 语言经典库使用分析（八）| 变量数据结构调试利器 go-spew](https://www.flysnow.org/2019/02/03/golang-classic-libs-go-spew.html)

[有点不安全却又一亮的 Go unsafe.Pointer](https://blog.csdn.net/RA681t58CJxsgCkJ31/article/details/85241470)

[Go 语言性能优化- For Range 性能研究](https://www.flysnow.org/2018/10/20/golang-for-range-slice-map.html)

[深入 Golang Runtime 之 Golang GC 的过去,当前与未来](https://www.jianshu.com/p/bfc3c65c05d1)

[Golang 使用指针来操作结构体就一定更高效吗？](https://medium.com/@blanchon.vincent/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963)
