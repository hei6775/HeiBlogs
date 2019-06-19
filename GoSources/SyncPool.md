# Golang 的 Sync.Pool 源码探究

如有错误，敬请指正。

## 设计 sync.pool 的目的

这个的设计目的是用来保存和复用临时对象，以减少内存分配，降低 CG 压力。
sync.Pool 是可伸缩的，并发安全的。其大小仅受限于内存的大小。

## sync.pool 的结构

池是一组可以单独保存和检索的临时对象。可以随时自动删除存储在池中的任何项目，而无需通知。如果池在发生这种情况时保留唯一引用，则可以取消分配该项。
池是并发安全的，池的目的是缓存已分配但未使用的项目以供以后重用，从而减轻对 GC 的压力。
也就是说，它可以轻松构建高效，线程安全的空闲列表。但是，它不适用于所有空闲列表。

池的适当使用是管理一组默认共享的临时项，并且可能由包的并发独立客户端重用。池提供了一种在许多客户端上分摊分配开销的方法。

良好使用 Pool 的一个例子是 fmt 包，它维护一个动态大小的临时输出缓冲区存储。缓冲区在负载下（当许多 goroutines 正在积极打印时）进行扩展，并在静止时收缩。

另一方面，作为短期对象的一部分维护的空闲列表不适合用于池，因为在该场景中开销不能很好地摊销。使这些对象实现自己的空闲列表更有效。

首次使用后不得复制池。

```golang
type Pool struct {
	noCopy noCopy

	local     unsafe.Pointer // 固定大小的针对每个P的池，实际类型是[P]poolLocal
	localSize uintptr        // local数组的大小

    // New是一个指定的函数去生成某个值，否则就会返回nil
	// It may not be changed concurrently with calls to Get.
	New func() interface{}
}

// Local per-P Pool appendix.
type poolLocalInternal struct {
	private interface{}   //只可以被各自的P使用，无法共享
	shared  []interface{} //可以被任何P使用
	Mutex                 // Protects shared.
}

type poolLocal struct {
	poolLocalInternal

	// Prevents false sharing on widespread platforms with
	// 128 mod (cache line size) = 0 .
	pad [128 - unsafe.Sizeof(poolLocalInternal{})%128]byte
}
```

为了使得在多个`goroutine`中高效的使用`goroutine`，
`sync.Pool`为每个 P(对应 CPU)都分配一个本地池，当执
行`Get`或者`Put`操作的时候，会先将`goroutine`和某个 P 的子池
关联，再对该子池进行操作。 每个 P 的子池分为私有对象和
共享列表对象，私有对象只能被特定的 P 访问，共享列表对象
可以被任何 P 访问。因为同一时刻一个 P 只能执行
一个`goroutine`，所以无需加锁，但是对共享列表
对象进行操作时，因为可能有多个`goroutine`同时操
作，所以需要加锁。

