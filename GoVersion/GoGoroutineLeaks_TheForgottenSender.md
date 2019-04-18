# Goroutine Leaks - The Forgotten Sender

<p align="right">Author: Jacob WalkerNovember 12, 2018</p>

原文地址 ：[Goroutine Leaks - The Forgotten Sender](https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html)

## Introduction

并发编程允许开发人员使用多个进程，多线程解决问题，并且通常用于尝试提高性能。
并发并不意味着这些多进程，多线程并行执行; 这意味着这些进程是无序执行而不是顺序执行。
从历史上看，使用由标准库或第三方开发人员提供的库可以促进这种类型的编程。

在 Go 中，`Goroutines`和`channel`等并发功能内置于语言和运行时，以减少或消除对库的需求。
这造成了在 Go 中编写并发程序很容易的错觉。
在决定使用并发时必须谨慎，因为如果没有正确使用它会带来一些独特的副作用或陷阱。
如果你不小心，这些陷阱会产生复杂性和令人讨厌的错误。

我将在这篇文章中讨论的陷阱与`Goroutine`泄漏有关。

##　 Leaking Goroutines

在内存管理方面，Go 为您处理了许多细节。Go 编译器使用**转义分析**确定值在内存中的位置。运行时通过使用**垃圾收集器来**跟踪和管理堆分配。虽然在您的应用程序中创建**内存泄漏**并非不可能，但是大大降低了机会。

一种常见的内存泄漏类型是 Goroutines 泄漏。如果你开始一个你期望最终终止的 Goroutine，但它永远不会，那么它已经泄露。它存在于应用程序的生命周期中，并且无法释放为 Goroutine 分配的任何内存。这是 **“在不知道它将如何停止”的情况下永远不要开始 goroutine”**的建议的原因之一。

要说明基本的 Goroutine 泄漏，请查看以下代码：

