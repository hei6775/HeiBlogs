# [Command go](http://docs.studygolang.com/cmd/go/#hdr-Test_packages)

<a href="#Start a bug report">启动错误报告</a>&emsp;&emsp;&emsp;<a href="#File types">文件类型</a>

<a href="#Compile packages and dependencies">编译包和依赖项</a>&emsp;&emsp;&emsp;<a href="#The go.mod file">go.mod 文件</a>

<a href="#Remove object files and cached files">删除目标文件和缓存的文件</a>&emsp;&emsp;&emsp;<a href="#GOPATH environment variable">GOPATH 环境变量</a>

<a href="#Show documentation for package or symbol">显示包或符号的文档</a>&emsp;&emsp;&emsp;<a href="#GOPATH and Modules">GOPATH 和模块</a>

<a href="#Print Go environment information">打印 Go 环境信息</a>&emsp;&emsp;&emsp;<a href="#Internal Directories">内部目录</a>

<a href="#Update packages to use new APIs">更新包以使用新 API</a>&emsp;&emsp;&emsp;<a href="#Vendor Directories">Vendor 目录</a>

<a href="#Gofmt (reformat) package sources">Gofmt（重新格式化）包源码</a>&emsp;&emsp;&emsp;<a href="#Module proxy protocol">模块代理协议</a>

<a href="#Generate Go files by processing source">通过处理源生成 Go 文件</a>&emsp;&emsp;&emsp;<a href="#Import path syntax">导入路径语法</a>

<a href="#Download and install packages and dependencies">下载并安装包和依赖项</a>&emsp;&emsp;&emsp;<a href="#Relative import paths">相对导入路径</a>

<a href="#Compile and install packages and dependencies">编译并安装包和依赖项</a>&emsp;&emsp;&emsp;<a href="#Remote import paths">远程导入路径</a>

<a href="#List packages or modules">列出包或模块</a>&emsp;&emsp;&emsp;<a href="#Import path checking">导入路径检查</a>

<a href="#Module maintenance">模块维护</a>&emsp;&emsp;&emsp;<a href="#Modules, module versions, and more">模块，模块版本等</a>

<a href="#Download modules to local cache">将模块下载到本地缓存</a>&emsp;&emsp;&emsp;<a href="#Preliminary module support">初步模块支持</a>

<a href="#Edit go.mod from tools or scripts">从工具或脚本编辑 go.mod</a>&emsp;&emsp;&emsp;<a href="#Defining a module">定义一个模块</a>

<a href="#Print module requirement graph">打印模块要求图</a>&emsp;&emsp;&emsp;<a href="#The main module and the build list">主模块和构建列表</a>

<a href="#Initialize new module in current directory">在当前目录中初始化新模块</a>&emsp;&emsp;&emsp;<a href="#Maintaining module requirements">维护模块要求</a>

<a href="#Add missing and remove unused modules">添加缺失并删除未使用的模块</a>&emsp;&emsp;&emsp;<a href="#Pseudo-versions">伪版本</a>

<a href="#Make vendored copy of dependencies">制作依赖项的销售副本</a>&emsp;&emsp;&emsp;<a href="#Module queries">模块查询</a>

<a href="#Verify dependencies have expected content">验证依赖项是否具有预期内容</a>&emsp;&emsp;&emsp;<a href="#Module compatibility and semantic versioning">模块兼容性和语义版本控制</a>

<a href="#Explain why packages or modules are needed">解释为什么需要包或模块</a>&emsp;&emsp;&emsp;<a href="#Module code layout">模块代码布局</a>

<a href="#Compile and run Go program">编译并运行 Go 程序</a>&emsp;&emsp;&emsp;<a href="#Module downloading and verification">模块下载和验证</a>

<a href="#Test packages">测试包</a>&emsp;&emsp;&emsp;<a href="#Modules and vendoring">模块和销售</a>

<a href="#Run specified go tool">运行指定的 go 工具</a>&emsp;&emsp;&emsp;<a href="#Module-aware go get">模块感知 go get</a>

<a href="#Print Go version">打印 Go 版本</a>&emsp;&emsp;&emsp;<a href="#Package lists and patterns">包列表和模式</a>

<a href="#Report likely mistakes in packages
Build modes">报告包中可能出现的错误</a>&emsp;&emsp;&emsp;<a href="#Testing flags">测试标志</a>

<a href="#Calling between Go and C">在 Go 和 C 之间调用</a>&emsp;&emsp;&emsp;<a href="#Testing functions">测试函数</a>

<a href="#Build modes">构建模式</a>&emsp;&emsp;&emsp;<a href="#Subdirectories">子目录</a>

<a href="#Build and test caching">构建和测试缓存</a>&emsp;&emsp;&emsp;

<a href="#Environment variables">环境变量</a>&emsp;&emsp;&emsp;

Go is a tool for managing Go source code.

用法:

```
go <command> [arguments]
```

The commands are:

```golang
bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         download and install packages and dependencies
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages
Use "go help <command>" for more information about a command.
```

Additional help topics:

```golang
buildmode   build modes
c           calling between Go and C
cache       build and test caching
environment environment variables
filetype    file types
go.mod      the go.mod file
gopath      GOPATH environment variable
gopath-get  legacy GOPATH go get
goproxy     module proxy protocol
importpath  import path syntax
modules     modules, module versions, and more
module-get  module-aware go get
packages    package lists and patterns
testflag    testing flags
testfunc    testing functions
Use "go help <topic>" for more information about that topic.
```

# <a name="Start a bug report">启动错误报告</a>

用法:

```
go bug
```

Bug 打开默认浏览器并启动新的错误报告。该报告包含有用的系统信息。

# <a name="Compile packages and dependencies">编译包和依赖项</a>

语法:

```
go build [-o output] [-i] [build flags] [packages]
```

`build` 会编译导入的包及其依赖项，但不会安装生成的结果。

如果构建的参数是一个`.go`文件的列表，则 `build` 会将它们视为指定单个包的源文件列表。

编译单个主程序包时，`build` 会将生成的可执行文件以第一个源文件来进行命名（例如'go build ed.go rx.go'write'ed'或'ed.exe'）或源代码目录（ 'go build unix / sam'写'sam'或'sam.exe'）。编写 Windows 可执行文件时会添加“.exe”后缀。

在编译多个包或单个非主包时，`build` 会编译包但丢弃生成的对象，仅用于检查是否可以构建包。

编译包时，`build` 会忽略以“\_test.go”结尾的文件。

`-o`标志仅在编译单个包时允许，强制构建将可执行结果或对象以重命名的方式输出文件，而不是最后两段中描述的默认行为。

`-i`标志安装作为目标依赖项的软件包。

`build`标志与`build`，`clean`，`get`，`install`，`list`，`run`和`test` 命令共享：

```
-a
	强制重建已经是最新的软件包。
-n
	打印命令但不运行它们。
-pn
	可以并行运行的程序（例如构建命令或测试二进制文件）的数量。默认值是可用的CPU数。
-race
	启用数据竞争检测。
	仅支持linux / amd64，freebsd / amd64，darwin / amd64和windows / amd64。
-msan
	支持与内存清理程序的互操作。
	仅支持在linux / amd64，linux / arm64上，
	并且仅支持Clang / LLVM作为主机C编译器。
-v
	在编译时打印包的名称。
-work
	打印临时工作目录的名称，退出时不要删除它。
-x
	打印命令。

-asmflags  '[pattern =] arg list'
	传递每个 go tool asm 调用的参数。
-buildmode  mode
	构建模式使用。有关更多信息，请参阅“go help buildmode”。
-compiler
	要使用的编译器名称，如runtime.Compiler（gccgo或gc）。
-gccgoflags  '[pattern =] arg list'
	传递每个 gccgo compiler/linker 调用的参数。
-gcflags  '[pattern =] arg list'
	传递每个go tool compile调用的参数。
-installsuffix  suffix
	要在程序包安装目录的名称中使用的后缀，以便将输出与默认构建分开。
    如果使用-race标志，则安装后缀会自动设置为race，或者，如果明确设置，则会附加_race。 同样对于-msan标志。 使用需要非默认编译标志的-buildmode选项具有类似的效果。
-ldflags  '[pattern =] arg list'
	传递每个 go tool link调用的参数。
-linkshared
	link 之前使用-buildmode = shared创建的共享库。
-mod  mode
	模块下载模式使用：readonly或vendor。
	有关更多信息，请参阅“go help modules”。
-pkgdir dir
	从dir目录安装并加载所有包，而不是通常的位置。
	例如，使用非标准配置构建时，请使用-pkgdir将生成的包保留在单独的位置。
-tags  'tag list'
	在构建期间要考虑满足的以空格分隔的构建标记列表。有关构建标记的更多信息，请参阅
	go/build包文档中的构建约束说明。
-toolexec  'cmd args'
	用于调用vet和asm等工具链程序的程序。
	例如，运行asm 'cmd args /path/to/asm <asm>的参数'。
```

`-asmflags`，`-gccgoflags`，`-gcflags` 和`-ldflags` 标志接受
以空格分隔的参数列表，以在构建期间传递给基础工具。要在列表中的元
素中嵌入空格，请使用单引号或双引号将其括起来。参数列表可以在包模
式和等号之后，这限制了该参数列表的使用以构建匹配该模式的包（有关
包模式的描述，请参阅'go help packages'）。如果没有模式，参数列
表仅适用于命令行上指定的包。可以用不同的模式重复标志，以便为不同
的包组指定不同的参数。如果包与多个标志中给出的模式匹配，则命令
行上的最新匹配将获胜。例如，'go build -gcflags = -S fmt'

有关指定包的更多信息，请参阅“go help packages”。有关安装包和二进制文件的更多信息，请运行'go help gopath'。有关在 Go 和 C / C ++之间调用的更多信息，请运行'go help c'。

注意：`Build` 遵守某些约定，例如'go help gopath'所描述的约定。但是，并非所有项目都遵循这些惯例。具有自己的约定或使用单独的软件构建系统的安装可以选择使用较低级别的调用，例如“go tool compile”和“go tool link”，以避免构建工具的一些开销和设计决策。

另见：`go install`，`go get`，`go clean`。

# <a name="Remove object files and cached files">删除目标文件和缓存的文件</a>

用法：

```
go clean [clean flags] [build flags] [packages]
```

`Clean`从包源目录中删除目标文件。go 命令在临时目录中构建大多数对象，因此 go clean 主要关注其他工具留下的目标文件或对于 go build 的手动调用。

具体来说，clean 从与导入路径对应的每个源目录中删除以下文件：

```
_obj/  旧对象目录，从 Makefiles留下的
_test/ 旧测试目录，从 Makefiles留下的
_testmain.go 旧gotest文件，从 Makefiles留下的
test.out   旧测试日志，从 Makefiles留下的
build.out  旧测试日志，从 Makefiles留下的
*.[568ao]  对象文件，从 Makefiles留下的

DIR（.exe）      来自 go build
DIR.test（.exe） 来自 go test -c
MAINFILE（.exe） 来自 go build MAINFILE.go
*.so           来自 SWIG
```

在列表中，DIR 表示目录的最终路径元素，MAINFILE 是构建程序包时未包含的目录中任何 Go 源文件的基本名称。

-i 标志导致 clean 删除对应的已安装存档或二进制文件（'go install'创建的东西）。

-n 标志导致 clean 打印它将执行的 remove 命令，但不运行它们。

-r 标志使 clean 以递归方式应用于导入路径命名的包的所有依赖项。

-x 标志导致 clean 在执行它们时打印 remove 命令。

-cache 标志导致 clean 删除整个 go build 的缓存。

-testcache 标志导致 clean 使 go build 缓存中的所有测试结果失效。

-modcache 标志导致 clean 删除整个模块下载缓存，包括版本化依赖项的解压缩源代码。

有关构建标志的更多信息，请参阅“go help build”。

有关指定包的更多信息，请参阅“go help packages”。

# <a name="Show documentation for package or symbol">显示包或符号的文档</a>

用法：

```
go doc [-u] [-c] [package|[package.]symbol[.methodOrField]]
```

Doc 打印与其参数（a package, const, func, type, var, method, or struct field）标识的项目相关联的文档注释，然后是每个“下”的第一级项的一行摘要。 item（包的包级声明，类型的方法等）。

Doc 接受零个，一个或两个参数。

没有参数，也就是说，当运行时

```
go doc
```

它在当前目录中打印包的包文档。如果一个包是一个命令（package mian），除非提供-cmd 标志，否则将从表示中删除包的导出符号。

当使用一个参数运行时，该参数被视为要记录的项的类似 Go 语法的表示。参数选择取决于 GOROOT 和 GOPATH 中安装的内容，以及参数的形式，其中示意性之一：

```
go doc <pkg>
go doc <sym>[.<methodOrField>]
go doc [<pkg>.]<sym>[.<methodOrField>]
go doc [<pkg>.][<sym>.]<methodOrField>
```

此参数列表中与参数匹配的第一个项目是打印其文档的项目。（请参阅下面的示例。）但是，如果参数以大写字母开头，则假定它标识当前目录中的符号或方法。

