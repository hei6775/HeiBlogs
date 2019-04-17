# Language Mechanics On Stacks And Pointers

<p align="right">Author: William KennedyMay 18, 2017</p>

原文地址 ：[Language Mechanics On Stacks And Pointers](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html)

## 前言

本系列文章总共四篇，主要帮助大家理解 Go 语言中一些语法结构和其背后的设计原则，包括指针、栈、堆、逃逸分析和值/指针传递。这是第二篇，主要介绍堆和逃逸分析。

以下是本系列文章的索引：

1. Language Mechanics On Stacks And Pointers
2. Language Mechanics On Escape Analysis
3. Language Mechanics On Memory Profiling
4. Design Philosophy On Data And Semantics

## 介绍（Introduction）

在四部分系列的第一部分，我用一个将值共享给·栈的例子介绍了指针结构的基础。而我没有说的是值存在栈之上的情况。为了理解这个，你需要学习值存储的另外一个位置：堆。有这个基础，就可以开始学习逃逸分析。

逃逸分析是编译器用来决定你的程序中值的位置的过程。特别地，编译器执行静态代码分析，以确定一个构造体的实例化值是否会逃逸到堆。在 Go 语言中，你没有可用的关键字或者函数，能够直接让编译器做这个决定。只能够通过你写代码的方式来作出这个决定。

## 堆（Heaps）

堆是内存的第二区域，除了栈之外，用来存储值的地方。堆无法像栈一样能自清理，所以使用这部分内存会造成很大的开销（相比于使用栈）。重要的是，开销跟 `GC`（垃圾收集），即被牵扯进来保证这部分区域干净的程序，有很大的关系。当垃圾收集程序运行时，它会占用你的可用 CPU 容量的 25%。更有甚者，它会造成微秒级的 `“stop the world”` 的延时。拥有 GC 的好处是你可以不再关注堆内存的管理，这部分很复杂，是历史上容易出错的地方。

在 Go 中，会将一部分值分配到堆上。这些分配给 GC 带来了压力，因为堆上没有被指针索引的值都需要被删除。越多需要被检查和删除的值，会给每次运行 GC 时带来越多的工作。所以，分配算法不断地工作，以平衡堆的大小和它运行的速度。

## 共享栈（Sharing Stacks）

在 Go 语言中，不允许 goroutine 中的指针指向另外一个`goroutine`的栈。这是因为当栈增长或者收缩时，`goroutine`中的栈内存会被一块新的内存替换。如果运行时需要追踪指针指向其他的`goroutine`的栈，就会造成非常多需要管理的内存，以至于更新指向那些栈的指针将使`“stop the world”`问题更严重。

这里有一个栈被替换好几次的例子。看输出的第 2 和第 6 行。你会看到`main`函数中的栈的字符串地址值改变了两次。

## 逃逸机制（Escape Mechanics）

任何时候，一个值被分享到函数栈帧范围之外，它都会在堆上被重新分配。这是逃逸分析算法发现这些情况和管控这一层的工作。（内存的）完整性在于确保对任何值的访问始终是准确、一致和高效的。

通过查看这个语言机制了解逃逸分析。

https://play.golang.org/p/Y_VZxYteKO

### Listing 1

```golang
01 package main
02
03 type user struct {
04     name  string
05     email string
06 }
07
08 func main() {
09     u1 := createUserV1()
10     u2 := createUserV2()
11
12     println("u1", &u1, "u2", &u2)
13 }
14
15 //go:noinline
16 func createUserV1() user {
17     u := user{
18         name:  "Bill",
19         email: "bill@ardanlabs.com",
20     }
21
22     println("V1", &u)
23     return u
24 }
25
26 //go:noinline
27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

我使用`go:noinline`指令，阻止在`main`函数中，编译器使用内联代码替代函数调用。内联（优化）会使函数调用消失，并使例子复杂化。我将在下一篇博文介绍内联造成的副作用。

在 Listing 1 中，你可以看到创建 `user`值，并返回给调用者的两个不同的函数。在函数版本 1 中，返回值。
