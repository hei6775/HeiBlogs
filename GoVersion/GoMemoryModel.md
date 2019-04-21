# The Go Memory Model

## Introduce

&emsp;&emsp;本文翻译自 GoLang 官方 2014 年的文章：[The Go Memory Model](https://golang.org/ref/mem)，如有翻译不合理之处敬请指正，会及时修改。

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

如果程序中修改数据时有其他`goroutine`同时读取，那么必须将读取串行化。
为了串行化访问，请使用`channel`或其他同步原语，例如`sync`和`sync/atomic`来保护数据。。

&emsp;&emsp;别聪明。

### 先行发生 Happens Before

&emsp;&emsp;在一个`gouroutine`中，读和写一定是按照程序中的顺序执行的。
即编译器和处理器只有在不会改变这个`goroutine`的行为时才可能修改读和写的执行顺序。由于重排，不同的 goroutine 可能会看到不同的执行顺序。例如，一个`goroutine`执行`a = 1;b = 2`;，另一个`goroutine`可能看到 b 在 a 之前更新。

&emsp;&emsp;为了说明读和写的必要条件，我们定义了先行发生（`Happens Before`）--Go 程序中执行内存操作的偏序。如果事件 e1 发生在 e2 前，我们可以说 e2 发生在 e1 后。如果 e1 不发生在 e2 前也不发生在 e2 后，我们就说 e1 和 e2 是并发的。

&emsp;&emsp;在单独的 goroutine 中先行发生`Happens Before`的顺序即是程序中表达的顺序。

&emsp;&emsp;当下面条件满足时，对变量 v 的读操作 r 是被允许看到对 v 的写操作 w 的：：

&emsp;&emsp;&emsp;&emsp;1、r 不先行发生于 w；

&emsp;&emsp;&emsp;&emsp;2、在 w 后 r 前没有对 v 的其他写操作；

&emsp;&emsp;为了保证对变量 v 的读操作 r 看到对 v 的写操作 w,要确保 w 是 r 允许看到的唯一写操作。即当下面条件满足时，r 被保证看到 w：

&emsp;&emsp;&emsp;&emsp;1、w 先行发生于 r

&emsp;&emsp;&emsp;&emsp;2、其他对共享变量 v 的写操作要么在 w 前，要么在 r 后。；

&emsp;&emsp;这一对条件比前面的条件更严格，需要没有其他写操作与 w 或 r 并发发生。

&emsp;&emsp;单独的 goroutine 中没有并发，所以上面两个定义是相同的：读操作 r 看到最近一次的写操作 w 写入 v 的值。当多个 goroutine 访问共享变量 v 时，它们必须使用同步事件来建立先行发生这一条件来保证读操作能看到需要的写操作。 对变量 v 的零值初始化在内存模型中表现的与写操作相同。 对大于一个字的变量的读写操作表现的像以不确定顺序对多个一字大小的变量的操作。

### Synchronization

#### Initialization

&emsp;&emsp;程序的初始化在单独的 goroutine 中进行，但这个 goroutine 可能会创建出并发执行的其他 goroutine。

&emsp;&emsp;如果 A 包引入了 B 包，那么 B 包的`init`函数会发生在 A 包的`init`函数之前。

&emsp;&emsp;`main.main`函数会发生在所有的`init`函数结束之后。

#### Goroutine creation

&emsp;&emsp;启动新`goroutine`的`go`语句发生在`goroutine`的执行开始之前。

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

调用`hello`将在未来的某个时刻打印“hello，world”（也许在 hello 返回之后）。

#### Goroutine destruction

&emsp;&emsp;`goroutine`的退出不保证在程序中的任何事件之前发生。例如，在此程序中：

```golang
var a string

func hello() {
	go func() { a = "hello" }()
	print(a)
}
```

没有用任何同步操作限制对 a 的赋值，所以并不能保证其他`goroutine` 能看到 a 的变化。实际上，一个激进的编译器可能会删掉整个 go 语句。

如果想要在一个`goroutine`中看到另一个`goroutine`的执行效果，请使用锁或者`channel`这种同步机制来建立程序执行的相对顺序。

#### Channel communication

&emsp;&emsp;`channel`通信是`goroutine`同步的主要方法。每一个在特定`channel`的发送操作都会匹配到通常在另一个`goroutine`执行的接收操作。。

&emsp;&emsp;在`channel`的发送操作先行发生于对应的接收操作完成 例如：

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

这个程序能保证打印出"hello, world"。对 a 的写先行发生于在 c 上的发送，先行发生于在 c 上的对应的接收完成，先行发生于`print`。

**对 channel 的关闭先行发生于接收到零值，因为 channel 已经被关闭了。**

在上面的例子中，将 c <- 0 替换为 close(c)还会产生同样的结果。

**无缓冲 channel 的接收先行发生于发送完成**

如下程序（和上面类似，只交换了对 channel 的读写位置并使用了非缓冲 channel）

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

此程序也能保证打印出"hello, world"。对 a 的写先行发生于从 c 接收，先行发生于向 c 发送完成，先行发生于`print`。
如果是带缓冲的 channel（例如`c = make(chan int, 1`)），
程序不保证打印出"hello, world"(可能打印空字符，程序崩溃或其他行为)。