对于包，扫描的顺序是以广度优先顺序词汇确定的。也就是说，所呈现的包是与搜索匹配的包，并且最接近根并且词法上首先在其层级的级别。在 GOPATH 之前，GOROOT 树总是被完整扫描。

如果没有指定或匹配包，则选择当前目录中的包，因此“go doc Foo”显示当前包中符号 Foo 的文档。

包路径必须是合格路径或路径的正确后缀。go 工具的常用包机制不适用：包路径元素之类的。和...不是由 go doc 实现的。

当使用两个参数运行时，第一个必须是完整的包路径（不仅仅是后缀），第二个是符号，或带有方法或结构字段的符号。这类似于 godoc 接受的语法：

```
go doc <pkg> <sym>[.<methodOrField>]
```

在所有形式中，当匹配符号时，参数中的小写字母与两种情况匹配，但大写字母完全匹配。这意味着如果不同的符号具有不同的情况，则包中可能存在小写参数的多个匹配。如果发生这种情况，则打印所有匹配的文档。

例子：

```
go doc
    显示当前包的文档。
go doc Foo
    在当前包中显示 Foo 的文档。
    （Foo 以大写字母开头，因此它与
    包路径不匹配。）
go doc encoding / json
    显示 encoding / json 包的文档。
go doc json
    编码/ json 的简写。
go doc json.Number（或 go doc json.number）
    显示 json.Number 的文档和方法摘要。
go doc json.Number.Int64（或 go doc json.number.int64）
    显示 json.Number 的 Int64 方法的文档。
go doc cmd / doc
    显示 doc 命令的包文档。
    去 doc -cmd cmd / doc
    在 doc 命令中显示包文档和导出的符号。
go doc template.new
    显示 html / template 的新功能的文档。
    （html / template 在文本/模板之前是词法上的）
go doc text / template.new＃一个参数
    显示文本/模板的新功能的文档。
go doc text / template new＃两个参数
    显示文本/模板的新功能的文档。

至少在当前树中，这些调用都打印了json.Decoder 的 Decode 方法的文档：

go doc json.Decoder.Decode
go doc json.decoder.decode
go doc json.decode
cd go/src/encoding/json; go doc decode
```

标志：

```
-all
    显示包的所有文档。
-c
    在匹配符号时尊重大小写。
-cmd
    将命令（包 main）视为常规包。
    否则，在显示程序包的顶级文档时，将隐藏程序包主导出的符号。
-src
    显示符号的完整源代码。这将
    显示其声明和
    定义的完整 Go 源，例如函数定义（包括
    正文），类型声明或封闭 const
    块。因此输出可能包括未导出的
    细节。
-u
    显示未导出的符号，方法和字段的文档。
```

# <a name="Print Go environment information">打印 Go 环境信息</a>

用法：

```
go env [-json] [var ...]
```

Env 打印 Go 环境信息。

默认情况下，env 将信息打印为 shell 脚本（在 Windows 上，即批处理文件）。如果给出一个或多个变量名作为参数，则 env 在其自己的行上打印每个命名变量的值。

-json 标志以 JSON 格式而不是 shell 脚本打印环境。

有关环境变量的更多信息，请参阅“go help environment”。

# <a name="Update packages to use new APIs">更新包以使用新 API</a>

用法：

```
go fix [packages]
```

Fix 在导入路径命名的包上运行 Go fix 命令。

有关修复的更多信息，请参阅“go doc cmd/fix”。有关指定包的更多信息，请参阅“go help packages”。

要使用特定选项运行修复，请运行“go tool fix”。

另见：go fmt，go vet。

# <a name="Gofmt (reformat) package sources">Gofmt（重新格式化）包源</a>

用法：

```
go fmt [-n] [-x] [packages]
```

Fmt 在导入路径命名的包上运行命令'gofmt -l -w'。它打印修改的文件的名称。

有关 gofmt 的更多信息，请参阅“go doc cmd/gofmt”。有关指定包的更多信息，请参阅“go help packages”。

-n 标志打印将要执行的命令。-x 标志在打印执行的命令。

要使用特定选项运行 gofmt，请运行 gofmt 本身。

另见：go fix，go vet。

# <a name="Generate Go files by processing source">通过处理源码生成 Go 文件</a>

用法：

```
go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]
```

Generate 由现有文件中的指令描述的运行命令。这些命令可以运行任何进程，但目的是创建或更新 Go 源文件。

Go generate 永远不会通过 go build，go get，go test 等自动运行。它必须明确运行。

Go 生成扫描文件中的指令，这些指令是表单的行，

```
//go:generate command argument...
```

（注意：“//go”中没有前导空格和空格）其中 command 是要运行的生成器，对应于可以在本地运行的可执行文件。它必须位于 shell 路径（gofmt），完全限定的路径（/usr/you/bin/mytool）或命令别名中，如下所述。

为了向人类和机器工具传达生成代码，生成的源码应该具有与以下正则表达式匹配的行（在 Go 语法中）：

```
^// Code generated .* DO NOT EDIT\.$
```

该行可能出现在文件的任何位置，但通常位于开头附近，因此很容易找到。

请注意，go generate 不会解析文件，因此看起来像注释或多行字符串中的指令的行将被视为指令。

该指令的参数是空格分隔的标记或双引号字符串，它们在运行时作为单独的参数传递给生成器。

带引号的字符串使用 Go 语法并在执行之前进行评估; 带引号的字符串作为生成器的单个参数出现。

Go 运行生成器时生成几个变量：

```
$ GOARCH
	执行架构（arm，amd64等）
$ GOOS
执行操作系统（linux，windows 等）
$ GOFILE
	文件的基本名称。
$ GOLINE
    源文件中指令的行号。
$ GOPACKAGE
	包含指令的文件包的名称。
$ DOLLAR
    美元符号.
```

除了变量替换和引用字符串评估之外，在命令行上不执行诸如“globbing”之类的特殊处理。

作为运行命令之前的最后一步，任何具有字母数字名称的环境变量（例如`$GOFILE`或`$HOME`）的调用都将在整个命令行中进行扩展。变量扩展的语法是所有操作系统上的`$NAME`。由于评估的顺序，变量甚至在引用的字符串内扩展。如果未设置变量 NAME，则`$NAME`将扩展为空字符串。

表格的指示，

```
//go:generate -command xxx args...
```

仅为此源文件的剩余部分指定字符串 xxx 表示由参数标识的命令。这可用于创建别名或处理多字生成器。例如，

```
//go:generate -command foo go tool foo
```

指定命令“foo”表示生成器“go tool foo”。

按命令行上给出的顺序生成进程包，一次一个。如果命令行列出.go 文件，则将它们视为单个包。在包中，按文件名顺序生成处理包中的源文件，一次一个。在源文件中，按照它们在文件中出现的顺序生成运行生成器，一次一个。

如果任何生成器返回错误退出状态，“go generate”将跳过该包的所有进一步处理。

`The generator`在包的源目录中运行。

Go generate 接受一个特定的标志：

```
-run =“”
    如果非空，则指定正则表达式以选择
    其完整原始源文本（不包括任何尾随空格和最终换行符）与
    表达式匹配的指令。
```
它还接受标准构建标志，包括-v，-n 和-x。-v 标志在处理包时打印包和文件的名称。-n 标志打印将要执行的命令。-x 标志在执行时打印命令。

有关构建标志的更多信息，请参阅“go help build”。

有关指定包的更多信息，请参阅“go help packages”。

# <a name="Download and install packages and dependencies">下载并安装包和依赖项</a>

用法：
```
go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]
```
获取导入路径指定的包及其依赖项的下载。然后安装命名包，比如'go install'。

-d 标志指示在下载软件包后停止; 也就是说，它指示不安装软件包。

-f 标志仅在设置-u 时有效，强制 get -u 不验证每个包是否已从其导入路径隐含的源控制存储库中检出。如果源是原始的本地分支，这可能很有用。

-fix 标志指示 get 在解析依赖项或构建代码之前在下载的包上运行修复工具。

-insecure 标志允许从存储库中提取并使用不安全的方案（如 HTTP）解析自定义域。谨慎使用。

-t 标志指示 get 还下载构建指定包的测试所需的包。

-u 标志指示 get 使用网络更新命名包及其依赖项。默认情况下，get 使用网络检出丢失的包，但不使用它来查找现有包的更新。

-v 标志启用详细进度和调试输出。

Get 还接受构建标志来控制安装。请参阅'go help build'。

当检测出新包时，get 将创建目标目录 `GOPATH/src/<import-path>`。如果有多个 GOPATH ，则 get 使用第一个。有关详细信息，请参阅：'go help gopath'。

当检测出或者更新包时，查找与本地安装的 Go 版本匹配的分支或标记。最重要的规则是，如果本地安装运行版本为“go1”，则搜索名为“go1”的分支或标记。如果不存在此类版本，则会检索包的默认分支。

当检测出或更新 Git 存储库时，它还会更新存储库引用的任何 git 子模块。

Get永远不会检出或更新存储在`vendor `目录中的代码。

有关指定包的更多信息，请参阅“go help packages”。

有关“go get”如何找到要下载的源代码的更多信息，请参阅“go help importpath”。

本文描述了使用 GOPATH 管理源代码和依赖项时 get 的行为。如果 go 命令在模块感知模式下运行，则 get 的标志和效果的细节会发生变化，就像'go help get'一样。请参阅“go help modules”和“go help module-get”。

另见：go build，go install，go clean。

# <a name="Compile and install packages and dependencies">编译并安装包和依赖项</a>

用法：
```
go install [-i] [build flags] [packages]
```
安装编译并安装导入路径命名的包。

-i 标志也会安装命名包的依赖项。

有关构建标志的更多信息，请参阅“go help build”。有关指定包的更多信息，请参阅“go help packages”。

另见：go build，go get，go clean。

# <a name="List packages or modules">列出包或模块</a>

用法：
```
go list [-f format] [-json] [-m] [list flags] [build flags] [packages]
```
list列出了命名包，每行一个。最常用的标志是-f 和-json，它们控制为每个包打印的输出形式。下面记录的其他列表标志控制更具体的细节。

默认输出显示包导入路径：
```
bytes
encoding/json
github.com/gorilla/mux
golang.org/x/net/html
```
-f 标志使用包模板的语法指定列表的备用格式。默认输出等效于 -f `{{.ImportPath}}`。传递给模板的结构是：

```
type Package struct {
    Dir           string   //包含包源的目录
    ImportPath    string   //导入 dir 中包的路径
    ImportComment string   //导入时的路径对包语句的注释
    Name          string   //包名称
    Doc           string   //包文档字符串
    Target        string   //安装路径
    Shlib         string   //包含此包的共享库（仅在-linkshared 时设置）
    Goroot        bool     //是 Go root 中的这个包吗？
    Standard      bool     //是标准 Go 库的这个包的一部分吗？
    Stale         bool     //会'安装'为这个包做任何事情吗？
    StaleReason   string   //解释 Stale == true
    Root          string   //转到 root 或 Go 路径 dir 包含此包
    ConflictDir   string   //此目录阴影 in $GOPATH
    BinaryOnly    bool     //仅二进制包：无法从源再次编译
    ForTest       string   //包仅用于命名测试
    Export        string   //包含导出数据的文件（当使用-export 时）
    Module        *Module  //有关包含模块的信息，如果有的话（可以是 nil）
    Match         []string //与此包匹配的命令行模式
    DepOnly       bool     //包只是一个依赖项，没有明确列出


    //源文件
    GoFiles         []string //.go源文件（不包括CgoFiles，TestGoFile，XTestGoFiles）
    CgoFiles        []string //.go源文件导入“C”
    CompiledGoFiles []string //.go文件呈现给编译器（使用-compiled时） ）
    IgnoredGoFiles  []string //.go源文件因构建约束而被忽略
    CFiles          []string //.c源文件
    CXXFiles        []string //.cc，.cxx和.cpp源文件
    MFiles          []string //.m源文件
    HFiles          []string //.h，.hh，.hpp和.hxx源文件
    FFiles          []string //.f，.F，.for和.f90 Fortran源文件
    SFiles          []string // .s source 文件
    SwigFiles       []string // .swig 文件
    SwigCXXFiles    []string // .swigcxx 文件
    SysoFiles       []string // .syso object files to add to archive
    TestGoFiles     []string // _test.go files in package
    XTestGoFiles    []string // _test.go files outside package

    // Cgo指令
    CgoCFLAGS    []string // cgo: C编译器的标志
    CgoCPPFLAGS  []string // cgo: C预处理器的标志
    CgoCXXFLAGS  []string // cgo: C ++编译器的标志
    CgoFFLAGS    []string // cgo: Fortran编译器的标志
    CgoLDFLAGS   []string // cgo: flags for linker
    CgoPkgConfig []string // cgo: pkg-config names

    //依赖性信息
    Imports      []string          //此包使用的导入路径
    ImportMap    map[string]string // map from source import to ImportPath (identity entries omitted)
    Deps         []string          // all (recursively) imported dependencies
    TestImports  []string          // imports from TestGoFiles
    XTestImports []string          // imports from XTestGoFiles

    //错误信息
    Incomplete bool            //此包或依赖项有错误
    Error      *PackageError   //错误加载包
    DepsErrors []*PackageError //错误加载依赖项
}
```
Packages 存储在`vendor`目录中,报告一个ImportPath，其中包含`vendor`目录的路径（例如，“d/vendor/p”而不是“p”），以便 ImportPath 唯一地标识包的给定副本。Imports，Deps，TestImports 和 XTestImports 列表还包含这些扩展的导入路径。有关 vendoring 的更多信息，请参阅 golang.org/s/go15vendor。

