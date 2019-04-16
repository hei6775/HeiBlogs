# The Behavior Of Channels

<p align="right">Author: William KennedyOctober 24, 2017</p>

原文地址 ：[The Behavior Of Channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)

## Introduction
当我第一次开始使用Go的通道时，我错误地将通道视为数据结构。
我将通道视为队列，在goroutines之间提供自动同步访问。
这种结构性的理解使我编写了许多糟糕而复杂的并发代码。

我随着时间的推移了解到，最好忘记渠道的结构，并关注他们的行为方式。
所以现在谈到渠道，我想到一件事：信号。
一个通道允许一个goroutine向另一个关于特定事件的goroutine发出信号。
信令是您应该对频道所做的一切的核心。
将通道视为信号机制将允许您使用定义明确且更精确的行为编写更好的代码。

要了解信令如何工作，我们必须了解其三个属性：

- Guarantee Of Delivery

- State

- With or Without Data

这三个属性共同构成了围绕信号传递的设计理念。
在我讨论这些属性之后，我将提供一些代码示例来演示应用这些属性的信令。


## Guarantee Of Delivery

交付保证基于一个问题：“我是否需要保证已收到特定goroutine发送的信号？”

换句话说，在清单1中给出了这个例子：

#### Listing 1
```go
go func() {
    p := <-ch // Receive
}()

ch <- "paper" // Send
```

发送goroutine是否需要保证通过02号线上的goroutine接收通过05号线路发送的纸张才能继续前进？

根据这个问题的答案，您将知道要使用的两种类型的通道中的哪一种：无缓冲或缓冲。每个渠道都围绕交付保证提供不同的行为。

####　Figure 1 : Guarantee Of Delivery
![86_guarantee_of_delivery.png](./asset/86_guarantee_of_delivery.png)

担保是重要的，如果你不这么认为，我有很多东西要卖给你。
当然，我正在试着开个玩笑，但是当你没有生命保障时，你不会感到紧张吗？
在编写并发软件时，充分了解您是否需要保证是至关重要的。
随着我们的继续，您将学习如何决定。

## State

信道的行为直接受其当前状态的影响。通道的状态可以是零，开放或关闭。

下面的清单2显示了如何在这三种状态中声明或放置通道。

#### Listing 2
```go
// ** nil channel

// A channel is in a nil state when it is declared to its zero value
var ch chan string

// A channel can be placed in a nil state by explicitly setting it to nil.
ch = nil

// ** open channel

// A channel is in a open state when it’s made using the built-in function make.
ch := make(chan string)    


// ** closed channel

// A channel is in a closed state when it’s closed using the built-in function close.
close(ch)
```


## With and Without Data

## Signaling With Data

## Signaling Without Data

## Conclusion