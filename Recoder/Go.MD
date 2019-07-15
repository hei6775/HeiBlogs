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