错误信息（如果有）是
```
type PackageError struct {
    ImportStack   []string //从命令行命名的包到此一个的最短路径
    Pos           string   // 错误位置 (if present, file:line:col)
    Err           string   // the error itself
}

}
```
模块信息是 Module 结构，在下面列表-m 的讨论中定义。

模板函数`join`调用 `strings.Join`。

模板函数`context`返回构建上下文，定义如下：

```
type Context struct {
    GOARCH        string   //目标体系结构
    GOOS          string   //目标操作系统
    GOROOT        string   // Go root
    GOPATH        string   // Go path
    CgoEnabled    bool     //是否可以使用 cgo 使用
    UseAllFiles   bool     // use files regardless of +build lines, file names
    Compiler      string   //编译器在计算目标路径时假设
    BuildTags     []string // build constraints to match in +build lines
    ReleaseTags   []string // releases the current release is compatible with
    InstallSuffix string   // suffix to use in the name of the install dir
}
```
有关这些字段含义的更多信息，请参阅 `go/build` 包的 Context 类型的文档。

-json 标志使包数据以 JSON 格式打印，而不是使用模板格式。

-compiled 标志导致 list 将 CompiledGoFiles 设置为呈现给编译器的 Go 源文件。通常，这意味着它会重复 GoFiles 中列出的文件，然后还会添加通过处理 CgoFiles 和 SwigFiles 生成的 Go 代码。Imports 列表包含来自 GoFiles 和 CompiledGoFiles 的所有导入的并集。

-deps 标志使列表不仅迭代命名包而且迭代所有依赖关系。它在深度优先的后序遍历中访问它们，以便仅在所有依赖项之后列出包。未在命令行中明确列出的包将 DepOnly 字段设置为 true。

-e 标志更改错误包的处理，无法找到或错误的包。默认情况下，list 命令会为每个错误的包打印一个错误标准错误，并在通常的打印过程中省略所考虑的包。使用-e 标志，list 命令永远不会将错误打印到标准错误，而是使用通常的打印处理错误的包。错误的包将具有非空的 ImportPath 和非零错误字段; 其他信息可能会或可能不会丢失（归零）。

-export 标志使列表将 Export 字段设置为包含给定包的最新导出信息的文件的名称。

-find 标志使列表标识命名包但不解析它们的依赖关系：Imports 和 Deps 列表将为空。

-test 标志使列表不仅报告命名包而且报告测试二进制文件（对于带有测试的包），以准确地向源代码分析工具传达测试二进制文件的构造方式。报告的测试二进制文件的导入路径是包的导入路径，后跟“.test”后缀，如“math / rand.test”。在构建测试时，有时需要专门为该测试重建某些依赖项（最常见的是测试包本身）。报告的针对特定测试二进制文件重新编译的包的导入路径后跟一个空格和括号中的测试二进制文件的名称，如`“math/rand [math/rand.test]”`或“`regexp [sort.test]`”。ForTest 字段也设置为正在测试的包的名称（`“math/rand”`或`“sort”`在之前的例子）

`Dir`，`Target`，`Shlib`，`Root`，`ConflictDir` 和 `Export `文件路径都是绝对路径。

默认情况下，列表 GoFiles，CgoFiles 等保存 Dir 中的文件名（即相对于 Dir 的路径，而不是绝对路径）。使用-compiled 和-test 标志时添加的生成文件是引用生成的 Go 源文件的缓存副本的绝对路径。虽然它们是 Go 源文件，但路径可能不会以“.go”结尾。

-m 标志使列表列出模块而不是包。

列出模块时，-f 标志仍指定应用于 Go 结构的格式模板，但现在是 Module 结构：
```
type Module struct {
    Path     string       //模块路径
    Version  string       //模块版本
    Versions []string     //可用模块版本（带-versions）
    Replace  *Module      // replaced by this module
    Time     *time.Time   // time version was created
    Update   *Module      // available update, if any (with -u)
    Main     bool         // is this the main module?
    Indirect bool         // is this module only an indirect dependency of main module?
    Dir      string       // directory holding files for this module, if any
    GoMod    string       // path to go.mod file for this module, if any
    Error    *ModuleError //错误加载模块
}

type ModuleError struct {
    Err string //错误本身
}
```
默认输出是打印模块路径，然后打印有关版本和替换的信息（如果有）。例如，`'go list -m all'`可能会打印：
```
my/main/module
golang.org/x/text v0.3.0 => /tmp/text
rsc.io/pdf v0.1.1
```
请注意，更换模块后，其“替换”字段描述替换模块，其“目录”字段设置为替换的源代码（如果存在）。（也就是说，如果 Replace 为非 nil，则 Dir 设置为 Replace.Dir，无法访问替换的源代码。）

-u 标志添加有关可用升级的信息。当给定模块的最新版本比当前模块更新时，列表-u 将 Module 的 Update 字段设置为有关较新模块的信息。Module 的 String 方法通过在当前版本之后格式化括号中的较新版本来指示可用的升级。例如，'go list -m -u all'可能会打印：
```
my/main/module
golang.org/x/text v0.3.0 [v0.4.0] => /tmp/text
rsc.io/pdf v0.1.1 [v0.1.2]
```
（对于tools，`'go list -m -u -json all'`可能更方便解析。）

-versions 标志导致 list 将 Module 的 Versions 字段设置为该模块的所有已知版本的列表，按照语义版本排序，最早到最新。该标志还更改默认输出格式以显示模块路径，后跟空格分隔的版本列表。

list -m 的参数被解释为模块列表，而不是包。主模块是包含当前目录的模块。活动模块是主模块及其依赖项。没有参数，list -m 显示主模块。使用参数，list -m 显示参数指定的模块。任何活动模块都可以通过其模块路径指定。特殊模式“all”指定所有活动模块，首先是主模块，然后是依赖于模块路径的依赖项。包含“...”的模式指定模块路径与模式匹配的活动模块。表单路径@ version 的查询指定该查询的结果，该查询不限于活动模块。有关模块查询的更多信息，请参阅“go help modules”。

模板函数“module”采用单个字符串参数，该参数必须是模块路径或查询，并将指定的模块作为 Module 结构返回。如果发生错误，结果将是具有非零错误字段的 Module 结构。

有关构建标志的更多信息，请参阅“go help build”。

有关指定包的更多信息，请参阅“go help packages”。

有关模块的更多信息，请参阅“go help modules”。

# <a name="Module maintenance">模块维护</a>

Go mod 提供对模块操作的访问。

请注意，对所有 go 命令都内置了对模块的支持，而不仅仅是'go mod'。例如，应使用“go get”来完成依赖项的日常添加，删除，升级和降级。有关模块功能的概述，请参阅“go help modules”。

用法：

`go mod <command> [arguments]`

命令是：
```
download    下载模块到本地缓存
edit        编辑 go.mod 从工具或脚本
graph       打印模块需求图
init        初始化当前目录中的新模块
tidy        添加缺失并删除未使用的模块
vendor      make vendored copy of dependencies
verify      验证依赖项已预期内容
why         解释为什么包或模块是必需的
```

有关命令的更多信息，请使用“go help mod <command>”。

# <a name="Download modules to local cache">将模块下载到本地缓存</a>

用法：

`go mod download [-json] [modules]`

下载下载命名模块，可以是模块模式选择主模块的依赖关系或模板路径@版本的模块查询。没有参数，下载适用于主模块的所有依赖项。

go 命令将在普通执行期间根据需要自动下载模块。“go mod download”命令主要用于预填充本地缓存或计算 Go 模块代理的答案。

默认情况下，下载会将错误报告为标准错误，否则将保持静默。-json 标志导致下载将一系列 JSON 对象打印到标准输出，描述每个下载的模块（或失败），对应于此 Go 结构：

```
type Module struct {
    Path     string // 模块路径
    Version  string // 模块版本
    Error    string // 错误加载模块
    Info     string // 缓存的绝对路径.info 文件
    GoMod    string // 缓存的绝对路径.mod 文件
    Zip      string // 绝对缓存.zip 文件的路径
    Dir      string // 缓存源根目录的绝对路径
    Sum      string // checksum for path, version (as in go.sum)
    GoModSum string // checksum for go.mod (as in go.sum)
}
```

有关模块查询的更多信息，请参阅“go help modules”。

# <a name="Edit go.mod from tools or scripts">从工具或脚本编辑 go.mod</a>

用法：

`go mod edit [editing flags] [go.mod]`

Edit 提供了一个命令行界面，用于编辑 go.mod，主要用于工具或脚本。它只读取 go.mod; 它没有查找有关模块的信息。默认情况下，编辑读取和写入主模块的 go.mod 文件，但可以在编辑标志之后指定不同的目标文件。

编辑标志指定一系列编辑操作。

-fmt 标志重新格式化 go.mod 文件而不进行其他更改。使用或重写 go.mod 文件的任何其他修改也暗示了这种重新格式化。唯一需要此标志的是如果没有指定其他标志，如'go mod edit -fmt'。

-module 标志更改模块的路径（go.mod 文件的模块行）。

-require = path @ version 和-droprequire = path 标志在给定的模块路径和版本上添加和删除需求。请注意，-require 会覆盖路径上的所有现有要求。这些标志主要用于了解模块图的工具。用户应该更喜欢“go get path @ version”或“go get path @ none”，这样可以根据需要进行其他 go.mod 调整，以满足其他模块施加的限制。

-exclude = path @ version 和-dropexclude = path @ version flags 为给定的模块路径和版本添加和删除排除项。请注意，如果排除已存在，则--exclude = path @ version 是无操作。

-replace = old [@v] = new [@v]和-dropreplace = old [@v]标志添加和删除给定模块路径和版本对的替换。如果省略旧@v 中的@v，则替换适用于具有旧模块路径的所有版本。如果省略 new @ v 中的@v，则新路径应该是本地模块根目录，而不是模块路径。请注意，-replace 会覆盖旧[@v]的任何现有替换。

可以重复-require，-droprequire，-exclude，-dropexclude，-replace 和-dropreplace 编辑标志，并且按照给定的顺序应用更改。

-go = version 标志设置预期的 Go 语言版本。

-print 标志以文本格式打印最终的 go.mod，而不是将其写回 go.mod。

-json 标志以 JSON 格式打印最终的 go.mod 文件，而不是将其写回 go.mod。JSON 输出对应于这些 Go 类型：
```
type Module struct {
	Path string
	Version string
}

type GoMod struct {
	Module  Module
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
}

type Require struct {
	Path string
	Version string
	Indirect bool
}

type Replace struct {
	Old Module
	New Module
}
```

请注意，这仅描述了 go.mod 文件本身，而不是间接引用的其他模块。对于构建可用的完整模块集，请使用'go list -m -json all'。

例如，工具可以通过解析'go mod edit -json'的输出来获取 go.mod 作为数据结构，然后可以通过使用-require，-exclude 等调用'go mod edit'来进行更改。

# <a name="Print module requirement graph">打印模块需求图</a>

用法：

`go mod graph`

图形以文本形式打印模块需求图（应用了替换）。输出中的每一行都有两个以空格分隔的字段：一个模块和一个要求。每个模块都被标识为表单路径@版本的字符串，但主模块除外，它没有@version 后缀。

# <a name="Initialize new module in current directory">在当前目录中初始化新模块</a>

用法：

`go mod init [module]`

Init 初始化并将新的 go.mod 写入当前目录，实际上创建了一个以当前目录为根的新模块。文件 go.mod 必须不存在。如果可能，init 将从导入注释（请参阅“go help importpath”）或版本控制配置中猜测模块路径。要覆盖此猜测，请将模块路径作为参数提供。

# <a name="Add missing and remove unused modules">添加缺失并删除未使用的模块</a>

用法：

`go mod tidy [-v]`

Tidy 确保 go.mod 匹配模块中的源代码。它添加了构建当前模块的包和依赖项所需的任何缺少的模块，并删除了未提供任何相关包的未使用模块。它还将任何缺少的条目添加到 go.sum 并删除任何不必要的条目。

-v 标志导致整理将有关已删除模块的信息打印到标准错误

# <a name="Make vendored copy of dependencies">制作依赖项的销售副本</a>

用法：

`go mod vendor [-v]`
供应商重置主模块的供应商目录，以包括构建和测试所有主模块包所需的所有包。它不包括销售包裹的测试代码。

-v 标志使供应商将出售模块和包的名称打印为标准错误。

# <a name="Verify dependencies have expected content">验证依赖项是否具有预期内容</a>

用法：

`go mod verify`

验证检查当前模块的依赖关系（存储在本地下载的源缓存中）自下载以来未被修改。如果所有模块都未修改，请验证打印“所有模块已验证”。否则，它会报告哪些模块已被更改，并导致'go mod'以非零状态退出。

