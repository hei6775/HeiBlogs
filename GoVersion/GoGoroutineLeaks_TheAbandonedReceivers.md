# Goroutine Leaks - The Abandoned Receivers

<p align="right">Author: Jacob WalkerDecember 19, 2018</p>

原文地址 ：[Goroutine Leaks - The Abandoned Receivers](https://www.ardanlabs.com/blog/2018/12/goroutine-leaks-the-abandoned-receivers.html)

## 简介

## 泄漏：被遗弃的接收者

**For this leak example you will see multiple Goroutines blocked waiting to receive values that will never be sent.**

### Listing 1

[https://play.golang.org/p/Jtpla_UvrmN](https://play.golang.org/p/Jtpla_UvrmN)

```
35 // processRecords is given a slice of values such as lines
36 // from a file. The order of these values is not important
37 // so the function can start multiple workers to perform some
38 // processing on each record then feed the results back.
39 func processRecords(records []string) {
40
41     // Load all of the records into the input channel. It is
42     // buffered with just enough capacity to hold all of the
43     // records so it will not block.
44
45     total := len(records)
46     input := make(chan string, total)
47     for _, record := range records {
48         input <- record
49     }
50     // close(input) // What if we forget to close the channel?
51
52     // Start a pool of workers to process input and send
53     // results to output. Base the size of the worker pool on
54     // the number of logical CPUs available.
55
56     output := make(chan string, total)
57     workers := runtime.NumCPU()
58     for i := 0; i < workers; i++ {
59         go worker(i, input, output)
60     }
61
62     // Receive from output the expected number of times. If 10
63     // records went in then 10 will come out.
64
65     for i := 0; i < total; i++ {
66         result := <-output
67         fmt.Printf("[result  ]: output %s\n", result)
68     }
69 }
70
71 // worker is the work the program wants to do concurrently.
72 // This is a blog post so all the workers do is capitalize a
73 // string but imagine they are doing something important.
74 //
75 // Each goroutine can't know how many records it will get so
76 // it must use the range keyword to receive in a loop.
77 func worker(id int, input <-chan string, output chan<- string) {
78     for v := range input {
79         fmt.Printf("[worker %d]: input %s\n", id, v)
80         output <- strings.ToUpper(v)
81     }
82     fmt.Printf("[worker %d]: shutting down\n", id)
83 }
```

在第 39 行的`Listing 1`中，定义了一个名为`processRecords`的函数。
该函数接受一个`string`类型的切片。在第 46 行，一个名为`input`的缓冲通道。
第 47 和 48 行运行一个循环，复制`string`切片中的每个值并将它们发送到通道。
`input`创建的通道具有足够的容量来保存切片中的每个值，
因此第 48 行上的通道发送都不会阻塞。此通道是用于在多个`Goroutines`之间分配值的管道。

接下来在第 56 到 60 行，该程序创建了一个`Goroutines`池来接收来自管道中的工作。
在第 56 行，创建了一个名为`output`的缓冲通道; 这是每个`Goroutine`将发送其结果的地方。
第 57 到 59 行运行循环以使用该`worker`函数创建多个`Goroutines` 。
`Goroutines`的数量等于机器上的逻辑 CPU 数量。
循环变量的副本`i`以及`input`和`output`通道都传递给`Goroutine`。

该 worker 函数定义在在第 77 行。
函数的签名定义 input 为 a <-chan string，这意味着它是一个只接收通道。
该函数也接受 output，chan<- string 这意味着它是一个只发送通道。

在函数内部，`Goroutines`接受来自`input`通道的消息，
在第 78 行使用`range`循环从通道接收。
在通道上使用在`range`循环接收，直到通道关闭并且没有值。
对于每次迭代，将接收到的值分配给 v 在第 79 行打印的迭代变量。
然后在第 80 行，worker 函数传递 v 给`strings`.ToUpper 返回 new 的函数 string。
工作协程立即在 output 通道上发送新的`string`。

回到`processRecords`函数中，执行已经向下移动到第 65 行，
在那里运行另一个循环。该循环迭代，直到它接收并处理了来自 output 通道的所有值。
在第 66 行，该`processRecords`函数等待从一个工作者`Goroutines`接收一个值。
接收到的值打印在第 67 行。当程序收到每个输入的值时，它退出循环并终止该功能。

运行此程序打印转换后的数据，
因此它似乎工作，但该程序正在泄漏多个`Goroutines`。
该程序从未到达第 82 行，该行将宣布工人正在关闭。
即使在`processRecords`函数返回之后，
每个工作者`Goroutines`仍处于活动状态并等待第 78 行的输入。
通过一个通道接收直到通道关闭并为空。问题是程序永远不会关闭。

##　 Fix: Signal Completion

###　 Listing 2
[https://play.golang.org/p/QNsxbT0eIay](https://play.golang.org/p/QNsxbT0eIay)

```
45     total := len(records)
46     input := make(chan string, total)
47     for _, record := range records {
48         input <- record
49     }
50     close(input)
```

关闭缓冲区中仍有值的缓冲通道是有效的;通道仅关闭发送而不是接收。
工作人员 Goroutines 运行范围输入将通过缓冲区工作，直到它们被通知通道已关闭。
这可以让工人在终止之前完成循环。

## Conclusion

正如前一篇文章中所提到的，Go 使得启动 Goroutines 变得简单，但是你有责任仔细使用它们。在这篇文章中，我展示了另一个可以轻易做出的 Goroutine 错误。还有很多方法可以创建 Goroutine 泄漏以及使用并发时可能遇到的其他陷阱。未来的帖子将继续讨论这些陷阱。与往常一样，我将继续重复这一建议：“ 如果不知道它会如何停止，就不要开始使用 goroutine ”。

并发是一种有用的工具，但必须谨慎使用。
