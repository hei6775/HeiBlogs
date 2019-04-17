# The Behavior Of Channels

<p align="right">Author: William KennedyOctober 24, 2017</p>

原文地址 ：[The Behavior Of Channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)

## Introduction

当我第一次开始使用 Go 的通道时，我错误地将通道视为数据结构。
我将通道视为队列，在 goroutines 之间提供自动同步访问。
这种结构性的理解使我编写了许多糟糕而复杂的并发代码。

我随着时间的推移了解到，最好忘记通道的结构，转而关注他们的行为方式。
所以现在谈到通道，我想到一件事：信号(signaling)。
一个通道允许一个`goroutine`向另一个关于特定事件的`goroutine`发出信号。
信号应该是您使用通道来做的事情的核心。
将通道视为信号机制将允许您使用定义明确且更精确的行为编写更好的代码。

要了解信号如何工作，我们必须了解其三个属性：

- Guarantee Of Delivery

- State

- With or Without Data

这三个属性共同构成了围绕信号传递的设计理念。
在我讨论这些属性之后，我将提供一些代码示例来演示应用这些属性的信号。

## Guarantee Of Delivery

交付保证基于一个问题：“我是否需要保证已收到特定`goroutine`发送的信号？”

换句话说，在 Listing 1 中给出了这个例子：

#### Listing 1

```go
go func() {
    p := <-ch // Receive
}()

ch <- "paper" // Send
```

发送`goroutine`是否需要保证通过 02 号线上的`goroutine`接收通过 05 号线路发送的信号才能继续前进？

根据这个问题的答案，您将知道要使用的两种类型的通道中的哪一种：无缓冲或缓冲。每个通道都围绕交付保证提供不同的行为。

####　 Figure 1 : Guarantee Of Delivery
![86_guarantee_of_delivery.png](./asset/86_guarantee_of_delivery.png)

担保是重要的，如果你不这么认为，我有很多相关的东西想要卖给你。
当然，我只是开个玩笑，但是当你没有生命保障时，你不会感到紧张吗？
在编写并发软件时，充分了解您是否需要保证是至关重要的。
随着我们的继续，您将学习如何决定。

## State

通道的行为直接受其当前状态的影响。通道的状态可以是 `nil`，`open` 或 `closed`。

下面的 Listing 2 显示了如何在这三种状态中声明或放置通道。

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

状态确定发送和接收操作的行为方式。

信号通过信道发送和接收。不要说`read/write`，因为通道不提供 `I / O`.

#### Figure 2 : State

![86_state.png](./asset/86_state.png)

当通道处于`nil`状态时，通道上尝试的任何发送或接收都将导致阻塞。当通道处于`open`状态时，可以发送和接收信号。当通道处于`closed`状态时，不再能够发送信号，但仍然可以接收信号。

这些状态将为您遇到的不同情况提供所需的不同行为。将**State**与**Guarantee Of Delivery**相结合时，您可以开始分析由于您的设计选择而产生的成本/收益。在许多情况下，您还可以通过阅读代码快速发现`Bug`，因为您了解通道的行为方式。

## With and Without Data

需要考虑的最后一个信号属性是您是否需要使用或不使用数据发送信号。

您通过在通道上执行发送来发送数据信号。

#### Listing 3

```
ch <- "paper"
```

当你需要发送带有数据的信号时通常是因为：

- 一个`goroutine`被通知开始一项新任务。
- 一个`goroutine`需要返回结果

你可以通过关闭通道来发送没有数据的信号。

#### Listing 4

```
close(ch)
```

当你发送没有数据的信号时，通常是因为：

- 一个`goroutine`被通知停止它正在做的事情。
- 一个`goroutine`报告说他们没有结果。
- 一个`goroutine`通知它已经处理完并且准备结束，销毁。

这些规则有例外，但这些是主要用例以及我们将在本文中关注的用例。我会认为这些规则的例外是初始的代码味道。

### Signaling With Data

#### Figure 3 : Signaling With Data

![86_signaling_with_data.png](./asset/86_signaling_with_data.png)

这三个通道分别是`Unbuffered`,`Buffered >1` or `Buffered =1`.