# <a name="Explain why packages or modules are needed">解释为什么需要包或模块</a>

用法：

`go mod why [-m] [-vendor] packages...`

为什么在导入图中显示从主模块到每个列出的包的最短路径。如果给出-m 标志，为什么将参数视为模块列表并找到每个模块中任何包的路径。

默认情况下，为什么查询与“go list all”匹配的包的图形，其中包括对可访问包的测试。-vendor 标志导致为什么要排除依赖项的测试。

输出是一系列节，一个用于命令行上的每个包或模块名称，用空行分隔。每个节都以注释行“#package”或“＃module”开头，给出目标包或模块。后续行给出了导入图的路径，每行一个包。如果未从主模块引用包或模块，则该节将显示指示该事实的单个带括号的注释。

例如：
```
$ go mod why golang.org/x/text/language golang.org/x/text/encoding
# golang.org/x/text/language
rsc.io/quote
rsc.io/sampler
golang.org/x/text/language

# golang.org/x/text/encoding
(main module does not need package golang.org/x/text/encoding)
$
```

# <a name="Compile and run Go program">编译并运行 Go 程序</a>

用法：

`go run [build flags] [-exec xprog] package [arguments...]`

运行编译并运行命名的主 Go 包。通常，包被指定为.go 源文件的列表，但它也可以是与单个已知包匹配的导入路径，文件系统路径或模式，如“go run”。或'去运行我的/ cmd'。

默认情况下，'go run'直接运行已编译的二进制文件：'a.out arguments ...'。如果给出-exec 标志，'go run'使用 xprog 调用二进制文件：

`'xprog a.out arguments...'.`

如果未给出-exec 标志，则 GOOS 或 GOARCH 与系统默认值不同，并且可以在当前搜索路径上找到名为 go* \$ GOOS* \$ GOARCH_exec 的程序，“go run”使用该程序调用二进制文件，例如'go_nacl_386_exec a.out arguments ...'。这允许在模拟器或其他执行方法可用时执行交叉编译的程序。

Run 的退出状态不是已编译二进制文件的退出状态。

有关构建标志的更多信息，请参阅“go help build”。有关指定包的更多信息，请参阅“go help packages”。

另见：go build。

# <a name="Test packages">测试包</a>

用法：

`go test [build/test flags] [packages] [build/test flags & test binary flags]`

“Go test”自动测试导入路径命名的包。它以以下格式打印测试结果的摘要：

```
ok   archive/tar   0.011s
FAIL archive/zip   0.022s
ok   compress/gzip 0.033s
...
```
然后是每个失败包的详细输出。

“Go test”重新编译每个包以及名称与文件模式“\* _test.go”匹配的任何文件。这些附加文件可以包含测试函数，基准函数和示例函数。有关更多信息，请参阅“go help testfunc”。每个列出的包都会导致执行单独的测试二进制文件。名称以“_”开头的文件（包括“\_test.go”）或“。” 被忽略了。

声明具有后缀“\_test”的包的测试文件将被编译为单独的包，然后链接并与主测试二进制文件一起运行。

go 工具将忽略名为“testdata”的目录，使其可用于保存测试所需的辅助数据。

作为构建测试二进制文件的一部分，测试运行对包及其测试源文件进行检查以识别重大问题。如果发现任何问题，请去测试报告那些并且不运行测试二进制文件。仅使用默认 go vet 检查的高可信子集。该子集是：'atomic'，'bool'，'buildtags'，'nilfunc'和'printf'。您可以通过“go doc cmd / vet”查看这些和其他兽医测试的文档。要禁用 go vet 的运行，请使用-vet = off 标志。

所有测试输出和汇总行都打印到 go 命令的标准输出，即使测试将它们打印到自己的标准错误。（go 命令的标准错误保留用于构建测试的打印错误。）

Go 测试以两种不同的模式运行：

第一种称为本地目录模式，在没有包参数的情况下调用 go test 时发生（例如，'go test'或'go test -v'）。在此模式下，go test 将编译当前目录中的包源和测试，然后运行生成的测试二进制文件。在此模式下，禁用缓存（下面讨论）。包测试完成后，go test 打印一条摘要行，显示测试状态（'ok'或'FAIL'），包名称和已用时间。

第二种叫做包列表模式，在使用显式包参数调用 go test 时发生（例如'go test math'，'go test。/ ...'，甚至'go test。'）。在此模式下，go test 编译并测试命令行中列出的每个包。如果包测试通过，则 go test 仅打印最终的'ok'摘要行。如果包测试失败，则 go test 打印完整的测试输出。如果使用-bench 或-v 标志调用，则即使传递包测试，go test 也会打印完整输出，以显示请求的基准测试结果或详细日志记录。

仅在包列表模式下，go test 缓存成功的包测试结果，以避免不必要的重复运行测试。当可以从缓存中恢复测试结果时，go test 将重新显示先前的输出，而不是再次运行测试二进制。发生这种情况时，请在测试打印'（缓存）'代替摘要行中的已用时间。

缓存中匹配的规则是运行涉及相同的测试二进制文件，命令行上的标志完全来自一组受限制的“可缓存”测试标志，定义为-cpu，-list，-parallel，-run ，-short 和-v。如果运行 go 测试在此集合之外有任何测试或非测试标志，则不会缓存结果。要禁用测试缓存，请使用除可缓存标志之外的任何测试标志或参数。显式禁用测试缓存的惯用方法是使用-count = 1。在包的源根目录（通常是\$ GOPATH）中打开文件或参考环境变量的测试仅匹配文件和环境变量未更改的未来运行。缓存的测试结果在任何时候都被视为执行，因此无论-timeout 设置如何，都将缓存并重用成功的包测试结果。

除了构建标志之外，'go test'本身处理的标志是：

-args
将命令行的其余部分（-args 之后的所有内容）
传递给测试二进制文件，取消解释并保持不变。
由于此标志占用命令行的其余部分，
因此包列表（如果存在）必须出现在此标志之前。

-c
将测试二进制文件编译为 pkg.test 但不运行它
（其中 pkg 是包的导入路径的最后一个元素）。
可以使用-o 标志更改文件名。

-exec xprog
使用 xprog 运行测试二进制文件。行为与
'go run'中的行为相同。有关详细信息，请参阅“go help run”。

-i
安装作为测试依赖项的包。
不要运行测试。

-json
将测试输出转换为适合自动处理的 JSON。
有关编码详细信息，请参阅“go doc test2json”。

-o file
将测试二进制文件编译为指定文件。
测试仍然运行（除非指定了-c 或-i）。
测试二进制文件还接受控制测试执行的标志; 这些标志也可以通过'go test'访问。有关详细信息，请参阅“go help testflag”。

有关构建标志的更多信息，请参阅“go help build”。有关指定包的更多信息，请参阅“go help packages”。

另见：go build，go vet。

# <a name="Run specified go tool">运行指定的 go 工具</a>

用法：

go tool [-n]命令[args ...]
Tool 运行由参数标识的 go 工具命令。没有参数，它打印已知工具列表。

-n 标志使工具打印将要执行但不执行它的命令。

有关每个工具命令的更多信息，请参阅“go doc cmd / <command>”。

# <a name="Print Go version">打印 Go 版本</a>

用法：

去版本
版本打印 Go 版本，由 runtime.Version 报告。

# <a name="Report likely mistakes in packages">报告包中可能出现的错误</a>

用法：

go vet [-n][-x] [-vettool prog][build flags] [vet flags][包]
Vet 在导入路径命名的包上运行 Go vet 命令。

有关兽医及其旗帜的更多信息，请参阅“go doc cmd / vet”。有关指定包的更多信息，请参阅“go help packages”。有关检查器及其标志的列表，请参阅“go tool vet help”。有关特定检查器（如“printf”）的详细信息，请参阅“go tool vet help printf”。

-n 标志打印将要执行的命令。-x 标志在执行时打印命令。

-vettool = prog 标志选择具有替代或附加检查的不同分析工具。例如，可以使用以下命令构建和运行'shadow'分析器：

go install
golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow go vet -vettool = \$（which shadow）
go vet 支持的构建标志是控制包解析和执行的构建标志，例如-n，-x，-v，-tags 和-toolexec。有关这些标志的更多信息，请参阅“go help build”。

另见：go fmt，go fix。

# <a name="Build modes">构建模式</a>

'go build'和'go install'命令采用-buildmode 参数，该参数指示要构建哪种对象文件。目前支持的值是：

-buildmode = archive
将列出的非主包构建到.a 文件中。名为
main 的包将被忽略。

-buildmode = c-archive
将列出的主程序包及其导入的所有程序包构建
到 C 归档文件中。唯一可调用的符号将是
使用 cgo // export 注释导出的函数。只需要
列出一个主要包。

-buildmode = c-shared
将列出的主程序包及其导入的所有程序包构建
到 C 共享库中。唯一可调用的符号将
是使用 cgo // export 注释导出的函数。
只需要列出一个主要包。

-buildmode =默认
列出的主程序包内置于可执行文件中，列出的
非主程序包内置于.a 文件中（默认
行为）。

-buildmode = shared 将
所有列出的非主包合并到一个共享
库中，该库将在使用-linkshared
选项构建时使用。名为 main 的包将被忽略。

-buildmode = exe
构建列出的主包及其导入
可执行文件的所有内容。未命名为 main 的包将被忽略。

-buildmode = pie
构建列出的主包及其导入的
位置独立可执行文件（PIE）。未命名为
main 的包将被忽略。

-buildmode =插件
将列出的主要包以及它们
导入的所有包构建到 Go 插件中。未命名为 main 的包将被忽略。

# <a name="Calling between Go and C">在 Go 和 C 之间调用</a>

在 Go 和 C / C ++代码之间调用有两种不同的方法。

第一个是 cgo 工具，它是 Go 发行版的一部分。有关如何使用它的信息，请参阅 cgo 文档（go doc cmd / cgo）。

第二个是 SWIG 程序，它是语言之间接口的通用工具。有关 SWIG 的信息，请参阅 http://swig.org/。运行go build 时，任何扩展名为.swig 的文件都将传递给 SWIG。任何扩展名为.swigcxx 的文件都将使用-c ++选项传递给 SWIG。

当使用 cgo 或 SWIG 时，go build 会将任何.c，.m，.s 或.S 文件传递给 C 编译器，将任何.cc，.cpp，.cxx 文件传递给 C ++编译器。可以设置 CC 或 CXX 环境变 ​​ 量以分别确定要使用的 C 或 C ++编译器。

# <a name="Build and test caching">构建和测试缓存</a>

go 命令缓存构建输出以便在将来的构建中重用。缓存数据的默认位置是当前操作系统的标准用户缓存目录中名为 go-build 的子目录。设置 GOCACHE 环境变量会覆盖此默认值，并且运行“go env GOCACHE”将打印当前缓存目录。

go 命令定期删除最近未使用的缓存数据。运行'go clean -cache'会删除所有缓存的数据。

构建缓存正确地考虑了对 Go 源文件，编译器，编译器选项等的更改：在典型使用中不应该明确清除缓存。但是，构建缓存不会检测使用 cgo 导入的 C 库的更改。如果您对系统上的 C 库进行了更改，则需要显式清理缓存，或者使用-a build 标志（请参阅“go help build”）强制重建依赖于更新的 C 库的包。

go 命令还可以缓存成功的包测试结果。有关详细信息，请参阅“go help test”。运行'go clean -testcache'会删除所有缓存的测试结果（但不会缓存构建结果）。

GODEBUG 环境变量可以打印有关缓存状态的调试信息：

GODEBUG = gocacheverify = 1 导致 go 命令绕过任何缓存条目的使用，而是重建所有内容并检查结果是否与现有缓存条目匹配。

GODEBUG = gocachehash = 1 导致 go 命令打印用于构造缓存查找键的所有内容哈希的输入。输出很大，但可用于调试缓存。

GODEBUG = gocachetest = 1 导致 go 命令打印关于是否重用缓存的测试结果的决定的详细信息。

# <a name="Environment variables">环境变量</a>
go命令及其调用的工具检查几个不同的环境变量。对于其中许多，您可以通过运行'go env NAME'来查看系统的默认值，其中NAME是变量的名称。

通用环境变量：

GCCGO
	运行'go build -compiler = gccgo'的gccgo命令。
GOARCH
	用于编译代码的体系结构或处理器。
	例子是amd64,386，arm，ppc64。
GOBIN'go
	install'将安装命令的目录。
GOCACHE
	go命令将存储缓存
	信息的目录，以便在将来的构建中重用。
GOFLAGS
	当
	当前命令知道给定标志时，默认情况下
	应用于go命令的空格分隔的-flag = value设置列表。命令行
	中列出的标志将在此列表后应用，因此会覆盖它。
GOOS
	编译代码的操作系统。
	例如linux，darwin，windows，netbsd。
GOPATH
	欲了解更多详情，请参阅：'go help gopath'。
	Go模块代理的
GOPROXY URL。请参阅'go help goproxy'。
GORACE
	竞赛探测器的选项。
	请参阅https://golang.org/doc/articles/race_detector.html。
GOROOT
	go树的根。
