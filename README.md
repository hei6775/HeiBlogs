# Algorithm
Principe University Algorithm Course


## 简介
&emsp;&emsp;数据结构、算法的学习之路

## 说明

| Path | Description |
| :------:| -----------|
| Astar   | A星寻路算法实现. |
| beego | beego小部分源码，主要为log模块和tool模块 |
| DB    | 预留数据库，已含Leaf的mongo数据库模块 |
| go   | Leaf框架封装的goroutine. |
| htp | Http相关练习 |
| lecture01    | 普林斯顿大学算法课程一 |
| lecture02   | 普林斯顿大学算法课程二. |
| lecture03 | 普林斯顿大学算法课程三 |
| myLecTest    | 课程练习使用 |
| recoder | 记录文件 |
| rf    | Leaf框架的ReadFile模块 |
| socket | socket练习 |
| until    | Leaf框架的部分模块，以及常用工具函数 |
| tree | 树数据结构练习 |
| ws    | websocket练习 |
| zk    | zookeeper封装 |
## 记录

&emsp;&emsp;golang中赋值都是复制，如果赋值了一个指针，那我们就复制了一个指针副本。
如果赋值了一个结构体，那我们就复制了一个结构体副本。往函数里传参也是同样的情况。

&emsp;&emsp;但是有一点点不同的是，函数传参：

1、指针传递，传递的是指针的地址，但是形参的地址是另外一个，存储的是实参的地址，修改形参会直接修改实参

2、数组传递，传递的是数组的“值拷贝”，对形参进行操作并不会影响到实参

3、数组名传递，和2相同

4、Slice传递，地址拷贝，传递的是底层数组的内存地址，修改形参实际上会修改实参

5、函数传递

## Others

- 前序遍历： 根结点 ---> 左子树 ---> 右子树

- 中序遍历：左子树---> 根结点 ---> 右子树

- 后序遍历：左子树 ---> 右子树 ---> 根结点

## Links

待补充 