<ul>
  <li>
    <b>Guarantee</b>
    <ul>
        <li>无缓冲通道为您提供发送后必须接受的保证</li>
        <li>Because the Receive of the signal <b>Happens Before</b> the Send of the signal completes.</li>
    </ul>
  </li>
    
  <li><b>No Guarantee</b>
      <ul>
        <li>A <b>Buffered</b> channel of size >1 gives you <b>No Guarantee</b> that a signal being sent has been received.</li>
        <li>Because the Send of the signal <b>Happens Before</b> the Receive of the signal completes.</li>
    </ul>
  </li>
  <li><b>Delayed Guarantee</b>
        <ul>
        <li>A <b>Buffered</b> channel of size =1 gives you a <b>Delayed Guarantee</b>. It can guarantee that the previous signal that was sent has been received.</li>
        <li>Because the <b>Receive</b> of the First Signal, <b>Happens Before</b> the Send of the Second Signal completes.</li>
    </ul>
    </li>
</ul>

缓冲区的大小绝不能是随机数，必须始终针对某些明确定义的约束进行计算。
计算中没有无穷大，无论是时间还是空间，一切都必须有一些明确的约束。

### Signaling Without Data

没有数据的信号主要用于取消。
它允许一个`goroutine`发出信号通知另一个`goroutine`取消他们正在做的事情并继续前进。取消可以使用无缓冲和缓冲通道来实现，但是当没有数据将被发送时使用缓冲通道是代码味道。

#### Figure 4 : Signaling Without Data

![86_signaling_without_data.png](./asset/86_signaling_without_data.png)

内置函数 close 用于在没有数据的情况下发出信号。
如上面`State`部分所述，您仍然可以在已关闭的通道上接收信号。
实际上，封闭通道上的任何接收都不会阻塞，并且接收操作总是返回。

在大多数情况下，您希望使用标准库`context`包来实现没有数据的信号。
`context`包使用下面的`Unbuffered`通道用于信号，内置函数接近信号而没有数据。

如果您选择使用自己的通道进行取消，而不是`context`包，则您的通道应为`chan struct {}`类型。这是`zero-space`，用于表示这是用于发信号的惯用方式。

## Scenarios

有了这些属性，进一步了解它们在实践中如何工作的最佳方法是运行一系列代码方案。
当我在阅读和编写基于通道的代码时，我喜欢将 goroutines 视为人。
这种可视化确实很有帮助，我将在下面使用它作为辅助。

### Signal With Data - Guarantee - Unbuffered Channels

当您需要知道已收到正在发送的信号时，会出现两种情况。这些是等待任务和等待结果。

#### Scenario 1 - Wait For Task

假设你成为一名经理并雇用一名新员工。在这种情况下，您希望新员工执行任务，但他们需要等到您准备好之后。这是因为你需要在它们开始之前给它们一张纸。

#### Listing 5

```
 func waitForTask() {
    ch := make(chan string)

    go func() {
        p := <-ch

       // Employee performs work here.

       // Employee is done and free to go.
    }()

   time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

   ch <- "paper"
}
```

在 Listing 5 的第 02 行，创建了一个`Unbuffered`通道，其中包含字符串数据作为随信号一起发送的属性。然后在第 04 行，雇佣一名员工并告诉他在工作前等待 05 号线上的信号。第 05 行是通道接收，导致员工在等待您将要发送的`paper`时阻塞。一旦员工收到`paper`，员工就会完成工作，然后就可以自由地完成工作了。

您作为经理与您的新员工同时工作。因此，在第 04 行雇用员工后，您会发现自己（在第 12 行）正在做您需要做的事情以解除阻塞并向员工发出信号。请注意，不知道准备这张纸需要多长时间才能发送。

最终，您已准备好向员工发出信号。在第 14 行，您执行带有数据的信号，数据就是那张纸。由于正在使用无缓冲通道，因此您可以保证在发送操作完成后员工已收到纸张。接收发生在发送之前。

技术上你所知道的是，当你的通道发送操作完成时，员工有`paper`。在两个通道操作之后，调度程序可以选择执行它想要的任何语句。由您或员工执行的下一行代码是不确定的。这意味着使用`print`语句可以欺骗你的事情顺序。

#### Scenario 2 - Wait For Result

在下一个场景中，情况正好相反。这次您希望新员工在被雇用时立即执行任务，您需要等待他们的工作结果。你需要等待，因为你需要他们的`paper`才能继续。

#### Listing 6

```
func waitForResult() {
    ch := make(chan string)

    go func() {
        time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

        ch <- "paper"

        // Employee is done and free to go.
    }()

    p := <-ch
}
```

在清单 6 的第 02 行，创建了一个`Unbuffered`通道，其中包含将随信号一起发送字符串数据的属性
。然后在第 04 行，雇用一名员工并立即投入工作。在第 04 行雇用员工后，您会发现自己在第 12 行等待`paper`。