GOTMPDIR
	go命令将写入
	临时源文件，包和二进制文件的目录。
GOFLAGS列表中的每个条目都必须是独立标志。由于条目是以空格分隔的，因此标志值不得包含空格。

与cgo一起使用的环境变量：

CC
	用于编译C代码的命令。
CGO_ENABLED
	是否支持cgo命令。0或1.
CGO_CFLAGS
	在编译
	C代码时cgo将传递给编译器的标志。
CGO_CFLAGS_ALLOW
	一个正则表达式，指定允许
	出现在#cgo CFLAGS源代码指令中的其他标志。
	不适用于CGO_CFLAGS环境变量。
CGO_CFLAGS_DISALLOW
	一个正则表达式，指定必须禁止
	出现在#cgo CFLAGS源代码指令中的标志。
	不适用于CGO_CFLAGS环境变量。
CGO_CPPFLAGS，CGO_CPPFLAGS_ALLOW，CGO_CPPFLAGS_DISALLOW
	像CGO_CFLAGS，CGO_CFLAGS_ALLOW和CGO_CFLAGS_DISALLOW，
	但是对于C预处理器。
CGO_CXXFLAGS，CGO_CXXFLAGS_ALLOW，CGO_CXXFLAGS_DISALLOW
	与CGO_CFLAGS，CGO_CFLAGS_ALLOW和CGO_CFLAGS_DISALLOW类似，
	但是对于C ++编译器。
CGO_FFLAGS，CGO_FFLAGS_ALLOW，CGO_FFLAGS_DISALLOW
	与CGO_CFLAGS，CGO_CFLAGS_ALLOW和CGO_CFLAGS_DISALLOW类似，
	但对于Fortran编译器。
CGO_LDFLAGS，CGO_LDFLAGS_ALLOW，CGO_LDFLAGS_DISALLOW
	与CGO_CFLAGS，CGO_CFLAGS_ALLOW和CGO_CFLAGS_DISALLOW类似，
	但是对于链接器。
CXX
	用于编译C ++代码的命令。
PKG_CONFIG
	pkg-config工具的路径。
AR
	使用
	gccgo编译器构建时用于操作库归档的命令。
	默认为'ar'。
体系结构特定的环境变量：

GOARM
	对于GOARCH = arm，要编译的ARM体系结构。
	有效值为
5,6,7。GO386
	对于GOARCH = 386，浮点指令集。
	有效值为387，sse2。
GOMIPS
	对于GOARCH = mips {，le}，是否使用浮点指令。
	有效值为hardfloat（默认），softfloat。
GOMIPS64
	对于GOARCH = mips64 {，le}，是否使用浮点指令。
	有效值为hardfloat（默认），softfloat。
专用环境变量：

GCCGOTOOLDIR
	If set, where to find gccgo tools, such as cgo.
	The default is based on how gccgo was configured.
GOROOT_FINAL
	The root of the installed Go tree, when it is
	installed in a location other than where it is built.
	File names in stack traces are rewritten from GOROOT to
	GOROOT_FINAL.
GO_EXTLINK_ENABLED
	Whether the linker should use external linking mode
	when using -linkmode=auto with code that uses cgo.
	Set to 0 to disable external linking mode, 1 to enable it.
GIT_ALLOW_PROTOCOL
	Defined by Git. A colon-separated list of schemes that are allowed to be used
	with git fetch/clone. If set, any scheme not explicitly mentioned will be
	被'去得'认为不安全。
“go env”提供的其他信息，但未从环境中读取：

GOEXE
	可执行文件名后缀（Windows上为“.exe”，其他系统上为“”）。
GOHOSTARCH
	Go工具链二进制文件的体系结构（GOARCH）。
GOHOSTOS
	Go工具链二进制文件的操作系统（GOOS）。
GOMOD
	主模块的go.mod 的绝对路径，
	如果不使用模块则为空字符串。
GOTOOLDIR
	安装go工具（编译，封面，doc等）的目录。
# <a name="File types">文件类型</a>
go命令检查每个目录中受限文件集的内容。它根据文件名的扩展名标识要检查的文件。这些扩展是：

.go
	Go源文件。
.c，.h
	C源文件。
	如果软件包使用cgo或SWIG，这些将使用
	OS本机编译器（通常是gcc）进行编译; 否则会
	触发错误。
.cc，.cpp，.cxx，.hh，.hpp，.hxx
	C ++源文件。仅适用于cgo或SWIG，并且始终
	使用OS本机编译器进行编译。
.m
	Objective-C源文件。仅适用于cgo，并始终
	使用OS本机编译器进行编译。
.s，.S
	汇编源文件。
	如果软件包使用cgo或SWIG，它们将与
	OS本机汇编程序（通常是gcc（sic））组装在一起; 否则他们
	将与Go汇编程序组装在一起。
.swig，.swigcxx
	SWIG定义文件。
.syso
	系统对象文件。
除.syso之外的每个类型的文件都可能包含构建约束，但是go命令会停止扫描文件中第一个不是空行或//样式行注释的构建约束。有关更多详细信息，请参阅go / build包文档。

通过Go 1.12版本，非测试Go源文件还可以包含// go：binary-only-package注释，指示包源仅包含在文档中，不得用于构建包二进制文件。这样就可以单独以编译形式分发Go包。即使是仅二进制包也需要准确的导入块来列出所需的依赖关系，以便在链接生成的命令时可以提供这些依赖关系。请注意，此功能计划在Go 1.12发布后删除。
# <a name="The go.mod file">go.mod 文件</a>
模块版本由源文件树定义，其根目录中包含go.mod文件。当运行go命令时，它会查找当前目录，然后查找连续的父目录，以查找标记主（当前）模块根目录的go.mod。

go.mod文件本身是面向行的，带有//注释但没有/ * * / comments。每行包含一个指令，由一个动词后跟参数组成。例如：

模块my / thing
go 1.12
要求其他/东西v1.0.2
要求new / thing / v2 v2.3.4
排除old / thing v1.2.3
替换bad / thing v1.4.5 => good / thing v1.4.5
动词是

模块，定义模块路径;
去，设置预期的语言版本;
要求，要求给定版本或更高版本的特定模块;
排除，排除特定模块版本的使用; 并
替换，以使用不同的模块版本替换模块版本。
排除和替换仅适用于主模块的go.mod，并在依赖项中被忽略。有关详细信息，请参阅https://research.swtch.com/vgo-mvs。

前导动词可以从相邻行中分解出来以创建一个块，就像在Go导入中一样：

要求（
	new / thing v2.3.4
	old / thing v1.2.3
）
go.mod文件的设计既可以直接编辑，也可以通过工具轻松更新。'go mod edit'命令可用于从程序和工具中解析和编辑go.mod文件。请参阅'go help mod edit'。

go命令每次使用模块图时都会自动更新go.mod，以确保go.mod始终准确地反映现实并且格式正确。例如，考虑这个go.mod文件：

模块M

要求（
        A v1
        B v1.0.0
        C v1.0.0
        D v1.2.3
        E dev
）

排除D v1.2.3
更新将非规范版本标识符重写为semver格式，因此A的v1变为v1.0.0，而E的dev变为dev分支上最新提交的伪版本，可能是v0.0.0-20180523231146-b3f5c0f6e5f1。

更新修改了要求以遵守排除，因此对已排除的D v1.2.3的要求将更新为使用D的下一个可用版本，可能是D v1.2.4或D v1.3.0。

此更新消除了冗余或误导性要求。例如，如果A v1.0.0本身需要B v1.2.0和C v1.0.0，则go.mod对B v1.0.0的要求具有误导性（由A需要v1.2.0取代），并且要求C v1。 0.0是冗余的（A对同一版本的需要暗示），因此两者都将被删除。如果模块M包含直接从B或C导入包的包，那么将保留需求但更新为正在使用的实际版本。

最后，更新以规范格式重新格式化go.mod，以便将来的机械更改将导致最小的差异。

因为模块图定义了import语句的含义，所以加载包的任何命令也都使用并因此更新go.mod，包括go build，go get，go install，go list，go test，go mod graph，go mod tidy，and去mod为什么。
# <a name="GOPATH environment variable">GOPATH 环境变量</a>
Go路径用于解析import语句。它由go / build包实现并记录。

GOPATH环境变量列出了查找Go代码的位置。在Unix上，该值是以冒号分隔的字符串。在Windows上，该值是以分号分隔的字符串。在计划9中，值是一个列表。

如果未设置环境变量，GOPATH默认为用户主目录中名为“go”的子目录（在Unix上为$ HOME / go，在Windows上为％USERPROFILE％\ go），除非该目录包含Go分发。运行“go env GOPATH”查看当前的GOPATH。

请参阅https://golang.org/wiki/SettingGOPATH以设置自定义GOPATH。

GOPATH中列出的每个目录都必须具有规定的结构：

src目录包含源代码。src下面的路径确定导入路径或可执行文件名。

pkg目录包含已安装的包对象。与Go树一样，每个目标操作系统和体系结构对都有自己的子目录pkg（pkg / GOOS_GOARCH）。

如果DIR是GOPATH中列出的目录，则可以将包含DIR / src / foo / bar源的包导入为“foo / bar”，并将其编译形式安装到“DIR / pkg / GOOS_GOARCH / foo / bar.a” ”。

bin目录保存已编译的命令。每个命令都以其源目录命名，但仅以最终元素命名，而不是整个路径。也就是说，DIR / src / foo / quux中带有源的命令安装在DIR / bin / quux中，而不是DIR / bin / foo / quux中。剥离“foo /”前缀，以便您可以将DIR / bin添加到PATH以获取已安装的命令。如果设置了GOBIN环境变量，则命令将安装到它命名的目录而不是DIR / bin。GOBIN必须是绝对的道路。

这是一个示例目录布局：

GOPATH = / home / user / go

/ home / user / go /
    src /
        foo /
            bar /（转到包中的代码）
                x.go
            quux /（转到包main中的代码）
                y.go
    bin /
        quux（安装命令）
    pkg /
        linux_amd64 /
            foo /
                bar.a（已安装的包对象）
Go搜索GOPATH中列出的每个目录以查找源代码，但新包始终下载到列表中的第一个目录中。

有关示例，请参阅https://golang.org/doc/code.html。
# <a name="GOPATH and Modules">GOPATH 和模块</a>
使用模块时，GOPATH不再用于解析导入。但是，它仍然用于存储下载的源代码（在GOPATH / pkg / mod中）和编译的命令（在GOPATH / bin中）。
# <a name="Internal Directories">内部目录</a>
名为“internal”的目录中或下面的代码只能由以“internal”的父目录为根的目录树中的代码导入。这是上面目录布局的扩展版本：

/ home / user / go /
    src /
        crash /
            bang /（go包中的代码）
                b.go
        foo /（go包foo中的代码）
            f.go
            bar /（go包中的代码）
                x.go
            internal /
                baz /（转到包baz中的代码）
                    z.go
            quux /（go package in package main）
                y.go
z.go中的代码导入为“foo / internal / baz”，但该import语句只能出现在以foo为根的子树中的源文件中。源文件foo / f.go，foo / bar / x.go和foo / quux / y.go都可以导入“foo / internal / baz”，但源文件crash / bang / b.go不能。

有关详细信息，请参阅https://golang.org/s/go14internal。
# <a name="Vendor Directories">Vendor 目录</a>
Go 1.6包括支持使用外部依赖项的本地副本来满足这些依赖项的导入，通常称为vendoring。

名为“vendor”的目录下的代码只能由以“vendor”的父目录为根的目录树中的代码导入，并且只能使用省略前缀并包括vendor元素的导入路径。

这是上一节中的示例，但将“internal”目录重命名为“vendor”并添加了新的foo / vendor / crash / bang目录：

/ home / user / go /
    src /
        crash /
            bang /（go包中的代码）
                b.go
        foo /（go包foo中的代码）
            f.go
            bar /（go包中的代码）
                x.go
            vendor /
                crash /
                    bang /（转到包bang中的代码）
                        b.go
                baz /（转到包baz中的代码）
                    z.go
            quux /（go package in package main）
                y.go
相同的可见性规则适用于内部，但z.go中的代码导入为“baz”，而不是“foo / vendor / baz”。

源树中较深的供应商目录中的代码在较高目录中影响代码。在以foo为根的子树中，“崩溃/爆炸”的导入解析为“foo / vendor / crash / bang”，而不是顶级“崩溃/爆炸”。

供应商目录中的代码不受导入路径检查的限制（请参阅“go help importpath”）。

当'go get'检出或更新git存储库时，它现在也会更新子模块。

供应商目录不会影响第一次通过“go get”检出的新存储库的位置：这些存储库始终位于主GOPATH中，而不是位于供应商子树中。

有关详细信息，请参阅https://golang.org/s/go15vendor。
# <a name="Module proxy protocol">模块代理协议</a>
默认情况下，go命令直接从版本控制系统下载模块，就像'go get'一样。GOPROXY环境变量允许进一步控制下载源。如果未设置GOPROXY，是空字符串，或者是字符串“direct”，则下载使用默认的直接连接到版本控制系统。将GOPROXY设置为“off”不允许从任何来源下载模块。否则，GOPROXY应该是模块代理的URL，在这种情况下，go命令将从该代理获取所有模块。无论模块的来源如何，下载的模块必须与go.sum中的现有条目相匹配（有关验证的讨论，请参阅“go help modules”）。

