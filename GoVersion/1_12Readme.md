## Golang 1.12版本改动

Go team如期在2月末[发布了Go 1.12](https://tip.golang.org/doc/go1.12)版本。从Go 1.12的[Release Notes](https://golang.org/doc/go1.12)粗略
来看，这个版本相较于之前增加了[go modules机制](https://tonybai.com/2018/07/15/hello-go-module)、
[WebAssembly支持](https://github.com/golang/go/wiki/WebAssembly)的[Go 1.11](https://tonybai.com/2018/11/19/some-changes-in-go-1-11/)，
变化略“小”。这也给下一个[Go 1.13](https://github.com/golang/go/milestone/83)版本预留了足够的“惊喜”空间:)。从目前
的plan来看，Go 1.13很可能落地的包括：Go2的几个proposals：[GO 2 number literals](https://github.com/golang/proposal/blob/master/design/19308-number-literals.md)
,[error values](https://github.com/golang/proposal/blob/master/design/29934-error-values.md)和[signed shift counts](https://github.com/golang/proposal/blob/master/design/19113-signed-shift-counts.md)等
，以及[优化版Escape Analysis](https://github.com/golang/go/issues/23109)等。

言归正传，我们来看看Go 1.12版本中值得我们关注的几个变化。

### 一. Go 1.12的可移植性
Go 1.12一如既往的保持了[Go1兼容性规范](https://golang.org/doc/go1compat)，使用Go 1.12编译以往编写的遗留代码，理论上都可以编译通过并
正常运行起来。这是很难得的，尤其是在"Go2"有关proposal逐步落地的“时间节点”，想必Go team为了保持Go1付出了不少额外的努力。

Go语言具有超强的[可移植性](https://tonybai.com/2017/06/27/an-intro-about-go-portability/)。在Go 1.12中，Go又增加了对aix/ppc64、
windows/arm的支持，我们可以在运行于树莓派3的Windows 10 IoT Core上运行Go程序了。

但是对于一些较老的平台系统，Go也不想背上较重的包袱。Go也在逐渐“放弃”一些老版本的系统，比如Go 1.12是最后一个支持macOS 10.10、FreeBSD
10.x的版本。在我的一台Mac 10.9.2的老机器上运行go 1.12将会得到下面错误：

```go
$./go version
dyld: Symbol not found: _unlinkat
  Referenced from: /Users/tony/.bin/go1.12/bin/./go
  Expected in: flat namespace

[1]    2403 trace trap  ./go version
```

###　二. Go modules机制的优化
####　1. GO111MODULE=on时，获取go module不再显式需要go.mod

用过Go 1.11[go module机制](https://tonybai.com/2018/07/15/hello-go-module/)的童鞋可能都遇到过这个问题，那就是在GO111MODULE=on的
情况下(非GOPATH路径)，我要go get某个package时，如果compiler没有在适当位置找到go.mod，就会提示如下错误：

```go
//go 1.11.2

# go get github.com/bigwhite/gocmpp
go: cannot find main module; see 'go help modules'

或

# go get github.com/bigwhite/gocmpp
go: cannot determine module path for source directory /Users/tony/test/go (outside GOPATH, no import comments)
```

这显然非常不方便。为了go get 一个package，我还需要显式地创建一个go.mod文件。在Go 1.12版本中，这个问题被优化掉了。

```go
//go 1.12

# go get github.com/bigwhite/gocmpp
go: finding github.com/bigwhite/gocmpp latest
go: finding golang.org/x/text/encoding/unicode latest
go: finding golang.org/x/text/transform latest
go: finding golang.org/x/text/encoding/simplifiedchinese latest
go: finding golang.org/x/text/encoding latest
go: downloading golang.org/x/text v0.3.0
go: extracting golang.org/x/text v0.3.0
```

其他在go 1.11.x中对go.mod显式依赖的命令，诸如go list、go mod download也在Go 1.12版本中和go get一样不再显式依赖go.mod。

并且在Go 1.12中go module的下载、解压操作支持并发进行，前提是go module的Cache路径：$GOPATH/pkg/mod必须在一个支持file locking的文件
系统中。

####　2. go.mod中增加go指示字段(go directive)

go 1.12版本在go.mod文件中增加了一个go version的指示字段，用于指示该module内源码所使用的 go版本。使用go 1.12创建的go.mod类似下面这样：

```go
# go mod init github.com/bigwhite/test
go: creating new go.mod: module github.com/bigwhite/test
# cat go.mod
module github.com/bigwhite/test

go 1.12

```

按照release notes中的说法，如果go.mod中go指示器指示的版本高于你使用的go tool链版本，那么go也会尝试继续编译。如果编译成功了，那也是ok的。
但是如果编译失败，那么会提示：module编译需要更新版本的go tool链。

我们使用go 1.11.4版本go [compiler编译下面的github.com/bigwhite/test](http://compiler编译下面的上面github.com/bigwhite/test) module的代码：

```go
// main.go

package main

import (
    "fmt"
)

func main() {
    fmt.Println("go world")
}

# go build main.go
# ./main
go world
```

我们看到，虽然go tool chain版本是1.11.4，低于go.mod中的go 1.12，但go 1.11.4仍然尝试继续编译代码，并且顺利通过。

如果我们将代码“故意”修改为下面这样：

```go
//main.go

package main

import (
        "fmt"
)

func main() {
        fmt.Printl("go world") // 这里我们故意将Println写成Printl
}
```

再用go 1.11.4编译这段代码：

```go
# go build main.go
# command-line-arguments
./main.go:8:2: undefined: fmt.Printl
note: module requires Go 1.12
```

我们看到go 1.11.4 compiler提示“需要go 1.12"版本编译器。从这里我们看出，我们可以使用go指示器用作module最低version约束的标识。在没有go
指示器时，我们只能在文档上显式增加这种约束的描述。

不过，这里有一个小插曲，那就是这种不管go.mod中go版本号是多少，仍然尝试继续编译的机制仅适用于go 1.11.4以及后续高版本。从引入go module
的go 1.11到go 1.11.3目前都还不支持这种机制，如果用go 1.11.3尝试编译以下上面的代码，会得到如下结果：

```go
# go build main.go
go build command-line-arguments: module requires Go 1.12
```

go 1.11.3不会继续尝试编译，而是在对比当前go tool chain版本与go.mod中go指示器的version后，给出了错误的提示并退出。

如果非要使用低于go 1.11.4版本的编译器去编译的话，我们可以使用go 1.12工具链的go mod edit -go命令来修改一下go.mod中的版本为go 1.11。
然后再用go 1.11.4以下的版本去编译：

```bash
# go mod edit -go=1.11
# cat go.mod
module github.com/bigwhite/test

go 1.11

# go build main.go  //使用go 1.11.3编译器
```

这样，我们就可用go 1.11~go 1.11.3正常编译源码了。

### 三. 对binary-only package的最后支持
我在2015的一篇文章 《[理解Golang包导入](https://tonybai.com/2015/03/09/understanding-import-packages/)》中提及到Go的编译对源码的
依赖性。对于开源工程中的包，这完全不是问题。但是对于一些商业公司而言，源码是公司资产，是不能作为交付物提供给买方的。为此，Go team
在[Go 1.7](https://tonybai.com/2016/06/21/some-changes-in-go-1-7/)中增加了
对[binary-only package](https://github.com/golang/proposal/blob/master/design/2775-binary-only-packages.md)的机制。

所谓"binary-only package"就是允许开发人员发布不包含源码的二进制形式的package，并且可直接基于该二进制package进行编译。比如下面这个例子：

```bash
// 创建二进制package

# cat $GOPATH/src/github.com/bigwhite/foo.go
package foo

import "fmt"

func HelloGo() {
    fmt.Println("Hello,Go")
}


# go build -o  $GOPATH/pkg/linux_amd64/github.com/bigwhite/foo.a

# ls $GOPATH/pkg/linux_amd64/github.com/bigwhite/foo.a
/root/.go/pkg/linux_amd64/github.com/bigwhite/foo.a

# mkdir temp
# mv foo.go temp

# touch foo.go

# cat foo.go

//go:binary-only-package

package foo

import "fmt"

# cd $GOPATH

# zip -r foo-binary.zip src/github.com/bigwhite/foo/foo.go pkg/linux_amd64/github.com/bigwhite/foo.a
updating: pkg/linux_amd64/github.com/bigwhite/foo.a (deflated 42%)
  adding: src/github.com/bigwhite/foo/foo.go (deflated 11%)
```

我们将foo-binary.zip发布到目标机器上后，进行如下操作：

```bash
# unzip foo-binary.zip -d $GOPATH/
Archive:  foo-binary.zip
  inflating: /root/.go/pkg/linux_amd64/github.com/bigwhite/foo.a
  inflating: /root/.go/src/github.com/bigwhite/foo/foo.go
```

接下来，我们就基于二进制的foo.a来编译依赖它的包:

```bash
//$GOPATH/src/bar.go

package main

import "github.com/bigwhite/foo"

func main() {
        foo.HelloGo()
}

# go build -o bar bar.go
# ./bar
Hello,Go
```


但是经过几个版本的迭代，Go team发现：[对binary-only package越来越难以提供安全支持](https://github.com/golang/go/issues/28152)，
无法保证binary-only package的编译使用的是与最终链接时相同的依赖版本，这很可能会造成因内存问题而
导致的崩溃。并且经过调查，似乎用binary-only package的gopher并不多，并且gopher可以使用plugin、shared library、c-shared library等
来替代binary-only package，以避免源码分发。于是Go 1.12版本将成为支持binary-only package的最后版本。

### 四. 运行时与标准库
经过[Go 1.5](https://tonybai.com/2015/07/10/some-changes-in-go-1-5/)~[Go 1.10](https://tonybai.com/2018/02/17/some-changes-in-go-1-10/)
对运行时，尤其是GC的大幅优化和改善后，Go 1.11、Go 1.12对运行时的改善相比之下都是小幅度的。

在Go 1.12中，一次GC后的内存分配延迟得以改善，这得益于在大量heap依然存在时清理性能的提升。运行时也会更加积极地将释放的内存归还给操作系统，
以应对大块内存分配无法重用已存在的堆空间的问题。在linux上，运行时使用MADV_FREE释放未使用的内存，这更为高效，操作系统内核可以在需要时重用
这些内存。

在多CPU的机器上，运行时的timer和deadline代码运行性能更高了，这对于提升网络连接的deadline性能大有裨益。

标准库最大的改变应该算是对[TLS 1.3](https://www.rfc-editor.org/info/rfc8446)的支持了。不过默认不开启。Go 1.13中将成为默认开启功能。
大多数涉及TLS的代码无需修改，使用Go 1.12重新编译后即可无缝支持TLS 1.3。

另一个”有趣“的变化是syscall包增加了Syscall18，依据syscall包中函数名字惯例，Syscall18支持最多传入18个参数，这个函数的引入是为了Windows
准备的。现在少有程序员会设计包含10多个参数的函数或方法了，这估计也是为了满足Windows中“遗留代码”的需求。

### 五. 工具链及其他

#### 1. go安装包中移除go tour
go tour被从go的安装包中移除了，Go的安装包从go 1.4.x开始到go 1.11.x变得日益“庞大”：以linux/amd64的tar.gz包为例，变化趋势如下：

```bash
go 1.4.3:  53MB
go 1.5.4:  76MB
go 1.6.4:  83MB
go 1.7.6:  80MB
go 1.8.7:  96MB
go 1.9.7:  113MB
go 1.10.8: 97MB
go 1.11.5: 134MB
go 1.12:   121MB


```

后续预计会有更多的非核心功能将会从go安装包中移除来对Go安装包进行瘦身，即便不能瘦身，也至少要保持在现有的size水平上。

本次go tour被挪到：[golang.org/x/tour中了，gopher们可单独安装tour](http://golang.org/x/tour%E4%B8%AD%E4%BA%86%EF%BC%8Cgopher%E4%BB%AC%E5%8F%AF%E5%8D%95%E7%8B%AC%E5%AE%89%E8%A3%85tour%EF%BC%9A):

```bash
# go get -u golang.org/x/tour

# tour //启动tour
```

Go 1.12也是godoc作为web server被内置在Go安装包的最后一个版本，在Go 1.13中该工具也会被从安装包中剔除，如有需要，可像go tour一样通
过go get来单独安装。

####　2. Build cache成为必需
build cache在[Go 1.10](https://tonybai.com/2018/02/17/some-changes-in-go-1-10/)被引入以加快Go包编译构建速度，但是在Go 1.10
和[Go 1.11](https://tonybai.com/2018/11/19/some-changes-in-go-1-11/)中都可以使用GOCACHE=off关闭build cache机制。但是在Go 1.12中build cache
成为必需。如果设置GOCACHE=off，那么编译器将报错：

```bash
# GOCACHE=off  go build github.com/bigwhite/gocmpp
build cache is disabled by GOCACHE=off, but required as of Go 1.12
```

#### 3. Go compiler支持-lang flag
Go compiler支持-lang flag，可以指示编译过程使用哪个版本的Go语法（注意不包括标准库变化等，仅限于语言自身语法）。比如：

```bash
//main.go

package main

import "fmt"

type Int = int

func main() {
        var a Int = 5
        fmt.Println(a)
}


# go run main.go

5
```

上面是一个使用了[Golang 1.9](https://tonybai.com/2017/07/14/some-changes-in-go-1-9/)才引入的type alias语法的Go代码，我们使用
Go 1.12可以正常编译运行它。但是如果我使用-lang flag，指定使用go1.8语法编译
该代码，我们会得到如下错误提示：

```bash
# go build  -gcflags "-lang=go1.8" main.go
# command-line-arguments
./main.go:5:6: type aliases only supported as of -lang=go1.9

换成-lang=go1.9就会得到正确结果：


# go build  -gcflags "-lang=go1.9" main.go
# ./main
5
```