员工在第 05 行完成工作后，他们会通过执行带数据的通道发送将结果发送给您。
由于这是一个无缓冲通道，因此接收在发送之前发生，并且保证员工已收到结果。
一旦员工获得此保证，他们就可以自由地完成任务。
在这种情况下，您不知道雇员完成任务需要多长时间。

#### Cost/Benefit

无缓冲通道可确保接收到正在发送的信号。
这很棒，但没有什么是免费的。这种保证的成本是未知的延迟。
在等待任务方案中，员工不知道发送该文件需要多长时间。
在等待结果方案中，您不知道雇员将该结果发送给您的时间有多长。

### Signal With Data - No Guarantee - Buffered Channels >1

当您不需要知道已收到正在发送的信号时，这两种情况就会发挥作用：`Fan Out`and`Drop`.

缓冲通道具有明确定义的空间，可用于存储正在发送的数据。那么你如何决定你需要多少空间呢？回答这些问题：

- 我是否有完美的定义多少的工作任务需要去完成？
- 有多少工作？
- 如果我的员工无法继续执行，我可以放弃任何新工作吗？
- 多少出色的工作让我有能力？
- 如果我的程序意外终止，我愿意接受多大程度的风险
- 在缓冲区中等待的任何内容都将丢失。

如果这些问题对您正在建模的行为没有意义，那么使用大于 1 的缓冲通道的代码气味可能是错误的。

#### Scenario 1 - Fan Out

`Fan Out`模式允许您针对同时工作的问题抛出明确数量的员工。由于每个任务都有一名员工，因此您确切知道将收到多少报告。您可以确保框中有足够的空间来接收所有这些报告。这样，您的员工无需等待您提交报告。然而，如果他们在同一时间或几乎同时到达盒子，他们每个人都需要轮流将报告放在你的盒子里。

想象一下，你再次担任经理，但这次你雇佣了一个员工队伍。您希望每个员工执行一项单独的任务。当每个员工完成任务时，他们需要向您提供`paper`，该报告必须放在您桌面上的盒子中。

#### Listing 7

```
func fanOut() {
    emps := 20
    ch := make(chan string, emps)

    for e := 0; e < emps; e++ {
        go func() {
            time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
             ch <- "paper"
        }()
    }

    for emps > 0 {
        p := <-ch
        fmt.Println(p)
        emps--
    }
}
```

在清单 7 的第 03 行上，创建了一个缓冲通道，
其中包含将随信号一起发送字符串数据的属性。
这次使用 20 个缓冲区创建了通道，这要归功于第 02 行声明的 emps 变量。

在 05 至 10 行之间，雇用了 20 名员工，他们立即开始工作。
你不知道每个员工要在 07 行上走多长时间。
然后在第 08 行，员工发送`paper`，但这次发送不会阻止等待接收。
由于每个员工的方框都有空间，因此通道上的发送仅与可能希望在同一时间或几乎同时发送报告的其他员工竞争。

第 12 到 16 行之间的代码就是你。
在这里，您可以等待所有 20 名员工完成工作并发送报告。
在第 12 行，您处于循环中，在第 13 行，您将被阻塞在通道接收等待您的报告。
收到报告后，报告将打印在第 14 行，并且本地计数器变量将递减以指示员工已完成。

#### Scenario 2 - Drop

`Drop`模式允许您在员工满负荷时放弃工作。
这样做的好处是可以继续接受客户端的工作，而不会在接受该工作时施加背压或延迟。
这里的关键是知道什么时候你真正有能力，所以你不会低估或过度承诺你将尝试完成的工作量。
通常，集成测试或指标是帮助您识别此编号所需的。

想象一下，您再次成为经理，并聘请一名员工完成工作。
您有一个您希望员工执行的单独任务。
当员工完成任务时，您不必知道他们已完成任务。
所有重要的是你是否可以在盒子里放置新的工作。
如果您无法执行发送，那么您就知道您的盒子已满，而且员工已达到容量。
在这一点上，新工作需要被抛弃，因此事情可以继续发展。

#### Listing 8

```
01 func selectDrop() {
02     const cap = 5
03     ch := make(chan string, cap)
04
05     go func() {
06         for p := range ch {
07             fmt.Println("employee : received :", p)
08         }
09     }()
10
11     const work = 20
12     for w := 0; w < work; w++ {
13         select {
14             case ch <- "paper":
15                 fmt.Println("manager : send ack")
16             default:
17                 fmt.Println("manager : drop")
18         }
19     }
20
21     close(ch)
22 }
```