Go模块代理是可以响应对指定表单的URL的GET请求的任何Web服务器。请求没有查询参数，因此即使是从固定文件系统（包括文件：/// URL）提供服务的站点也可以是模块代理。

发送到Go模块代理的GET请求是：

GET $ GOPROXY / <module> / @ v / list返回给定模块的所有已知版本的列表，每行一个。

GET $ GOPROXY / <module> / @ v / <version> .info返回有关给定模块的该版本的JSON格式的元数据。

GET $ GOPROXY / <module> / @ v / <version> .mod返回给定模块的该版本的go.mod文件。

GET $ GOPROXY / <module> / @ v / <version> .zip返回给定模块的该版本的zip存档。

为了避免在区分大小写的文件系统中提供问题，<module>和<version>元素是大小写编码的，用感叹号替换每个大写字母后跟相应的小写字母：github.com/Azure编码为github.com/!azure。

关于给定模块的JSON格式的元数据对应于此Go数据结构，可以在将来进行扩展：

type Info struct {
    Version string // version string
    Time time.Time // commit time
}
给定模块的特定版本的zip存档是标准zip文件，其包含与模块的源代码和相关文件对应的文件树。存档使用斜杠分隔的路径，存档中的每个文件路径必须以<module> @ <version> /开头，其中模块和版本直接替换，而不是大小写编码。模块文件树的根对应于存档中的<module> @ <version> /前缀。

即使直接从版本控制系统下载，go命令也会合成显式的info，mod和zip文件，并将它们存储在本地缓存中，$ GOPATH / pkg / mod / cache / download，就像它直接从下载它们一样代理人。缓存布局与代理URL空间相同，因此在（或复制到）https://example.com/proxy上提供$ GOPATH / pkg / mod / cache / download 会让其他用户访问这些缓存的模块版本GOPROXY = https://example.com/proxy。
# <a name="Import path syntax">导入路径语法</a>
导入路径（请参阅“go help packages”）表示存储在本地文件系统中的包。通常，导入路径表示标准包（例如“unicode / utf8”）或在其中一个工作空间中找到的包（有关详细信息，请参阅：'go help gopath'）。
# <a name="Relative import paths">相对导入路径</a>
以./或../开头的导入路径称为相对路径。工具链以两种方式支持相对导入路径作为快捷方式。

首先，相对路径可以用作命令行上的简写。如果您在包含导入为“unicode”的代码的目录中工作并且想要运行“unicode / utf8”的测试，则可以键入“go test ./utf8”而不是需要指定完整路径。同样，在相反的情况下，“go test ..”将从“unicode / utf8”目录中测试“unicode”。也允许相对模式，例如“go test。/ ...”来测试所有子目录。有关模式语法的详细信息，请参阅“go help packages”。

其次，如果您正在编译不在工作空间中的Go程序，则可以在该程序的import语句中使用相对路径来引用附近的代码，而不是在工作空间中。这样可以很容易地在通常的工作空间之外试验小型多包装程序，但是这些程序不能通过“go install”安装（没有可以安装它们的工作空间），所以每次它们都是从头开始重建的。建成了。为避免歧义，Go程序无法在工作空间中使用相对导入路径。
# <a name="Remote import paths">远程导入路径</a>
某些导入路径还描述了如何使用修订控制系统获取程序包的源代码。

一些常见的代码托管站点具有特殊语法：

到位桶（GIT，水银）

	进口“bitbucket.org/user/project”
	进口“bitbucket.org/user/project/sub/directory”

的GitHub（GIT）

	进口“github.com/user/project”
	进口“github.com/ user / project / sub /目录“

Launchpad（Bazaar）

	import”launchpad.net/project“
	import”launchpad.net/project/series“
	import”launchpad.net/project/series/sub/directory“

	import”launchpad.net/ ~user / project / branch“
	import”launchpad.net/~user/project/branch/sub/directory“

IBM DevOps Services（Git）

	import”hub.jazz.net/git/user/project“
	import”hub.jazz。净/ git的/用户/项目/子/目录”
对于托管在其他服务器上的代码，导入路径可以使用版本控制类型进行限定，或者go工具可以通过https / http动态获取导入路径，并从HTML中的<meta>标记中发现代码所在的位置。

声明代码位置，表单的导入路径

repository.vcs /路径
使用指定的版本控制系统指定具有或不包含.vcs后缀的给定存储库，然后指定该存储库中的路径。支持的版本控制系统是：

Bazaar .bzr
Fossil .fossil
Git .git
Mercurial .hg
Subversion .svn
例如，

导入“example.org/user/foo.hg”
表示example.org/user/foo或foo.hg中的Mercurial存储库的根目录

导入“example.org/repo.git/foo/bar”
表示example.org/repo或repo.git中Git存储库的foo / bar目录。

当版本控制系统支持多种协议时，在下载时依次尝试每种协议。例如，Git下载尝试https：//，然后是git + ssh：//。

默认情况下，下载仅限于已知的安全协议（例如https，ssh）。要覆盖Git下载的此设置，可以设置GIT_ALLOW_PROTOCOL环境变量（有关详细信息，请参阅：“go help environment”）。

如果导入路径不是已知的代码托管站点且缺少版本控制限定符，则go工具会尝试通过https / http获取导入，并在文档的HTML <head>中查找<meta>标记。

元标记具有以下形式：

<meta name =“go-import”content =“import-prefix vcs repo-root”>
import-prefix是与存储库根目录对应的导入路径。它必须是使用“go get”获取的包的前缀或完全匹配。如果它不是完全匹配，则在前缀处生成另一个http请求以验证<meta>标记是否匹配。

元标记应尽可能早地出现在文件中。特别是，它应该出现在任何原始JavaScript或CSS之前，以避免混淆go命令的受限解析器。

vcs是“bzr”，“fossil”，“git”，“hg”，“svn”之一。

repo-root是包含方案且不包含.vcs限定符的版本控制系统的根。

例如，

导入“example.org/pkg/foo”
将导致以下请求：

https://example.org/pkg/foo?go-get=1（首选）
 http://example.org/pkg/foo?go-get=1   （后备，仅限-insecure）
如果该页面包含元标记

<meta name =“go-import”content =“example.org git https://code.org/r/p/exproj ”>
go工具将验证https://example.org/?go-get=1是否包含相同的元标记，然后git clone https://code.org/r/p/exproj进入GOPATH / src / example.org 。

使用GOPATH时，下载的包将写入GOPATH环境变量中列出的第一个目录。（参见'go help gopath-get'和'go help gopath'。）

使用模块时，下载的包存储在模块缓存中。（参见'go help module-get'和'go help goproxy'。）

使用模块时，会识别go-import元标记的其他变体，并且优先于那些列出版本控制系统。该变体使用“mod”作为内容值中的vcs，如：

<meta name =“go-import”content =“example.org mod https://code.org/moduleproxy ”>
此标记表示从URL https://code.org/moduleproxy上提供的模块代理获取带有以example.org开头的路径的模块。有关代理协议的详细信息，请参阅“go help goproxy”。

# <a name="Import path checking">导入路径检查</a>
当上述自定义导入路径功能重定向到已知代码托管站点时，每个生成的包都有两个可能的导入路径，使用自定义域或已知的托管站点。

如果通过对这两种形式之一的注释立即跟随（在下一个换行符之前），则声称包语句具有“导入注释”：

包数学//导入“路径”
包数学/ *导入“路径”* /
go命令将拒绝安装带有导入注释的包，除非该导入路径引用该包。通过这种方式，导入注释可以让包作者确保使用自定义导入路径，而不是直接指向底层代码托管站点的路径。

对供应商树中的代码禁用导入路径检查。这使得可以将代码复制到供应商树中的备用位置，而无需更新导入注释。

使用模块时也会禁用导入路径检查。导入路径注释由go.mod文件的模块语句废弃。

有关详细信息，请参阅https://golang.org/s/go14customimport。
# <a name="Modules, module versions, and more">模块，模块版本等</a>
模块是相关Go包的集合。模块是源代码交换和版本控制的单元。go命令直接支持使用模块，包括记录和解析对其他模块的依赖性。模块替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件。
# <a name="Preliminary module support">初步模块支持</a>
Go 1.11包括对Go模块的初步支持，包括一个新的模块感知'go get'命令。我们打算继续修改这种支持，同时保持兼容性，直到它可以被宣布为官方（不再是初步的），然后在稍后我们可以删除对GOPATH工作的支持和旧的'go get'命令。

利用新的Go 1.11模块支持的最快方法是将您的存储库签出到GOPATH / src之外的目录中，在那里创建一个go.mod文件（在下一节中描述），并从该文件中运行go命令树。

对于更精细的控制，Go 1.11中的模块支持尊重临时环境变量GO111MODULE，该变量可以设置为三个字符串值之一：off，on或auto（默认值）。如果GO111MODULE = off，则go命令从不使用新模块支持。相反，它查找供应商目录和GOPATH以查找依赖项; 我们现在将其称为“GOPATH模式”。如果GO111MODULE = on，则go命令需要使用模块，从不咨询GOPATH。我们将此称为模块感知或以“模块感知模式”运行的命令。如果GO111MODULE = auto或未设置，则go命令根据当前目录启用或禁用模块支持。仅当当前目录位于GOPATH / src之外并且其本身包含go.mod文件或位于包含go的目录下时，才启用模块支持。

在模块感知模式下，GOPATH不再定义构建期间导入的含义，但它仍然存储下载的依赖项（在GOPATH / pkg / mod中）和已安装的命令（在GOPATH / bin中，除非设置了GOBIN）。
# <a name="Defining a module">定义一个模块</a>
模块由Go源文件树定义，并在树的根目录中包含go.mod文件。包含go.mod文件的目录称为模块根目录。通常，模块根目录也将对应于源代码存储库根目录（但通常不需要）。该模块是模块根目录及其子目录中所有Go包的集合，但不包括具有自己的go.mod文件的子树。

“模块路径”是与模块根对应的导入路径前缀。go.mod文件定义模块路径，并列出在构建期间解析导入时应使用的其他模块的特定版本，方法是提供模块路径和版本。

例如，这个go.mod声明包含它的目录是带有路径example.com/m的模块的根目录，并且它还声明该模块依赖于golang.org/x/text和gopkg.in的特定版本。 /yaml.v2：

module example.com/m

require（
	golang.org/x/text
	v0.3.0 gopkg.in/yaml.v2 v2.1.0
）
go.mod文件还可以指定仅在直接构建模块时应用的替换和排除版本; 当模块合并到更大的构建中时，它们将被忽略。有关go.mod文件的更多信息，请参阅“go help go.mod”。

要启动一个新模块，只需在模块目录树的根目录中创建一个go.mod文件，该文件只包含一个模块语句。'go mod init'命令可用于执行此操作：

去mod init example.com/m
在已经使用现有依赖关系管理工具（如godep，glide或dep）的项目中，“go mod init”还将添加与现有配置匹配的require语句。

一旦go.mod文件存在，就不需要额外的步骤：像'go build'，'go test'，甚至'go list'这样的命令将根据需要自动添加新的依赖项以满足导入。
# <a name="The main module and the build list">主模块和构建列表</a>
“主模块”是包含运行go命令的目录的模块。go命令通过查找当前目录中的go.mod或者当前目录的父目录，或者父目录的父目录等来查找模块root。

主模块的go.mod文件定义了go命令可以通过require，replace和exclude语句使用的精确软件包集。通过以下require语句找到的依赖关系模块也有助于定义该组包，但只能通过其go.mod文件的require语句：依赖模块中的任何replace和exclude语句都将被忽略。因此，replace和exclude语句允许主模块完全控制其自己的构建，而不受依赖项的完全控制。

提供构建包的模块集称为“构建列表”。构建列表最初仅包含主模块。然后，go命令以递归方式向列表添加列表中已有模块所需的确切模块版本，直到没有任何内容可添加到列表中。如果将特定模块的多个版本添加到列表中，则最后仅保留最新版本（根据语义版本排序）以用于构建。

'go list'命令提供有关主模块和构建列表的信息。例如：

go list -m＃主模块的打印路径
go list -m -f = {{.dir}} #print主模块的根目录
go list -m all #print build list

# <a name="Maintaining module requirements">维护模块要求</a>
go.mod文件是程序员和工具可读和可编辑的。go命令本身会自动更新go.mod文件，以维护标准格式和require语句的准确性。

任何找到不熟悉的导入的go命令都会查找包含该导入的模块，并自动将该模块的最新版本添加到go.mod中。因此，在大多数情况下，只需在源代码中添加导入并运行“go build”，“go test”，甚至“go list”即可：作为分析包的一部分，go命令将发现并解析导入并更新go.mod文件。

任何go命令都可以确定缺少模块要求并且必须添加，即使仅考虑模块中的单个包也是如此。另一方面，确定不再需要并且可以删除模块要求需要在所有可能的构建配置（体系结构，操作系统，构建标记等）中完整查看模块中的所有包。'go mod tidy'命令构建该视图，然后添加任何缺少的模块要求并删除不必要的模块要求。

