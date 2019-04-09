# The Go Memory Model

## Introduce

&emsp;&emsp;本文介绍 GoLang 的内存模型，原文是 GoLang 官方 2014 年的文章：[The Go Memory Model](https://golang.org/ref/mem)
。

## 正文

### Version of May 31, 2014

Introduction <br />
Advice <br />
Happens Before <br />
Synchronization <br />
&emsp;&emsp;Initialization <br />
&emsp;&emsp;Goroutine creation <br />
&emsp;&emsp;Goroutine destruction <br />
&emsp;&emsp;Channel communication <br />
&emsp;&emsp;Locks <br />
&emsp;&emsp;Once <br />
Incorrect synchronization <br />

### Introduce

&emsp;&emsp;GoLang 的内存模型指定了一个条件，在这个条件下，可以保证在一个`goroutine`中读取变量可以被另外一个写入相同变量
的`goroutine`观察到。

### Advice

&emsp;&emsp;修改同时被多个`goroutine`访问的数据的程序必须序列化这样的访问。

&emsp;&emsp;为了序列化访问，通过`channel`操作和其它同步操作保护数据，这些操作都在`sync`和`sync/atomic`包中可以找到。

&emsp;&emsp;如果您必须阅读本文档的其余部分以了解程序的行为，那么您就太聪明了。

&emsp;&emsp;别聪明。

### Happens Before

&emsp;&emsp;在一个`goroutine`中，读和写必须表现得好像它们是按程序的顺序被执行的。换一句话说，编译器和处理器
会重新对一个`goroutine`中的读和写的执行顺序进行排序但是必然不会影响语言规范定义的`goroutine`内的行为。因为
编译器的优化功能，进行的重新排序，一个`goroutine`中观察到的执行的顺序可能完全不同于另一个`goroutine`感知到
的顺序。比如说一个`goroutine`中的执行`a=1;b=2`；另一个`goroutine`b 的赋值在 a 的赋值操作之前。

&emsp;&emsp;为了指定读和写的需求，我们定义了`Happens Before`，这是一个在 Go 程序中的局部的内存操作执行顺序。
如果事件`e1`发生在`e2`之前，那我们则说`e2`发生在`e1`之后。当然，如果`e1`没发生在`e2`之前，也没有发生在`e2`
之后，那我们则说`e1`和`e2`同时执行。

&emsp;&emsp;在一个`goroutine`，`Happens Before`是一个程序执行表现的顺序。

&emsp;&emsp;如果满足下面的要求，则对允许一个变量`v`的读操作`r`感知到一个写操作`w`：

&emsp;&emsp;&emsp;&emsp;1、`r`没有发生在`w`之前；

&emsp;&emsp;&emsp;&emsp;2、没有其它对变量`v`的写操作`w'`发生在`w`操作之后,`r`操作之前；

&emsp;&emsp;为了保证变量`v`的读取`r`观察到`v`的特定写入`w`，确保`w`是允许观察的唯一写入`r`。也就是说，
如果以下两个都成立，则`r`保证观察到`w`：

&emsp;&emsp;&emsp;&emsp;1、`w`发生在`r`之前；

&emsp;&emsp;&emsp;&emsp;2、对共享变量`v`的任何其他写入要么发生在`w`之前，要么发生在`r`之后；

&emsp;&emsp;这对条件比第一对更强;它要求没有其他写入操作与 w 或 r 同时发生。

&emsp;&emsp;在单个 goroutine 中，没有并发性，所以这两个定义是等价的：读取操作`r`可以观察到最近写入操作`w`对
变量`v`写入的值。当多个`goroutine`访问共享变量`v`时，它们必须使用同步事件来确保`Happens Before`条件读取
观察到所需要的写入。

### Synchronization

#### Initialization

&emsp;&emsp;程序初始化运行在一个`goroutine`中，但是这个`goroutine`会创建其它`goroutine`，这个`goroutine`会并发执行。

&emsp;&emsp;如果 A 包引入了 B 包，那么 B 包的`init`函数会发生在 A 包的`init`函数之前。

&emsp;&emsp;`main.main`函数会发生在所有的`init`函数结束之后。

#### Goroutine creation

&emsp;&emsp;启动新 goroutine 的 go 语句发生在 goroutine 的执行开始之前。

比如说这段程序：

```golang
var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}
```

调用 hello 将在未来的某个时刻打印“hello，world”（也许在 hello 返回之后）。

#### Goroutine destruction

&emsp;&emsp;goroutine 的退出不保证在程序中的任何事件之前发生。例如，在此程序中：

```golang
var a string

func hello() {
	go func() { a = "hello" }()
	print(a)
}
```

对 a 的赋值没有跟随任何同步事件，因此不保证任何其他 goroutine 都能观察到它。实际上，一个积极的编译器可能会删除整个 go 语句。

如果必须由另一个 goroutine 观察到 goroutine 的影响，请使用`lock`或`channel`等同步机制来建立相关的顺序。

#### Channel communication

&emsp;&emsp;`channel`通信是两个`goroutine`同步通讯的主要方式。
发送到通道需要一个对应的通道来接受，发送和接受通常是是不同的`goroutine`。

&emsp;&emsp;在通道发送发生在对应的通道接受之前。

这个程序：

```golang
var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
```

这段程序保证打印“hellp, world”,对`a`变量的写入操作发生在往 C 通道发送数据之前，而这些又发生在通道接受信息操作之前，，通道接受数据操作又在`print`操作之前。在上一个例子，使用`close(c)`替换`c<-0`，程序依然会保证打印出“hello, world”。

来自无缓冲通道的接收在该通道上的发送完成之前发生。这个程序（如上所述，但发送和接收语句交换并使用无缓冲通道）：

```golang
var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}
func main() {
	go f()
	c <- 0
	print(a)
}
```

这个程序依然会打印出“print, world”。`a`变量的复制发生在`c`的接受之前，而`c`的接受必须要有`c`的发送，`c`的发送发生在`print`之前。

如果通道是有缓存的（比如说：`c = make(chan int,1)`），那么程序就不能保证一定会打印出“hello, world”了，它可能打印出空字符串，意外等其它情况。

`The kth receive on a channel with capacity C happens before the k+Cth send from that channel completes.`

这句话的意思是：缓存大小为 C 的通道的第 k 次接受发生在它的第 K+C 次发送之前。

此规则将先前的规则概括为缓冲的通道。它允许计数信号量由缓冲通道塑造：通道中的元素数量对应于正在使用的个数，通道的容量对应于最大同时使用个数，发送元素获取信号量，以及接收项目会释放信号量。这是限制并发的常用习惯用法。

这个程序为每个输入开启一个`goroutine`，但是`goroutine`的数量对应于通道的容量，最多只能有三个`goroutine`同时运行。

```golang
var limit = make(chan int, 3)

func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}
```

#### Locks

&emsp;&emsp;`sync`包提供了两种锁，一种是通用锁`sync.Mutex`，另一种是读写锁`sync.RWMutex`。

不管通用锁还是读写锁而言，对于变量`l`且`n<m`，调用 n 次`l.Unlock()`发生在调用 m 次`l.Lock()`返回之前。

这个程序：

```golang
var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
```

保证打印“hello, world”。第一次调用`l.Unlock()`（在`f`中）发生在第二次调用 `l.Lock`（在`main`中）返回之前，这发生在`print`之前。

对于在`sync.RWMutex`变量`l`上对`l.RLock`的任何调用，调用`n`次的`l.RLock`在发生在调用`n`次的`l.Unlock`，并且匹配的`l.RUnlock`在调用`n + 1`到`l`之前发生在调用`n+1`次的`l.Lock`。

#### Once

&emsp;&emsp;`sync`包通过使用`Once`type，来提供一种安全的在多个`goroutine`初始化的机制。多个线程可以执行`once.Do(f)`,但是只有一个线程可以执行`f()`,其它的则阻塞住知道它返回。

`A single call of f() from once.Do(f) happens (returns) before any call of once.Do(f) returns.`

在这段程序中：

```golang
var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

调用`twoprint`会调用`setup`，而`setup`函数会发生在每个`print`操作之前，所以会打印两次`hello,world`。

## Incorrect synchronization

&emsp;&emsp;注意，读`r`可以观察与`r`同时发生的写`w`所写的值。即使发生这种情况，也不意味着在`r`之后发生的读取将观察到在`w`之前发生的写入。

在这个程序中：

```golang
var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}

```

结果可能是`g`打印`2`和`0`。这些因素使得一些常用的语法失效。

双重检查锁定是为了避免同步开销。例如，`twoprint`程序可能被错误地写为:

```golang
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

但是不能保证，在`doprint`中，观察`done`中的写入意味着观察到写入`a`。此版本可以（错误地）打印空字符串而不是“hello，world”。

另一个错误的语法是等待一个值：

```golang
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
```

和上面一样，在`main`中没有保证观察到`done`的写入会意味着`a`的写入,所以这个程序也可能打印出空字符串。更糟糕的是，这里并不能保证`done`一定会被`main`观察到，在这两个线程之间可能没有同步事件，`main`里的循环可能无法保证结束。

这个主题有更微妙的变体，例如这个程序：

```golang
type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}
```

即使`main`中可以通过 g!=nil 来退出循环，但是并不保证它可以观察到`g.msg`的初始化。

`In all these examples, the solution is the same: use explicit synchronization.`

在所有这些示例中，解决方案是相同的：使用显式同步。
