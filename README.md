# Algorithm
Principe University Algorithm Course


## 简介
&emsp;&emsp;数据结构、算法的学习之路

## 说明

| Path | Description |
| :------:| -----------|
| Astar   | A星寻路算法实现. |
| Beego | beego小部分源码，主要为log模块和tool模块 |
| DB    | 预留数据库，已含Leaf的mongo数据库模块 |
| Tools | 个人常用工具函数，附带单元测试 |
| Go   | Leaf框架封装的goroutine. |
| Http | Http相关练习 |
| Lecture01    | 普林斯顿大学算法课程一 |
| Lecture02   | 普林斯顿大学算法课程二. |
| Lecture03 | 普林斯顿大学算法课程三 |
| MyLecTest    | 课程练习使用 |
| Recoder | 记录文件 |
| Rf    | Leaf框架的ReadFile模块 |
| Socket | socket练习 |
| Until    | Leaf框架的部分模块，以及常用工具函数 |
| Tree | 树数据结构练习 |
| Tools | 个人使用工具 |
| Ws    | websocket练习 |
| Zk    | zookeeper封装 |
## 记录

&emsp;&emsp;golang中赋值都是复制，如果赋值了一个指针，那我们就复制了一个指针副本。
如果赋值了一个结构体，那我们就复制了一个结构体副本。往函数里传参也是同样的情况。

&emsp;&emsp;但是有一点点不同的是，函数传参：

1、指针传递，传递的是指针的地址，但是形参的地址是另外一个，存储的是实参的地址，修改形参会直接修改实参

2、数组传递，传递的是数组的“值拷贝”，对形参进行操作并不会影响到实参

3、数组名传递，和2相同

4、Slice传递，地址拷贝，传递的是底层数组的内存地址，修改形参实际上会修改实参

5、函数传递

&emsp;&emsp;Golang反射三大定律

1、反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”。

2、反射第二定律：反射可以将“反射类型对象”转换为“接口类型变量”。

3、反射第三定律：如果要修改“反射类型对象”，其值必须是“可写的”（settable）


Golang中byte、string、rune的关系

&emsp;&emsp;首先我们要知道golang的默认编码是utf-8，中文unicode下是占两个字节，在utf-8下占三个字节，而在string底层使用byte数组
存，并且不可改变。直接对中文字符串`len()`操作得出的不一定是真实的长度，这是因为byte等同于int8，常用来处理ascii字符，而rune等于int32，常用来处理unicode和utf-8字符。
想要获取中文的话需要使用rune转换



## Others

> #### 树
- 前序遍历： 根结点 ---> 左子树 ---> 右子树
```go
//1、创建一个栈对象，将根节点入栈；

//2、当栈为非空时，将栈顶结点弹出栈并访问该结点；

//3、对当前访问的非空左孩子结点相继依次访问（不需要入栈），并将访问结点的非空右孩子结点入栈

//4、重复执行步骤②和步骤③，直到栈为空为止
func PreStackTraverse(t *TreeNode){//先根遍历，非递归
	if t != nil {
		S := CreateStack()
		S.Push(t)
		for !S.IsEmpty() {
			T,_ := S.Pop()
			fmt.Printf("%d ",T.Value)
			for T != nil {
				if T.Left != nil {
					fmt.Printf("%d ",T.Left.Value)
				}
				if T.Right != nil {
					S.Push(T.Right)
				}
				T = T.Left
			}
		}
	}
	fmt.Println()
}
```

- 中序遍历：左子树---> 根结点 ---> 右子树

- 后序遍历：左子树 ---> 右子树 ---> 根结点

> #### 进制转换

十进制转其他进制

- 原则1：整数部分与小数部分分别转换；
  
- 原则2：整数部分采用除基数(转换为2进制则每次除2，转换为8进制每次除8，以此类推)取余法，直到商为0，而余数作为转换的结果，第一次除后的余数为最低位，最后一次的余数为最高位；
  
- 原则3：小数部分采用乘基数(转换为2进制则每次乘2，转换为8进制每次乘8，以此类推)取整法，直至乘积为整数或达到控制精度。

```python
725.625D=1011010101.101B
725/2 余数1 商362  0.625*2=1.25 整数部分1 小数部分0.25
362/2 余数0 商181  0.25*2=0.5 整数部分0 小数部分0.5
....               ......
2/2 余数0 商1       0.5*2=1.0 整数部分1 小数部分0
1/2 余数1 商0
```

其他进制转十进制

- 按权展开法，即把各数位乘权的i次方后相加

```python
#二进制转十进制
01011010.01B=0×2^7+1×2^6+0×2^5+1×2^4+1×2^3+0×2^2  +1×2^1+0×2^0+0×2^-1+1×2^-2 = 90.25
```

### 跨域

| URL | Description | 是否允许通信 |
| :------: | -----------| :------: |
| `http://www.d.com/d.js`<br>`http://www.d.com/w.js` | 同一域名下 | 允许 |
| `http://www.d.com/lab/a.js`<br>`http://www.d.com//src/b.js` | 同一域名下不同文件夹 | 允许 |
| `http://www.d.com:3333/a.js`<br>`http://www.d.com:4444/b.js` | 同一域名不同端口 | 不允许 |
| `http://www.d.com/a.js`<br>`http://46.33.22.44/b.js` | 域名和域名对应IP | 不允许 |
| `http://www.d.com/a.js`<br>`http://script.d.com/b.js` | 主域相同，子域不同 | 不允许 |
| `http://www.d.com/a.js`<br>`http://d.com/w.js` | 同一域名,不同二级域名<br>（同上） | 不允许(cookie这种情况下也不允许访问) |
| `http://www.d.com/d.js`<br>`http://www.v.com/w.js` | 不同域名 | 不允许 |

js

```javascript
res.header('Access-Control-Allow-Origin','*')
res.header('Access-Control-Allow-Headers','X-Requested-With,Content-Type')
res.header('Access-Control-Allow-Methods','PUT,POST,GET,DELETE,OPTIONS')
```
golang

```go
	//	Origin := r.Header.Get("Origin")
	//	if Origin != "" {
	//		w.Header().Add("Access-Control-Allow-Origin", "*")
	//		w.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	//		w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type")
	//		w.Header().Add("Access-Control-Allow-Credentials", "true")
	//	}
```
## 等比数列 等差数列

![等差数列](Asset/等差数列求和公式.png)

![等比数列](Asset/等比数列求和公式.png)

## Links

[网络计算机基础篇](https://hit-alibaba.github.io/interview/basic/network/HTTP.html)