作为在go.mod中维护require语句的一部分，go命令跟踪哪些提供由当前模块直接导入的包，哪些提供仅由其他模块依赖性间接使用的包。仅在间接使用时需要的要求在go.mod文件中标有“// indirect”注释。一旦其他直接要求暗示间接要求，就会自动从go.mod文件中删除。间接要求仅在使用未能说明其某些自身依赖关系的模块或在其自己声明的要求之前明确升级模块的依赖关系时出现。

由于这种自动维护，go.mod中的信息是构建的最新可读描述。

'go get'命令更新go.mod以更改构建中使用的模块版本。升级一个模块可能意味着升级其他模块，同样一个模块的降级可能意味着降级其他模块。'go get'命令也会产生这些隐含的变化。如果直接编辑go.mod，“go build”或“go list”等命令将假定升级是预期的，并自动进行任何隐含的升级并更新go.mod以反映它们。

'go mod'命令提供了用于维护和理解模块和go.mod文件的其他功能。请参阅'go help mod'。

-mod build标志提供了对go.mod更新和使用的额外控制。

如果使用-mod = readonly调用，则不允许从上述go.mod的隐式自动更新中执行go命令。相反，当需要对go.mod进行任何更改时，它会失败。此设置对于检查go.mod是否不需要更新非常有用，例如在持续集成和测试系统中。即使使用-mod = readonly，“go get”命令仍然允许更新go.mod，而“go mod”命令不接受-mod标志（或任何其他构建标志）。

如果使用-mod = vendor调用，则go命令假定供应商目录包含正确的依赖项副本，并忽略go.mod中的依赖项描述。
# <a name="Pseudo-versions">伪版本</a>
go.mod文件和go命令通常使用语义版本作为描述模块版本的标准形式，因此可以比较版本以确定哪个版本应该比其他版本更早或更晚。通过在底层源存储库中标记修订版来引入类似v1.2.3的模块版本。可以使用像v0.0.0-yyyymmddhhmmss-abcdefabcdef这样的“伪版本”来引用未标记的修订，其中时间是UTC的提交时间，最后的后缀是提交哈希的前缀。时间部分确保可以比较两个伪版本以确定稍后发生的版本，提交哈希标识基础提交，并且前缀（在此示例中为v0.0.0-）是从提交图中的最新标记版本派生的在此提交之前。

有三种伪版本形式：

当目标提交之前没有具有适当主要版本的早期版本化提交时，将使用vX.0.0-yyyymmddhhmmss-abcdefabcdef。（这最初是唯一的形式，因此一些较旧的go.mod文件甚至可以使用此表单来执行跟随标记的提交。）

当目标提交之前的最新版本化提交是vX.YZ-pre时，使用vX.YZ-pre.0.yyyymmddhhmmss-abcdefabcdef。

当目标提交之前的最新版本化提交是vX.YZ时，使用vX.Y.（Z + 1）-0.yyyymmddhhmmss-abcdefabcdef。

伪版本永远不需要手动输入：go命令将接受普通提交哈希并自动将其转换为伪版本（或标记版本，如果可用）。此转换是模块查询的示例。
# <a name="Module queries">模块查询</a>
go命令在命令行和主模块的go.mod文件中接受“模块查询”来代替模块版本。（在评估主模块的go.mod文件中找到的查询后，go命令会更新文件以将查询替换为其结果。）

完全指定的语义版本（例如“v1.2.3”）将评估该特定版本。

语义版本前缀（例如“v1”或“v1.2”）将评估具有该前缀的最新可用标记版本。

语义版本比较（例如“<v1.2.3”或“> = v1.5.6”）评估最接近比较目标的可用标记版本（<和<=的最新版本，>和>的最早版本=）。

字符串“latest”与最新的可用标记版本匹配，或者与底层源存储库的最新未标记版本匹配。

底层源存储库的修订标识符（例如提交哈希前缀，修订标记或分支名称）选择该特定代码修订。如果修订版还标记了语义版本，则查询将评估该语义版本。否则，查询将评估为提交的伪版本。

所有查询都喜欢发布版本到预发布版本。例如，“<v1.2.3”将更喜欢返回“v1.2.2”而不是“v1.2.3-pre1”，即使“v1.2.3-pre1”更接近比较目标。

主模块go.mod中的exclude语句不允许的模块版本被视为不可用，并且查询无法返回。

例如，这些命令都是有效的：

go get github.com/gorilla/mux@latest#相同（@latest默认为'go get'）
go get github.com/gorilla/mux@v1.6.2#recored v1.6.2
go get github.com/gorilla/ mux @ e3702bed2＃records v1.6.2
go get github.com/gorilla/mux@c856192#recored v0.0.0-20180517173623-c85619274f5d
go get github.com/gorilla/mux@master#dov master的当前含义
# <a name="Module compatibility and semantic versioning">模块兼容性和语义版本控制</a>
go命令要求模块使用语义版本并期望版本准确地描述兼容性：它假定v1.5.4是v1.5.3，v1.4.0甚至v1.0.0的向后兼容替代品。更常见的是，go命令期望包遵循“导入兼容性规则”，其中说：

“如果旧软件包和新软件包具有相同的导入路径，则新软件包必须向后兼容旧软件包。”

由于go命令采用导入兼容性规则，因此模块定义只能设置其依赖项之一的最低要求版本：它无法设置最大值或排除所选版本。但是，导入兼容性规则并不能保证：v1.5.4可能是错误的，而不是v1.5.3的向后兼容替代品。因此，go命令永远不会从旧版本更新到未安装的模块的较新版本。

在语义版本控制中，更改主版本号表示缺少与早期版本的向后兼容性。为了保持导入兼容性，go命令要求主要版本为v2或更高版本的模块使用具有该主要版本的模块路径作为最终元素。例如，example.com/m的v2.0.0版必须使用模块路径example.com/m/v2，该模块中的包将使用该路径作为其导入路径前缀，如example.com/m/v2 /子/ PKG。以这种方式包括模块路径中的主要版本号和导入路径称为“语义导入版本控制”。主要版本为v2及更高版本的模块的伪版本以该主要版本而非v0开头，如v2.0.0-20180326061214-4fc5987536ef。

作为一种特殊情况，以gopkg.in/开头的模块路径继续使用在该系统上建立的约定：主要版本始终存在，并且前面有一个点而不是斜杠：gopkg.in/yaml.v1和gopkg.in/yaml.v2，而不是gopkg.in/yaml和gopkg.in/yaml/v2。

go命令将具有不同模块路径的模块视为不相关：它在example.com/m和example.com/m/v2之间没有任何连接。具有不同主要版本的模块可以在构建中一起使用，并且由于它们的包使用不同的导入路径而保持独立。

在语义版本控制中，主要版本v0用于初始开发，表示没有期望稳定性或向后兼容性。主要版本v0没有出现在模块路径中，因为这些版本是为v1.0.0做准备，并且v1也没有出现在模块路径中。

在引入语义导入版本控制约定之前编写的代码可以使用主要版本v2和更高版本来描述与v0和v1中使用的相同的未版本化导入路径集。为了适应这样的代码，如果源代码存储库对于没有go.mod的文件树具有v2.0.0或更高版本的标记，则该版本被认为是v1模块的可用版本的一部分，并且在转换时被赋予+不兼容的后缀到模块版本，如在v2.0.0 +不兼容。+不兼容标记也适用于从此类版本派生的伪版本，如v2.0.1-0.yyyymmddhhmmss-abcdefabcdef +不兼容。

通常，在v0版本，预发行版本，伪版本或+不兼容版本上的构建列表中具有依赖性（由“go list -m all”报告）表明升级时出现问题的可能性更大这种依赖性，因为没有期望兼容性。

见https://research.swtch.com/vgo-import关于语义进口版本的更多信息，并查看https://semver.org/更多关于语义版本。
# <a name="Module code layout">模块代码布局</a>
现在，请参阅https://research.swtch.com/vgo-module，以获取有关如何将版本控制系统中的源代码映射到模块文件树的信息。
# <a name="Module downloading and verification">环模块下载和验证境变量</a>
go命令在主模块的根目录中与go.mod一起维护一个名为go.sum的文件，其中包含特定模块版本内容的预期加密校验和。每次使用依赖项时，如果缺少，则将其校验和添加到go.sum，或者需要匹配go.sum中的现有条目。

go命令维护下载包的缓存，并在下载时计算和记录每个包的加密校验和。在正常操作中，go命令会针对主模块的go.sum文件检查这些预先计算的校验和，而不是在每次命令调用时重新计算它们。'go mod verify'命令检查模块下载的缓存副本是否仍然匹配记录的校验和和go.sum中的条目。

根据GOPROXY环境变量的设置，go命令可以从代理获取模块，而不是直接连接到源控制系统。

有关代理的详细信息以及缓存的已下载软件包的格式，请参阅“go help goproxy”。
# <a name="Modules and vendoring">模块和销售</a>
使用模块时，go命令完全忽略供应商目录。

默认情况下，go命令通过从其源下载模块并使用下载的副本来满足依赖性（在验证之后，如上一节中所述）。为了允许与旧版本的Go进行互操作，或者为了确保用于构建的所有文件一起存储在单个文件树中，'go mod vendor'在主模块的根目录中创建一个名为vendor的目录并存储在那里来自依赖模块的包，这些包是支持主模块中包的构建和测试所需的。

要使用主模块的顶级供应商目录来构建以满足依赖性（禁用常用网络源和本地缓存的使用），请使用“go build -mod = vendor”。请注意，仅使用主模块的顶级供应商目录; 其他位置的供应商目录仍被忽略。
# <a name="Module-aware go get">模块感知 go get</a>
'go get'命令根据go命令是在模块感知模式还是传统GOPATH模式下运行来改变行为。即使在传统的GOPATH模式下，此帮助文本也可以作为“go help module-get”访问，它描述了“go get”，因为它在模块感知模式下运行。

用法：go get [-d] [-m] [-u] [-v] [-insecure] [build flags] [packages]

获取解析并将依赖项添加到当前开发模块，然后构建并安装它们。

第一步是解决要添加的依赖项。

对于每个命名的包或包模式，get必须决定使用哪个版本的相应模块。默认情况下，get选择最新的标记发行版本，例如v0.4.5或v1.2.3。如果没有标记的发行版本，请选择最新的标记预发布版本，例如v0.0.1-pre1。如果根本没有标记版本，请选择最新的已知提交。

可以通过在package参数中添加@version后缀来覆盖此默认版本选择，如'go get golang.org/x/text@v0.3.0'。对于存储在源控制存储库中的模块，版本后缀也可以是提交哈希，分支标识符或源控制系统已知的其他语法，如'go get golang.org/x/text@master'。版本后缀@latest显式请求上述默认行为。

如果正在考虑的模块已经是当前开发模块的依赖项，那么get将更新所需的版本。指定早于当前所需版本的版本是有效的，并降低依赖性。版本后缀@none表示应根据需要完全删除依赖项，降级或删除模块。

虽然默认使用包含命名包的模块的最新版本，但它不使用该模块的最新版本的依赖项。相反，它更喜欢使用该模块请求的特定依赖版本。例如，如果最新的A需要模块B v1.2.3，而B v1.2.4和v1.3.1也可用，那么'go get A'将使用最新的A但是然后使用B v1.2.3，按照A的要求。（如果对特定模块有竞争要求，那么'go get'通过获取最大请求版本来解决这些要求。）

-u标志指示get更新依赖关系以在可用时使用较新的次要或补丁版本。继续前面的例子，'go get -u A'将使用最新的A与B v1.3.1（不是B v1.2.3）。

-u = patch标志（不是-u patch）指示get更新依赖关系以在可用时使用更新的补丁版本。继续前面的例子，'go get -u = patch A'将使用最新的A和B v1.2.4（不是B v1.2.3）。

通常，添加新的依赖项可能需要升级现有的依赖项以保持工作的构建，并且“go get”会自动执行此操作。同样，降级一个依赖项可能需要降级其他依赖项，“go get”也会自动执行此操作。

在解析，升级和降级模块以及更新go.mod之后，-m标志指示get停在这里。使用-m时，每个指定的包路径也必须是模块路径，而不是模块根目录下的包的导入路径。

-insecure标志允许从存储库中提取并使用不安全的方案（如HTTP）解析自定义域。谨慎使用。

第二步是下载（如果需要），构建和安装命名包。

如果参数命名模块但不命名包（因为模块的根目录中没有Go源代码），则跳过该参数的安装步骤，而不是导致构建失败。例如，即使没有与该导入路径对应的代码，'go get golang.org/x/perf'也会成功。

请注意，允许使用包模式，并在解析模块版本后进行扩展。例如，'go get golang.org/x/perf/cmd / ...'添加最新的golang.org/x/perf，然后在最新版本中安装命令。

-d标志指示get下载构建命名包所需的源代码，包括下载必要的依赖项，但不构建和安装它们。