`The kth receive on a channel with capacity C happens before the k+Cth send from that channel completes.`

在容量为 C 的 channel 上的第 k 个接收先行发生于从这个 channel 上的第 k+C 次发送完成。

这条规则将前面的规则推广到了带缓冲的`channel`上。可以通过带缓冲的`channel`来实现计数信号量：`channel`中的元素数量对应着活动的数量，`channel`的容量表示同时活动的最大数量，发送元素获取信号量，接收元素释放信号量，这是限制并发的通常用法。

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

不管通用锁还是读写锁而言，**对任意的 sync.Mutex 或 sync.RWMutex 变量 l 和 n < m，n 次调用 l.Unlock()先行发生于 m 次 l.Lock()返回。**

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

能保证打印出"hello, world"。第一次调用`l.Unlock()`（在`f()`中）先行发生于`main`中的第二次`l.Lock()`返回, 先行发生于`print`。

对于`sync.RWMutex`变量 l，任意的函数调用`l.RLock`满足第 n 次`l.RLock`后发生于第 n 次调用`l.Unlock`，对应的`l.RUnlock`先行发生于第 n+1 次调用`l.Lock`。

#### Once

&emsp;&emsp;`sync`包的`Once`为多个`goroutine`提供了安全的初始化机制。能在多个线程中执行`once.Do(f)`，但只有一个`f()`会执行，其他调用会一直阻塞直到`f()`返回。

通过执行先行发生（指`f()`返回）于其他的返回。

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

调用`twoprint`会打印"hello, world"两次。`setup`只在第一次`doprint`时执行。

## Incorrect synchronization

&emsp;&emsp;注意，读操作 r 可能会看到并发的写操作 w。即使这样也不能表明 r 之后的读能看到 w 之前的写。

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

g 可能先打印出 2 然后是 0。

这个事实证明一些旧的习惯是错误的。

双重检查锁定是为了避免同步的资源消耗。例如`twoprint`程序可能会错误的写成

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

在`doprint`中看到`done`被赋值并不保证能看到对 a 赋值。此程序可能会错误地输出空字符而不是"hello, world"。

另一个错误的习惯是忙等待 例如：

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

和之前程序类似，在`main`中看到`done`被赋值不能保证看到 a 被赋值，所以此程序也可能打印出空字符。更糟糕的是因为两个线程间没有同步事件，在`main`中可能永远不会看到`done`被赋值，所以`main`中的循环不保证能结束。

对程序做一个微小的改变：

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

即使`main`看到了`g != nil`并且退出了循环，也不能保证看到`g.msg`的初始化值。

`In all these examples, the solution is the same: use explicit synchronization.`

在所有这些示例中，解决方案是相同的：明确的使用同步。