值得注意的是`poolLocal`结构体中
有个`pad`成员，目的是为了防止`false sharing`。`cache`使用
中常见的一个问题是`false sharing`。当不同的线程同时读
写同一`cache line`上不同数据时就可能发
生`false sharing`。`false sharing`会导致多核处理器上严重
的系统性能下降。具体的可以参考伪共享[(`False Sharing`)](http://ifeve.com/falsesharing/)。

## sync.pool 的主要方法

`sync.pool`的主要方法为`Get`和`Put`方法，先看`Get`的源码。

1.尝试从本地 P 对应的那个本地池中获取一个对象值, 并从本地池
冲删除该值。

2.如果获取失败，那么从共享池中获取, 并从共享队列中删除该值。

3.如果获取失败，那么从其他 P 的共享池中偷一个过来，并删除共享池
中的该值(p.getSlow())。

4.如果仍然失败，那么直接通过 New()分配一个返回值，注意这个分配
的值不会被放入池中。New()返回用户注册的 New 函数的值，如果用户未
注册 New，那么返回 nil。

```golang
func (p *Pool) Get() interface{} {
	if race.Enabled {
		race.Disable()
	}
	l := p.pin()//获取当前线程的poolLocal对象，也就是p.local[pid]
	x := l.private
	l.private = nil
	runtime_procUnpin()
	//否则从shared的最后面开始取
	if x == nil {
		l.Lock()
		last := len(l.shared) - 1
		if last >= 0 {
			x = l.shared[last]
			l.shared = l.shared[:last]
		}
		l.Unlock()
		if x == nil {
			x = p.getSlow()//当本线程的缓存对象已经没有，去其他线程缓存列表中取
		}
	}
	if race.Enabled {
		race.Enable()
		if x != nil {
			race.Acquire(poolRaceAddr(x))
		}
	}
	//如果private和shared都取不到，那么就New一个
	if x == nil && p.New != nil {
		x = p.New()
	}
	//否则返回nil
	return x
}
```

再看`Put`方法：

1.如果放入的值为空，直接 return.

2.检查当前 goroutine 的是否设置对象池私有值，如果没有则将 x 赋
值给其私有成员，并将 x 设置为 nil。

3.如果当前 goroutine 私有值已经被设置，那么将该值追加到共享列表。

```golang
// Put adds x to the pool.
func (p *Pool) Put(x interface{}) {
	if x == nil {
		return
	}
	if race.Enabled {
		if fastrand()%4 == 0 {
			// Randomly drop x on floor.
			return
		}
		race.ReleaseMerge(poolRaceAddr(x))
		race.Disable()
	}
	l := p.pin() //获取当前线程的poolLocal对象，也就是p.local[pid]。
	if l.private == nil {
		l.private = x
		x = nil
	}
	runtime_procUnpin()
	if x != nil {
		l.Lock()
		l.shared = append(l.shared, x)
		l.Unlock()
	}
	if race.Enabled {
		race.Enable()
	}
}
```

最后我们来看一下 init 函数。

```
func init() {
    runtime_registerPoolCleanup(poolCleanup)
}
```

可以看到在 init 的时候注册了一个`PoolCleanup`函数，他会
清除掉`sync.Pool`中的所有的缓存的对象，这个注册函数会
在每次 GC 开始的时候运行，所以`sync.Pool中`的值只在两次 GC 中
间的时段有效。

## sync.pool 的小陷阱

### sync.pool 每次 Put,Get 相同的对象，都是使用相同的内存地址吗？

答：这个问题可能需要分情况看，需要看 `New` 函数 `return` 时返回的是什么。
如果返回的是字符串，那么不好意思，看`Get`的源码`x := l.private`，由于 golang
中都是值传递，所以这个时候返回的`x`实际上肯定不是之前的内存地址了。但是如
果`return`的是一个切片，因为切片实际上是一个引用类型，变量`x`这个时候就是
一个类似指针变量的东西，它的地址当然不是之前的地址，但是它存储的 value 却指向的
是之前的地址，详细看代码输出：

```golang
package main

import (
	"sync"
	"fmt"
)

func main()  {
	// 建立对象
	var pipeString = &sync.Pool{New:func()interface{}{return "Hello, BeiJing"}}
	var pipeSlices = &sync.Pool{New:func()interface{}{
		slices := make([]byte, 1)
		return slices
	}}
	// 准备放入的字符串
	valString := "Hello,World!"
	valSlices := make([]byte,1)
	valSlices[0]=5
	//a := make([]byte,1)
	//b := a
	//fmt.Printf("value: %v address: %p %p \n",a,&a,a)
	//fmt.Printf("value: %v address: %p %p \n",b,&b,b)
	fmt.Printf("String放入前 value: %v 变量地址: %p 真实地址: %p \n",valString,&valString,valString)
	fmt.Printf("Slices放入前 value: %v 变量地址: %p 真实地址: %p \n",valSlices,&valSlices,valSlices)
	// 放入
	pipeString.Put(valString)
	pipeSlices.Put(valSlices)

	// 第一次取出
	firstString := pipeString.Get().(string)
	firstSlices := pipeSlices.Get().([]byte)
	fmt.Printf("String第一次取出 value: %v 变量地址: %p 真实地址: %p \n",firstString,&firstString,firstString)
	fmt.Printf("Slices第一次取出 value: %v 变量地址: %p 真实地址: %p \n",firstSlices,&firstSlices,firstSlices)
	// 再取就没有了,会自动调用NEW
	secondString := pipeString.Get().(string)
	secondSlices := pipeSlices.Get().([]byte)
	fmt.Printf("String第二次取出 value: %v 变量地址: %p 真实地址: %p \n",secondString,&secondString,secondString)
	fmt.Printf("Slices第二次取出 value: %v 变量地址: %p 真实地址: %p \n",secondSlices,&secondSlices,secondSlices)
	//再次放入
	pipeString.Put(firstString)
	pipeSlices.Put(firstSlices)
	//第三次取出
	thirdString := pipeString.Get().(string)
	thirdSlices := pipeSlices.Get().([]byte)
	fmt.Printf("String第三次取出 value: %v 变量地址: %p 真实地址: %p \n",thirdString,&thirdString,thirdString)
	fmt.Printf("Slices第三次取出 value: %v 变量地址: %p 真实地址: %p \n",thirdSlices,&thirdSlices,thirdSlices)
}

```

可以看到输出结果：

```
String放入前 value: Hello,World! 变量地址: 0xc00003e1c0 真实地址: %!p(string=Hello,World!)
Slices放入前 value: [5] 变量地址: 0xc000048460 真实地址: 0xc000052080
String第一次取出 value: Hello,World! 变量地址: 0xc00003e210 真实地址: %!p(string=Hello,World!)
Slices第一次取出 value: [5] 变量地址: 0xc000048520 真实地址: 0xc000052080
String第二次取出 value: Hello, BeiJing 变量地址: 0xc00003e240 真实地址: %!p(string=Hello, BeiJing)
Slices第二次取出 value: [0] 变量地址: 0xc000048580 真实地址: 0xc00005208c
String第三次取出 value: Hello,World! 变量地址: 0xc00003e280 真实地址: %!p(string=Hello,World!)
Slices第三次取出 value: [5] 变量地址: 0xc000048620 真实地址: 0xc000052080
```

每次获得的 string 类型的变量的内存地址都是不一样的，但是对于切片而言放入一个变量，取出则个变量时它们指向的内存地址都是相同的。

### sync.pool 在 GC 后取出的对象还是同一个地址吗？

```golang
package main

import (
	"sync"
	"fmt"
	"runtime"
)

func main()  {
	// 建立对象
	var pipeSlices = &sync.Pool{New:func()interface{}{
		slices := make([]byte, 1)
		return slices
	}}
	// 准备放入的字符串

	valSlices := make([]byte,1)
	valSlices[0]=5

	fmt.Printf("Slices放入前 value: %v 变量地址: %p 真实地址: %p \n",valSlices,&valSlices,valSlices)
	// 放入

	pipeSlices.Put(valSlices)

	// 第一次取出

	firstSlices := pipeSlices.Get().([]byte)
	fmt.Printf("Slices第一次取出 value: %v 变量地址: %p 真实地址: %p \n",firstSlices,&firstSlices,firstSlices)
	// 再取就没有了,会自动调用NEW
	secondSlices := pipeSlices.Get().([]byte)
	fmt.Printf("Slices第二次取出 value: %v 变量地址: %p 真实地址: %p \n",secondSlices,&secondSlices,secondSlices)
	//再次放入
	pipeSlices.Put(firstSlices)
	runtime.GC()
	//第三次取出
	thirdSlices := pipeSlices.Get().([]byte)
	fmt.Printf("Slices第三次取出 value: %v 变量地址: %p 真实地址: %p \n",thirdSlices,&thirdSlices,thirdSlices)
}

```

输出如下：

```golang
Slices放入前 value: [5] 变量地址: 0xc00004a440 真实地址: 0xc000054080
Slices第一次取出 value: [5] 变量地址: 0xc00004a4e0 真实地址: 0xc000054080
Slices第二次取出 value: [0] 变量地址: 0xc00004a540 真实地址: 0xc00005408c
Slices第三次取出 value: [0] 变量地址: 0xc00004a440 真实地址: 0xc000054000
```

可以看到在 GC 后取出的对象已经不是第一次的对象了，它的初始值变了，且指向的内存地址也发生了变化。

### sync.pool 适用在什么场合

答：由 `golang` 的值传递的特性和 `sync.pool`的 `GC`的特性，我们不难发现，
`sync.pool`不适合那种需要长久保持状态的池，
比如说连接池，因为如果每次`GC`都要重新初始化，
那么对于连接的开销而言太大了。而且对于 return 的不是引用类型或者说
不是指针类型的值，`sync.pool`的意义并不大，它依然重新分配内存，
并没有起到一个重复利用的作用。