如果没有包参数，则“go get”将应用于主模块，并应用于当前目录中的Go包（如果有）。特别是，'go get -u'和'go get -u = patch'更新主模块的所有依赖项。没有包参数也没有-u，'go get'不仅仅是'go install'，'go get -d'不仅仅是'go list'。

有关模块的更多信息，请参阅“go help modules”。

有关指定包的更多信息，请参阅“go help packages”。

本文描述了使用模块来管理源代码和依赖关系的行为。如果go命令在GOPATH模式下运行，则get的标志和效果的细节会发生变化，就像'go help get'一样。请参阅'go help modules'和'go help gopath-get'。

另见：go build，go install，go clean，go mod。
# <a name="Package lists and patterns">包列表和模式</a>
许多命令适用于一组包：

去行动[包]
通常，[packages]是导入路径列表。

导入路径是根路径或以a开头的路径。或..元素被解释为文件系统路径，表示该目录中的包。

否则，导入路径P表示在GOPATH环境变量中列出的某些DIR的目录DIR / src / P中找到的包（有关更多详细信息，请参阅：'go help gopath'）。

如果未指定导入路径，则该操作将应用于当前目录中的包。

路径有四个保留名称，不应该用于使用go工具构建的包：

- “main”表示独立可执行文件中的顶级包。

- “all”扩展到所有GOPATH树中的所有包。例如，'go list all'列出本地系统上的所有软件包。使用模块时，“all”扩展到主模块中的所有包及其依赖关系，包括任何这些包的测试所需的依赖关系。

- “std”就像扩展到标准Go库中的包一样。

- “cmd”扩展为Go存储库的命令及其内部库。

以“cmd /”开头的导入路径仅匹配Go存储库中的源代码。

导入路径是一种模式，如果它包含一个或多个“...”通配符，每个通配符都可以匹配任何字符串，包括空字符串和包含斜杠的字符串。这样的模式扩展到GOPATH树中找到的所有包目录，其名称与模式匹配。

为了使普通模式更方便，有两种特殊情况。首先，/ ...在模式的末尾可以匹配一个空字符串，以便net / ...匹配其子目录中的net和packages，如net / http。其次，任何包含通配符的斜杠分隔模式元素都不会参与vendored包路径中“vendor”元素的匹配，因此./ ...与./vendor或./的子目录中的包不匹配。 mycode / vendor，但./vendor / ...和./mycode/vendor / ... do。但请注意，名为vendor的目录本身包含代码不是销售包：cmd / vendor将是名为vendor的命令，并且模式cmd / ...与它匹配。有关vendoring的更多信息，请参阅golang.org/s/go15vendor。

导入路径还可以命名要从远程存储库下载的包。运行'go help importpath'了解详细信息。

程序中的每个包都必须具有唯一的导入路径。按照惯例，这是通过使用属于您的唯一前缀启动每个路径来安排的。例如，Google内部使用的路径都以“google”开头，而表示远程存储库的路径则以代码的路径开头，例如“github.com/user/repo”。

程序中的包不需要具有唯一的包名，但有两个具有特殊含义的保留包名。名称main表示命令，而不是库。命令内置于二进制文件中，无法导入。名称文档表示目录中非Go程序的文档。go命令会忽略包文档中的文件。

作为一种特殊情况，如果包列表是来自单个目录的.go文件列表，则该命令将应用于由这些文件组成的单个合成包，忽略这些文件中的任何构建约束并忽略其中的任何其他文件。目录。

以“。”开头的目录和文件名。go工具忽略或“_”，名为“testdata”的目录也是如此。
# <a name="Testing flags">测试标志</a>
'go test'命令接受适用于'go test'本身的两个标志和适用于生成的测试二进制文件的标志。

几个标志控制分析并编写适合“go tool pprof”的执行配置文件; 运行“go tool pprof -h”获取更多信息。pprof的--alloc_space，--alloc_objects和--show_bytes选项控制信息的呈现方式。

'go test'命令识别以下标志并控制任何测试的执行：

-bench regexp
    仅运行与正则表达式匹配的基准。
    默认情况下，不运行基准测试。
    要运行所有基准测试，请使用'-bench'。或'-bench =。'。
    正则表达式由未括号的斜杠（/）
    字符拆分为正则表达式序列，并且
    基准测试标识符的每个部分必须与
    序列中的相应元素匹配（如果有）。可能的匹配父项
    以bN = 1运行以识别子基准。例如，
    给定-bench = X / Y，匹配X的顶级基准测试
    以bN = 1 运行，以找到与Y匹配的任何子基准，
    然后完全运行。

-benchtime t
    运行每个基准测试的足够迭代以获取t，指定
    为time.Duration（例如，-benchtime 1h30s）。
    默认值为1秒（1秒）。
    特殊语法Nx意味着运行基准N次
    （例如，-benchtime 100x）。

-count n
    运行每个测试和基准n次（默认值1）。
    如果设置了-cpu，则为每个GOMAXPROCS值运行n次。
    示例总是运行一次。

-cover
    启用覆盖率分析。
    请注意，因为覆盖率通过
    在编译之前注释源代码来工作，所以
    启用覆盖率的编译和测试失败可能会报告不对应的行号
    原始来源。

-covermode set，count，atomic
    设置
    正在测试的软件包的覆盖率分析模式。除非启用了-race，否则默认为“set”，
    在这种情况下它是“原子”。
    值：
	set：bool：这个语句运行吗？
	count：int：这个语句运行了多少次？
	atomic：int：count，但在多线程测试中是正确的;
		显着更贵。
    设置 - 覆盖。

-coverpkg pattern1，pattern2，pattern3
    将每个测试中的覆盖率分析应用于与模式匹配的包。
    默认情况是每个测试仅分析正在测试的包。
    有关包模式的说明，请参阅“go help packages”。
    设置 - 覆盖。

-cpu 1,2,4
    指定
    应为其执行测试或基准测试的GOMAXPROCS值列表。默认值
    是GOMAXPROCS 的当前值。

-failfast
    在第一次测试失败后不要开始新的测试。

-list regexp
    列出与正则表达式匹配的测试，基准或示例。
    不会运行测试，基准测试或示例。这只
    列出顶级测试。不会显示子测试或子基准测试。

-parallel n
    允许并行执行调用t.Parallel的测试函数。
    该标志的值是
    同时运行的最大测试数; 默认情况下，它设置为GOMAXPROCS的值。
    请注意，-parallel仅适用于单个测试二进制文件。
    'go test'命令也可以
    根据-p标志的设置并行运行不同包的测试
    （参见'go help build'）。

-run regexp
    仅运行与正则表达式匹配的那些测试和示例。
    对于测试，正则表达式由未括号的斜杠（/）
    字符拆分为正则表达式序列，并且
    测试标识符的每个部分必须与相应的元素匹配。
    顺序，如果有的话。请注意，匹配的可能父项也会
    运行，因此-run = X / Y匹配并运行并报告
    与X匹配的所有测试的结果，即使没有匹配Y的子测试的结果，
    因为它必须运行它们以查找那些-tests。

-short
    告诉长时间运行的测试以缩短其运行时间。
    默认情况下它处于关闭状态，但在all.bash期间设置，以便安装
    Go树可以运行完整性检查但不花时间运行
    详尽的测试。

-timeout d
    如果测试二进制文件的运行时间超过持续时间d，则发生混乱。
    如果d为0，则禁用超时。
    默认值为10分钟（10米）。

-v
    详细输出：记录运行时的所有测试。
    即使测试成功，也会打印Log和Logf调用中的所有文本。

-vet list
    在“go test”期间配置“go vet”的调用，
    以使用以逗号分隔的兽医检查列表。
    如果list为空，则“go test”运行“go vet”，其中列出了一系列
    被认为总是值得解决的检查。
    如果列表是“关闭”，则“go test”根本不会运行“go vet”。
以下标志也可以通过'go test'识别，并可用于在执行期间对测试进行分析：

-benchmem
    打印基准的内存分配统计信息。

-blockprofile block.out
    在所有测试完成后
    ，将goroutine阻塞配置文件写入指定的文件。
    将测试二进制文件写为-c will。

-blockprofilerate n
    通过
    使用n调用runtime.SetBlockProfileRate来控制goroutine阻塞配置文件中提供的详细信息。
    请参阅'go doc runtime.SetBlockProfileRate'。
    分析器的目的是平均每隔
    n纳秒对程序所阻塞的一个阻塞事件进行采样。默认情况下，
    如果设置了-test.blockprofile而没有此标志，
    则会记录所有阻塞事件，相当于-test.blockprofilerate = 1。

-coverprofile cover.out
    在所有测试通过后，将覆盖配置文件写入文件。
    设置 - 覆盖。

-cpuprofile cpu.out
    在退出之前将CPU配置文件写入指定的文件。
    将测试二进制文件写为-c will。

-memprofile mem.out
    在所有测试通过后将分配配置文件写入文件。
    将测试二进制文件写为-c will。

-memprofilerate n
    通过
    设置runtime.MemProfileRate，启用更精确（和昂贵）的内存分配配置文件。请参阅'go doc runtime.MemProfileRate'。
    要分析所有内存分配，请使用-test.memprofilerate = 1。

-mutexprofile mutex.out
    所有测试完成后
    ，将互斥锁争用配置文件写入指定的文件。
    将测试二进制文件写为-c will。

-mutexprofilefraction n
    n堆栈中的样本1，包含
    争用互斥锁的goroutines 。

-outputdir目录
    将分析中的输出文件放在指定目录中，
    默认情况下是运行“go test”的目录。

-trace trace.out
    在退出之前将执行跟踪写入指定的文件。
这些标志中的每一个也通过可选的“测试”识别。前缀，如-test.v. 但是，当直接调用生成的测试二进制文件（'go test -c'的结果）时，前缀是必需的。

在调用测试二进制文件之前，'go test'命令在可选包列表之前和之后，根据需要重写或删除已识别的标志。

例如，命令

go test -v -myflag testdata -cpuprofile = prof.out -x
将编译测试二进制文件，然后运行它

pkg.test -test.v -myflag testdata -test.cpuprofile = prof.out
（-x标志被删除，因为它仅适用于go命令的执行，而不适用于测试本身。）

生成配置文件的测试标志（覆盖范围除外）也会将测试二进制文件保留在pkg.test中，以便在分析配置文件时使用。

当'go test'运行测试二进制文件时，它会从相应软件包的源代码目录中执行。根据测试，在直接调用生成的测试二进制文件时可能需要执行相同操作。

命令行程序包列表（如果存在）必须出现在go test命令未知的任何标志之前。继续上面的例子，包列表必须出现在-myflag之前，但可能出现在-v的两侧。

当'go test'在包列表模式下运行时，'go test'会缓存成功的包测试结果，以避免不必要的重复运行测试。要禁用测试缓存，请使用除可缓存标志之外的任何测试标志或参数。显式禁用测试缓存的惯用方法是使用-count = 1。

要保持测试二进制文件的参数不被解释为已知标志或包名称，请使用-args（请参阅“go help test”），它将命令行的其余部分传递给未解释且未更改的测试二进制文件。

例如，命令

去测试-v -args -x -v
将编译测试二进制文件，然后运行它

pkg.test -test.v -x -v
同样的，

去测试-args数学
将编译测试二进制文件，然后运行它

pkg.test数学
在第一个示例中，-x和第二个-v不变地传递给测试二进制文件，并且对go命令本身没有影响。在第二个示例中，参数math被传递给测试二进制文件，而不是被解释为包列表。
# <a name="Testing functions">测试功能</a>
'go test'命令期望在与测试包对应的“* _test.go”文件中找到测试，基准和示例函数。

一个名为TestXxx的测试函数（其中Xxx不以小写字母开头）并且应该具有签名，

func TestXxx（t * testing.T）{...}
基准函数是名为BenchmarkXxx的函数，应具有签名，

func BenchmarkXxx（b * testing.B）{...}
示例函数类似于测试函数，但不是使用* testing.T来报告成功或失败，而是将输出打印到os.Stdout。如果函数中的最后一个注释以“Output：”开头，那么输出将与注释完全比较（参见下面的示例）。如果最后一条注释以“无序输出：”开头，则将输出与注释进行比较，但忽略行的顺序。没有此类注释的示例已编译但未执行。“Output：”之后没有文本的示例被编译，执行，并且预期不会产生输出。

Godoc显示ExampleXxx的主体以演示函数，常量或变量Xxx的使用。具有接收器类型T或* T的方法M的示例被命名为ExampleT_M。给定函数，常量或变量可能有多个示例，由尾随_xxx区分，其中xxx是不以大写字母开头的后缀。

以下是一个示例示例：

func ExamplePrintln（）{
	Println（“这个例子的输出。”）
	//输出：
	//这个例子的输出。
}
这是另一个忽略输出顺序的例子：

func ExamplePerm（）{
	for _，value：= range Perm（4）{
		fmt.Println（value）
	}

	//无序输出：4
	// 2
	// 1
	// 3
	// 0
}
当整个测试文件包含单个示例函数，至少一个其他函数，类型，变量或常量声明，以及没有测试或基准函数时，它们将作为示例显示。

有关更多信息，请参阅测试包的文档。
# <a name="Subdirectories">子目录</a>
名称	概要
..
```