在 Listing 8 的第 03 行，创建了一个缓冲通道，其中包含将随信号一起发送字符串数据的属性。这次通道创建了 5 个缓冲区，这要归功于 02 行声明的上限常量。

在 05 到 09 行之间，雇用一名员工来处理工作。A 范围用于通道接收。每次收到一个`paper`，就会在第 07 行处理。

在第 11 行到第 19 行之间，您尝试向员工发送 20 张`paper`。这次使用 select 语句在第 14 行的第一个 case 中执行`send`。因为在第 16 行的`select`中使用了`default`子句，如果`send`将要阻塞，因为缓冲区中没有空间，通过执行第 17 行放弃发送。

最后在第 21 行，对通道调用内置函数`close`。这将在没有数据的情况下向员工发出信号，一旦他们完成指定的工作就可以自由行动。

#### Cost/Benefit

大小大于 1 带缓冲的通道不保证始终接收正在发送的信号。远离这种保证是有益的，这是两个`goroutines`之间通信的减少或没有延迟。在** Fan Out **场景中，每个将发送报告的员工都有一个缓冲区。在`Drop`方案中，缓冲区是针对容量进行测量的，如果达到容量，则会丢弃工作，因此可以继续移动。

在这两个选项中，我们不得不忍受这种缺乏保证，因为延迟的减少更为重要。从零到最小延迟的要求不会对系统的整体逻辑造成问题。

### Signal With Data - Delayed Guarantee - Buffered Channel 1

当需要知道在发送新信号之前是否已收到先前发送的信号时，`Wait For Tasks`方案将起作用。

#### Scenario 1 - Wait For Tasks

在这种情况下，您有一名新员工，但他们要做的不仅仅是一项任务。
你将一个接一个地为他们提供许多任务。
但是，他们必须先完成每项任务才能开始新任务。由于它们一次只能处理一项任务，因此在切换工作之间可能存在延迟问题。
如果可以减少延迟而不失去员工正在处理下一个任务的保证，那么它可能有所帮助。

这是大小为 1 的缓冲通道有益的地方。
如果一切都以您和员工之间的预期速度运行，那么您都不需要等待另一个。
每次发送一张`paper`时，缓冲区都是空的。
每当您的员工完成更多工作时，缓冲区就会满了。它是工作流程的完美对称。

最好的部分就是这个。
如果您在任何时候尝试发送一张`paper`而您不能因为缓冲区已满，您知道您的员工遇到了问题而您就停止了。
这是延迟保证的来源。
当缓冲区为空并执行发送时，您可以保证您的员工已完成您发送的最后一项工作。
如果您执行发送而您不能，则可以保证他们没有。

#### Listing 9

```
01 func waitForTasks() {
02     ch := make(chan string, 1)
03
04     go func() {
05         for p := range ch {
06             fmt.Println("employee : working :", p)
07         }
08     }()
09
10     const work = 10
11     for w := 0; w < work; w++ {
12         ch <- "paper"
13     }
14
15     close(ch)
16 }
```

在 Listing 9 的第 02 行，创建了一个大小为 1 的带缓冲通道，其中包含字符串数据随信号一起发送的属性。
在第 04 至 08 行之间雇用一名雇员来处理工作。
`for range`用于通道接收。每次收到一张`paper`，都会在 06 行处理。

在第 10 行到第 13 行之间，您开始将您的任务发送给员工。
如果您的员工可以尽可能快地运行，那么您的员工之间的延迟就会降低。
但是，每次发送都会成功执行，您可以保证您提交的最后一项工作正在进行中。

最后在第 15 行，对通道调用内置函数 close。
这将向员工发出没有数据的信号，他们已经完成并且可以自由行动
。但是，您提交的最后一项工作将在`for range`终止之前收到（`flushed`）。

### Signal Without Data - Context

在最后一个场景中，
您将看到如何使用`context`包中的`Context`值取消正在运行的`goroutine`。
这一切都可以通过利用关闭的无缓冲通道来执行发送没有数据的信号。

您最后一次是经理，并雇用一名员工完成工作。
这次你不愿意等待一段不明的时间让员工完成。
您处于离散的最后期限，如果员工没有及时完成，您不愿意等待。

#### Listing 10

