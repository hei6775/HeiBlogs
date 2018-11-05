# Learning

>## 设计模式

### 单例模式 
>>### golang
``` golang
package main
import (
	"sync"
)

type singleton struct{}

var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

```
>>### python
``` python
 class MySimple(type){
     _instances = {}
        def __call__(cls, *args, **kwargs):
            if cls not in cls._instances:
                cls._instances[cls] = super(MySimple,cls).__call__(*args,**kwargs)
            return cls._instances[cls]
}
#Python2
class MyClass(object):
    __metaclass__ = MySimple
#Python3
class MyClass(metaclass=MySimple):
    pass
```

### 生产-消费者模式
>>### golang
```golang
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//生产者
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

//消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	ch := make(chan int, 64)
	go Producer(2, ch)
	go Producer(5, ch)
	go Consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%#v)\n", <-sig)
}
```

--------
# **Wait For Do**
<p align="right">---hht</p>
<table>
    <tr>
        <th>项目</th>
        <th>开始时间</th>
        <th>结束时间</th>
    </tr>
    <tr>
        <th bgcolor=grey><font color="pink">MongDB</font></th>
        <th>2018-07-24</th>
        <th>2018-07-31</th>
    </tr>
    <tr>
        <th bgcolor=grey><font color="pink">Docker</font></th>
        <th>2018-07-24</th>
        <th>2018-08-15</th>
    </tr>
    <tr>
        <th bgcolor=grey><font color="pink">Shell</font></th>
        <th>2018-07-24</th>
        <th>2018-08-25</th>
    </tr>
    <tr>
        <th bgcolor=grey><font color="pink">MarkDown</font></th>
        <th>2018-07-24</th>
        <th>2018-08-31</th>
    </tr>
        <th bgcolor=grey><font color="pink">Golang</font></th>
        <th>2018-07-24</th>
        <th>2018-10-24</th>
    <tr>
        <th bgcolor=grey><font color="pink">Golang</font></th>
        <th>2018-07-24</th>
        <th>2018-10-24</th>
    </tr>
    <tr>
        <th bgcolor=grey><font color="pink">gRPC</font></th>
        <th>2018-08-05</th>
        <th>2018-08-20</th>
    </tr>
        <tr>
        <th bgcolor=grey><font color="pink">protobuf</font></th>
        <th>2018-08-24</th>
        <th>2018-10-30</th>
    </tr>
    <tr>
        <th bgcolor=grey><font color="pink">Python</font></th>
        <th>2018-08-24</th>
        <th>2018-10-30</th>
    </tr>
</table>
---------


>>second
>>>third

~~hello~~

    int i
    good

---

yiji
===

二级
---

**粗体**

*斜体*

第一段

第二段  
第三段

这是一行中的``代`码``块

[百度](http://www.baidu.com)

[谷歌][1]

[1]:google.com




