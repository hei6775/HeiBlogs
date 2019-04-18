# Integration Testing in Go: Part I - Executing Tests with Docker

<p align="right">Author: George ShawMarch 18, 2019</p>

原文地址 ：[Integration Testing in Go: Part I - Executing Tests with Docker](https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html)

## Introduction

**“Testing leads to failure, and failure leads to understanding.” - Burt Rutan**

Burt Rutan 是一名航空航天工程师，他设计了 Voyager，这是第一架在不停车或加油的情况下环球飞行的飞机。
虽然 Rutan 不是软件工程师，但他的话充分说明了测试的重要性，甚至测试软件。
所有形式的测试软件都非常重要，无论是单元，集成，系统还是验收测试。
但是，根据项目，一种形式的测试可能比其他形式更有价值。
换句话说，有时一种形式的测试可以比其他形式更好地理解软件的健康和完整性。

在开发 Web 服务时，我相信一组强大的集成测试可以比其他类型的测试更好地理解服务。
集成测试是一种软件测试形式，用于测试代码与应用程序利用的依赖项（如数据库和消息传递系统）之间的交互。
如果没有集成测试，很难信任 Web 服务的端到端操作。
我相信这是真的，因为在 Web 服务中测试的各个代码单元很少提供与集成测试相同的洞察力。

这是关于 Go 中集成测试的三部分系列的第一篇文章。
本系列中分享的想法，代码和流程旨在轻松扩展到您正在处理的 Web 服务项目。
在这篇文章中，
我将向您展示如何设置 Web 服务项目以使用 Docker 和 Docker Compose 在没有预先安装 Go 的限制性计算环境中运行 Go 测试和依赖项。

## Why use Docker and Docker Compose

吸引许多开发人员使用 Docker 的方法是如何在主机上加载应用程序，而无需手动安装和管理它们。
这意味着您可以加载复杂的软件，包括但不限于数据库（例如 Postgres），消息传递系统（例如 Kafka）和监控系统（例如 Prometheus）。
所有这一切都是通过下载一组代表应用程序及其所有依赖项的镜像来完成的。