```
01 func withTimeout() {
02     duration := 50 * time.Millisecond
03
04     ctx, cancel := context.WithTimeout(context.Background(), duration)
05     defer cancel()
06
07     ch := make(chan string, 1)
08
09     go func() {
10         time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
11         ch <- "paper"
12     }()
13
14     select {
15     case p := <-ch:
16         fmt.Println("work complete", p)
17
18     case <-ctx.Done():
19         fmt.Println("moving on")
20     }
21 }
```

在 Listing 10 的第 02 行，声明了持续时间值，表示员工必须完成任务的时间。
此行在第 04 行上用于创建`context.Context`值，超时为`50`毫秒。
`context`包中的`WithTimeout`函数返回`Context`值和取消函数。

`context`包创建一个`goroutine`，一旦满足持续时间，它将关闭与`Context`值关联的`Unbuffered`通道。
无论情况如何，您都有责任调用`cancel`函数。
这将清除为`Context`创建的内容。可以多次调用`cancel`函数。

在第 05 行，一旦该函数终止，取消函数将被推迟执行。
在第 07 行创建了一个大小为 1 的缓冲通道，员工将使用该通道向您发送工作结果。
然后在 09 到 12 行雇用员工并立即投入工作。你不知道雇员完成需要多长时间。

在第 14 行到第 20 行之间，您可以使用`select`语句在两个通道上接收。
在第 15 行收到，您等待员工向您发送结果。
在第 18 行接收，你等着看`context`包是否会发出 50 毫秒的信号。
您首先收到的信号将是处理过的信号。

此算法的一个重要方面是使用大小为 1 的缓冲通道。
如果员工没有及时完成，您将继续前进而不向员工发出任何通知。
从员工的角度来看，他们总是会在第 11 行向您发送报告，如果您在那里或者没有收到报告，他们就会失明。
如果您使用未缓冲的通道，员工将永久阻塞尝试向您发送报告（如果您继续前进）。
这会造成`goroutine`泄漏。
因此，使用大小为 1 的缓冲通道来防止这种情况发生。

## Conclusion

`guarantees`、`channel state`、`sending`的理解对于在并发中使用通道非常重要。
它们将指导您实现所编写的并发程序和算法所需的最佳行为。
它们将帮助您找到 Bug 并嗅出可能有害的代码。

在这篇文章中，我分享了一些示例程序，展示信号的属性如何在不同的场景中工作。每条规则都有例外，但这些模式是一个很好的基础。

查看这些概述，作为有效思考和使用通道的时间和方式的摘要：

### Language Mechanics

- 使用通道来编排和协调`goroutines`。
  - 使用通道发送信号，而不是使用共享数据。
  - 发送带有数据或者不带有数据的信号。
  - 使用通道用于同步机制来控制对共享状态的访问。
    - 有些情况下，使用通道可以更简单，但初始化是一个问题。
- Unbuffered channels
  - 接受发生在发送之前（`Receive happens before the Send.`）
  - 好处：百分百确认信号一定会被接受到
  - 不足：接收信号时的延迟时间未知。
- Buffered channels
  - 发送发生在接受之前（`Send happens before the Receive.`）
  - 好处：减少接受信号的延迟
  - 不足：无法保证信号已经被接受
    - 越大的缓存，保证能力越低
    - 大小为 1 的缓存通道给你一个延迟发送的保证
- Closing channels
  - 关闭发生在接受之前（`Close happens before the Receive (like Buffered)`）
  - 发送没有数据的信号。
  - 完美的信号取消和到期时间。
- nil channels
  - 发送和接受都是阻塞的
  - 无法发送信号
  - 适用于限速或短期停机。

### Design Philosophy

- 如果任何给定在通道的发送**能够**导致发送的`goroutine`阻塞：
  - 不允许使用大小大于 1 的缓存通道。
    - 使用缓存大于 1 必须有原因以及经过了计算。
  - 必须知道当发送的`gouroutine`阻塞时发生了什么
- 如果任何给定在通道的发送**不能够**导致发送的`goroutine`阻塞：
  - 每个发送都有准确的缓冲数量。
    - `Fan Out`模式
  - 你必须准确计算缓存的容量。
    - `Drop`模式
- 缓冲越少越好。
  - 在考虑缓冲时，不要考虑性能。
  - 缓冲可以帮助减少信号之间的阻塞延迟。
    - 将阻塞延迟减少到零并不一定意味着更好的吞吐量。
    - 如果一个缓冲为您提供足够好的吞吐量，那么保留它。
    - 询问缓冲大于 1 并测量大小。
    - 找到可能提供足够吞吐量的最小缓冲区
