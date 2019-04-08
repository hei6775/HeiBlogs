# The Go Memory Model

## Introduce

&emsp;&emsp;本文介绍GoLang的内存模型，原文是GoLang官方2014年的文章：[The Go Memory Model](!https://golang.org/ref/mem)
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

&emsp;&emsp;GoLang的内存模型指定了一个条件，在这个条件下，可以保证在一个`goroutine`中读取变量可以被另外一个写入相同变量
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
的顺序。比如说一个`goroutine`中的执行`a=1;b=2`；另一个`goroutine`b的赋值在a的赋值操作之前。

&emsp;&emsp;为了指定读和写的需求，我们定义了`Happens Before`，这是一个在Go程序中的局部的内存操作执行顺序。
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

&emsp;&emsp;这对条件比第一对更强;它要求没有其他写入操作与w或r同时发生。

&emsp;&emsp;在单个goroutine中，没有并发性，所以这两个定义是等价的：读取操作`r`可以观察到最近写入操作`w`对
变量`v`写入的值。当多个`goroutine`访问共享变量`v`时，它们必须使用同步事件来确保`Happens Before`条件读取
观察到所需要的写入。

### Synchronization