###　 Listing 1
[https://play.golang.org/p/dsu3PARM24K](https://play.golang.org/p/dsu3PARM24K)

```
31 // leak is a buggy function. It launches a goroutine that
32 // blocks receiving from a channel. Nothing will ever be
33 // sent on that channel and the channel is never closed so
34 // that goroutine will be blocked forever.
35 func leak() {
36     ch := make(chan int)
37
38     go func() {
39         val := <-ch
40         fmt.Println("We received a value:", val)
41     }()
42 }
```

Listing 1 定义了一个名为`leak`的函数。
该函数在第 36 行创建一个通道，允许 Goroutines 传递整数型数据。
然后在第 38 行创建 Goroutine，它在第 39 行阻塞，等待从通道接收值。
当 Goroutine 等待时，leak 函数返回。
此时，程序的其他任何部分都不能通过通道发送信号。
这使得 Goroutine 在第 39 行被无限期等待。第 40 行`fmt.Println`的调用永远不会发生。

在此示例中，可以在代码审查期间快速识别 Goroutine 泄漏。不幸的是，Goroutine 泄漏的生产代码通常更难以找到。我无法显示 Goroutine 泄漏可能发生的所有可能方式，但是这篇文章将详细说明您可能遇到的一种 Goroutine 泄漏：

## Leak: The Forgotten Sender

**对于此泄漏示例，您将看到无限期阻塞的 Goroutine，等待在通道上发送值。**

我们将要查看的程序根据某些搜索词找到记录，然后打印出来。该程序名为 search：

### Listing 2

[https://play.golang.org/p/o6_eMjxMVFv](https://play.golang.org/p/o6_eMjxMVFv)

```
29 // search simulates a function that finds a record based
30 // on a search term. It takes 200ms to perform this work.
31 func search(term string) (string, error) {
32     time.Sleep(200 * time.Millisecond)
33     return "some value", nil
34 }
```

Listing 2 中的第 31 行的`search`函数是一个模拟实现，用于模拟长时间运行的操作，如数据库查询或 Web 调用。在这个例子中，强行写为 200ms。

该程序调用该`search`函数，如 Listing 3 所示。

### Listing 3

[https://play.golang.org/p/o6_eMjxMVFv](https://play.golang.org/p/o6_eMjxMVFv)

```
17 // process is the work for the program. It finds a record
18 // then prints it.
19 func process(term string) error {
20     record, err := search(term)
21     if err != nil {
22         return err
23     }
24
25     fmt.Println("Received:", record)
26     return nil
27 }
```

在 Listing 3 中的第 19 行，定义的函数`process`接受一个`string`表示搜索的`term`的参数。
在第 20 行，`term`然后将变量传递给`search`返回记录和错误的函数。
如果发生错误，则错误将返回到第 22 行的调用方。
如果没有错误，则记录将打印在第 25 行。

对于某些应用程序，调用`search` 顺序时产生的延迟可能是不可接受的。
假设`search`无法使功能运行得更快，`process`可以将功能更改为不消耗所产生的总延迟成本`search`。

为此，可以使用 Goroutine，如下面的清单 4 所示。不幸的是，这第一次尝试是错误的，因为它造成了潜在的 Goroutine 泄漏

### Listing 4

[https://play.golang.org/p/m0DHuchgX0A](https://play.golang.org/p/m0DHuchgX0A)

```
38 // result wraps the return values from search. It allows us
39 // to pass both values across a single channel.
40 type result struct {
41     record string
42     err    error
43 }
44
45 // process is the work for the program. It finds a record
46 // then prints it. It fails if it takes more than 100ms.
47 func process(term string) error {
48
49     // Create a context that will be canceled in 100ms.
50     ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
51     defer cancel()
52
53     // Make a channel for the goroutine to report its result.
54     ch := make(chan result)
55
56     // Launch a goroutine to find the record. Create a result
57     // from the returned values to send through the channel.
58     go func() {
59         record, err := search(term)
60         ch <- result{record, err}
61     }()
62
63     // Block waiting to either receive from the goroutine's
64     // channel or for the context to be canceled.
65     select {
66     case <-ctx.Done():
67         return errors.New("search canceled")
68     case result := <-ch:
69         if result.err != nil {
70             return result.err
71         }
72         fmt.Println("Received:", result.record)
73         return nil
74     }
75 }
```

在 Listing 4 中的 第 50 行，重写`process`函数以创建`Context`，`context`的取消函数将在 100ms 内取消。
有关如何使用的更多信息，请`Context`阅读 [golang.org blog post](https://blog.golang.org/context)。

然后在第 54 行，程序创建一个无缓冲的通道，允许 Goroutines 传递该 result 类型的数据。
在第 58 到 61 行，定义了匿名函数，然后放在 Goroutine 中被调用。
此 Goroutine 调用`search`并尝试通过第 60 行的通道发送其返回值。

当`Goroutine`正在进行其工作时，该`process`函数执行第 65 行的`select`。
这边有两种`case`，它们都是通道接收操作。

在第 66 行中有一个`case`接受从 ctx.Done()信道来的数据。
如果 Context 取消（100ms 的持续时间后），将执行此`case`。
如果执行这个`case`，那么`process`将返回一个错误，表明它放弃了等待`search`第 67 行。

或者，第 68 行上的`case`从`ch`通道接收并将值分配给名为`result`的变量。
与之前的顺序实现一样，程序检查并处理第 69 行和第 70 行的错误。
如果没有错误，程序将在第 72 行打印记录并返回`nil`以指示成功。

此重构设置了`process`函数等待`search`完成的最长持续时间。
然而，这种实施也会产生潜在的`Goroutine`泄漏。
想想这个代码中的`Goroutine`正在做什么;
在第 60 行，它在通道上发送。在此通道上发送会阻止执行，直到另一个 Goroutine 准备好接收该值。
在超时情况下，接收器停止等待从`Goroutine` 接收并继续运行。
这将导致 Goroutine 永远阻止等待接收器出现，
这是永远不会发生的。这是 Goroutine 泄漏的时候。

## Fix: Make Some Space

解决此泄漏的最简单方法是将通道从无缓冲通道更改为容量为 1 的缓冲通道。

### Listing 5

[https://play.golang.org/p/u3xtQ48G3qK](https://play.golang.org/p/u3xtQ48G3qK)

```
53     // Make a channel for the goroutine to report its result.
54     // Give it capacity so sending doesn't block.
55     ch := make(chan result, 1)
```

现在在超时情况下，在接收器继续运行之后，search 的 Goroutine 将通过将 result 值放入通道来完成其发送，然后它将返回。Goroutine 的记忆内终将被收回以及该通道的内存。一切都会自然而然地发挥作用。

在`The Behavior of Channels`中威廉·肯尼迪提供了几个关于通道行为的很好的例子，并提供了有关其使用的哲学。该文章“Listing 10”的最后一个示例显示了一个类似于此超时示例的程序。阅读该文章，了解何时使用缓冲通道以及适当的容量级别的更多建议。

## Conclusion

Go 让启动 Goroutines 变得简单，但我们有责任明智地使用它们。在这篇文章中，我展示了如何错误地使用 Goroutines 的一个例子。有许多方法可以创建 Goroutine 泄漏以及使用并发时可能遇到的其他陷阱。在以后的文章中，我将提供更多 Goroutine 泄漏和其他并发陷阱的例子。现在我会给你这个建议; 任何时候你开始 Goroutine 你必须问自己：

- 什么时候会终止？
- 什么可以阻止它终止？

**并发是一种有用的工具，但必须谨慎使用。**