注意：[有关容器的更多信息](https://www.docker.com/resources/what-container)，Docker 有一个专门用于定义容器的网页，并突出显示容器和虚拟机之间的差异和相似之处。

Docker Compose 是一种编排工具，可帮助在一个沙箱内构建，运行和联网一组容器。
使用一个命令`docker-compose up`，您可以使 Docker Compose 文件变为现实。
在撰写文件中定义的所有服务将成为在其自己的网络沙箱中作为一个组运行的容器，并按配置运行
。这与手动构建，运行和联网每个容器形成对比，以便允许它们一起运行，相互通信并保留数据。

注意：有关 Docker Compose 的更多信息，请访问网页以获取 Docker 官方网站上 Docker Compose 的概述。

Docker 和 Docker Compose 的另一大好处是，它们有助于在将新开发人员引入项目时实现更轻松的过渡。
新开发人员只需执行一些 Docker 和 Docker Compose 命令即可开始使用，而不是有关于如何安装和管理开发环境的复杂文档。
如果在启动应用程序时主机上当前不存在所需的映像，Docker CLI 将负责下载所需的映像。

## Using Docker and Docker Compose to Run Tests

本系列中引用的 Web 服务应用程序使用 Postgres 数据库公开了一个简单的基于`CRUD`的`REST`风格的 API。
该项目使用 Docker 为生产和测试运行 Postgres 数据库。
此应用程序的测试需要能够在已安装 Go 的本地开发环境和 Go 不存在的受限环境中运行。

以下`Docker Compose`文件支持在上面提到的两种环境中为项目运行集成测试的能力。在本节中，我将分解我选择的配置选项以及我选择它们的原因。

### Listing 1

```
version: '3'

networks:
  integration-tests-example-test:
    driver: bridge

services:
  listd_tests:
    build:
      context: .
      dockerfile: ./cmd/listd/deploy/Dockerfile.test
    depends_on:
      - db
    networks:
      - integration-tests-example-test
  db:
    image: postgres:11.1
    ports:
      - "5432:5432"
    expose:
      - "5432"
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: root
      POSTGRES_DB: testdb
    restart: on-failure
    networks:
      - integration-tests-example-test

```

在 Listing 1 中，您可以看到`Docker Compose`文件，
该文件定义了运行测试所需的项目服务。
此文件有三个主要属性：版本（`version`），网络（`networks`）和服务（`services`）。
`version`属性定义了您正在使用的 Docker Compose 的版本。
`networks`属性定义了可用于给定服务的一个或多个网络配置。
`services`属性定义要启动的容器及其配置。

### Listing 2

```
networks:
  integration-tests-example:
    driver: bridge
```

通过将服务定义在一个`compose`文件中，
默认情况下它们会自动放置在同一网络中，
因此可以相互通信。
但是，与使用默认网络相比，为服务创建网络是最佳做法。
最上层的网络配置定义网络的名称及其使用的驱动程序，在这种情况下是桥接驱动程序。

桥接驱动程序是`Docker`提供的默认驱动程序，
它为容器内部通信创建一个私有内部网络。
服务被告知在`compose`文件中的服务定义配置中使用创建的网络。

### Listing 3

```
services:
  listd_tests:
    build:
      context: .
      dockerfile: ./cmd/listd/deploy/Dockerfile.test
// ... omitted code…
  db:
// ... omitted code…
```

`services`属性有两个直接子属性，`listd_tests`和`db`。`listd_tests`容器通过指定`Dockerfile`来定义其镜像。`context`属性表示所有主机路径应该相对于当前工作目录，如 a`..`表示的那样。

### Listing 4

```
listd_tests:
    build:
      context: .
      dockerfile: ./cmd/listd/deploy/Dockerfile.test
    depends_on:
      - db
    volumes:
      - $PWD:/go/src/github.com/george-e-shaw-iv/integration-tests-example
```

`depends_on`属性告诉`listd_tests`服务等待启动，直到`db`服务已经启动。
除了断言服务的启动顺序之外，此属性还将禁止`listd_tests`服务独立于 db 服务运行。
`volumes`属性告诉 compose 将当前工作目录（由`$ PWD`（打印工作目录）表示）挂载到容器中的`/go/src/github.com/george-e-shaw-iv/integration-tests-example`，是代码将被定位和测试的地方。

### Listing 5

```
listd_tests:
    build:
      context: .
      dockerfile: ./cmd/listd/deploy/Dockerfile.test
    depends_on:
      - db
    networks:
      - integration-tests-example-test
```

最后，该服务被赋予一个网络，以便在沙箱内运行时进行通信。这最初是在 Listing 2 中的顶级网络配置键中定义的。

### Listing 6

```
db:
    image: postgres:11.1
```

下一个服务定义在 db 中的容器通过使用 Docker Hub 托管的映像定义其镜像，即`postgres：1.11` 镜像。如果 Docker Hub 镜像存储库无法在本地计算机上找到镜像，那么 Doc​​ker CLI 非常智能，可以查看它。

### Listing 7

```
db:
    image: postgres:11.1
    ports:
      - "5432:5432"
```

出于安全原因，默认情况下，无法从主机访问任何容器端口。
在本地运行集成测试时，这被证明是一个问题，因为如果无法访问集成服务，那么它就毫无价值。
`ports`属性以下列格式定义从主机到容器的端口映射：`“HOST_PORT：CONTAINER_PORT”`。
Listing 7 中的上述定义可确保计算机上的端口 5432 映射到`db`容器上的端口 5432，因为默认情况下，这是`Postgres`在容器中运行的端口。

### Listing 8

```
db:
    image: postgres:11.1
    ports:
      - "5432:5432"
    expose:
      - "5432"
```

默认情况下，容器的端口未向主机开发，，
同样的容器端口默认也不会暴露给在联网沙箱中运行的其它容器。
即使它们位于同一网络上也是如此。
为了将端口公开给在网络沙箱中运行的其他容器，需要设置暴露端口`expose`的设置。

注意：对于`postgres：1.11`镜像，由于创建镜像的人已经暴露了端口 5432。
由于您不知道镜像是否是使用已暴露的端口创建的，除非您查看镜像的 Dockerfile，
因此最好定义公开`expose`密钥，即使它是多余的。

### Listing 9

```
db:
    image: postgres:11.1
    ports:
      - "5432:5432"
    expose:
      - "5432"
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: root
      POSTGRES_DB: testdb
    restart: on-failure
    networks:
      - integration-tests-example-test
```

db 所需要的最终配置选项是环境（`environment`），重新启动（`restart`）和网络（`networks`）。
`networks`密钥被赋予已定义网络的名称，与先前的服务定义不同。
重新启动密钥的值为`on-failure`，以确保服务在执行期间的任何时候失败时将自动重启。
`environment`选项可以接收环境变量列表，然后在容器的`shell`中设置。
流行应用程序（如 postgres）的大多数托管镜像都具有可以指定的环境变量，以配置镜像提供的应用程序。

## Running The Tests

准备好`Docker Compose`文件后，
下一步是根据`listd_tests`服务中引用的`dockerfile`构建镜像。
此`dockerfile`定义了一个能够运行整个服务的集成测试的镜像。
创建镜像后，即可运行测试。

## Building an Image Capable of Running Tests

为了构建能够运行测试的镜像，必须在`dockerfile`中定义四件事：

获取安装了最新稳定版 Go 的基本映镜像。为 Go 模块安装 git。将可测试代码复制到容器中。运行测试。

让我们打破这些步骤并分析`dockerfile`执行它们所需的指令。

### Listing 10

```
FROM golang:1.12-alpine
```

Listing 10 显示了第 4 步。我选择作为基本操作系统映像的镜像是`golang：1.11-alpine`。镜像在撰写此博客文章时预先安装了 Go 的最新稳定版本。

### Listing 11

```

FROM golang:1.11-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git
```

由于 Alpine OS 非常轻量级，因此必须在基本 Alpine 镜像之上手动安装`git`依赖项。
Listing 11 显示了第 2 步，其中 `git`被添加到镜像中以便使用`Go`模块。
在添加 git 之前运行`apk update`命令以确保安装了最新版本的`git`。
如果您的项目恰好使用 `cgo`，那么您还必须手动安装`gcc`及其所需的库。

### Listing 12

```
FROM golang:1.12-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /go/src/github.com/george-e-shaw-iv/integration-tests-example/
```

为了便于使用，在 Listing 12 中，`WORKDIR`指令设置为`/go/src/github.com/george-e-shaw-iv/integration-tests-example/`，以便其余指令将相对于该目录，这是在容器的`$GOPATH`内。由于在 Listing 4 中安装了带有可测试代码的卷这一事实，已经处理了将可测试代码复制到容器中的过程的第 3 步。

### Listing 13

```
FROM golang:1.12-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /go/src/github.com/george-e-shaw-iv/integration-tests-example/

CMD CGO_ENABLED=0 go test ./...
```

最后，Listing 13 显示了步骤 4，运行测试。

这是通过使用带有 CMD 指令的 `go test ./ ...`来完成的。

测试使用`CGO_ENABLED = 0`作为内联环境变量运行，因为示例项目中的测试不使用`cgo`，而 the alpine 基于镜像不附带 C 编译器。即使您的项目中没有`cgo`代码，也必须以这种方式禁用`cgo`，因为如果启用了`cgo`，Go 仍会尝试使用标准 C 库来执行某些网络任务。

注意：可以在此处找到定义能够从其中运行 Go 测试的自定义镜像的整个`Dockerfile`的代码。现在编写了定义镜像的 `dockerfile`，下面的`Docker Compose 命`令可以调出`listd_test`和`db`服务，这些服务将运行所有集成测试并报告结果。

### Listing 14

```
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```

`--abort-on-container-exit`这个标志是必要的，因为如果省略该标志，则在测试完成运行后，包含集成服务的其他容器将挂起。

## Clean-up

### Listing 15

```
test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes
```

停止和删除容器，卷和网络是一个非常重要的步骤，在运行测试后经常被忽略。
弄清楚为什么由于上次测试运行中持续存在的数据而导致测试被破坏的原因是一个不容易避免的小错误。
为了防止这种情况发生，我创建了一个简单的`makefile`规则，
`test`，在 Listing 14 中展示，用于构建，运行和拆卸容器，无需任何人为干预。

### Listing 16

```
test-db-up:
	docker-compose -f docker-compose.test.yml up --build db

test-db-down:
	docker-compose -f docker-compose.test.yml down --volumes db
```

清单 15 中的规则在受限制的环境中最有效，因为它们在`Compose`文件中启动了两个服务。
为了实现本地测试的相同效果，
可以在运行任何集成测试之前使用 Listing 16 中的`test-db-up`规则，并在运行所有测试之后使用`testdb-down`。

## Conclusion

在这篇文章中，我向您展示了如何设置 Web 服务项目以使用 Docker 和 Docker Compose。我审查的文件允许您在没有预先安装 Go 的限制性计算环境中运行 Go 测试和依赖项。在本系列的下一部分中，我将展示为 Web 服务设置测试套件所需的 Go 代码，这将是编写富有洞察力的集成测试的基础。

注意：这整个系列的帖子从此[存储库](https://github.com/george-e-shaw-iv/integration-tests-example)中提取其示例。
