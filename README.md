# Algorithm

&emsp;&emsp;Principe University Algorithm Course

## 简介

&emsp;&emsp;数据结构、算法的学习之路

## 说明

|   Path    | Description                                   |
| :-------: | --------------------------------------------- |
|   Astar   | A 星寻路算法实现.                             |
|   Beego   | beego 小部分源码，主要为 log 模块和 tool 模块 |
|    DB     | 预留数据库，已含 Leaf 的 mongo 数据库模块     |
|   Tools   | 个人常用工具函数，附带单元测试                |
|    Go     | Leaf 框架封装的 goroutine.                    |
| GoVersion | Go Blogs                                      |
|   Http    | Http 相关练习                                 |
| Lecture01 | 普林斯顿大学算法课程一                        |
| Lecture02 | 普林斯顿大学算法课程二.                       |
| Lecture03 | 普林斯顿大学算法课程三                        |
| MyLecTest | 课程练习使用                                  |
|  Recoder  | 记录文件                                      |
|    Rf     | Leaf 框架的 ReadFile 模块                     |
|  Socket   | socket 练习                                   |
|   Until   | Leaf 框架的部分模块，以及常用工具函数         |
|   Tree    | 树数据结构练习                                |
|   Tools   | 个人使用工具                                  |
|    Ws     | websocket 练习                                |
|    Zk     | zookeeper 封装                                |

## 记录

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

&emsp;&emsp;defer 是在 return 之前执行的

## Golang 的内存管理

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

如果k>n

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
