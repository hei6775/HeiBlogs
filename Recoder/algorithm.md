## 有趣的算法题

#### LeetCode 155题 Min Stack

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

- push(x) -- Push element x onto stack.
- pop() -- Removes the element on top of the stack.
- top() -- Get the top element.
- getMin() -- Retrieve the minimum element in the stack.

Example:
```bash
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
```

#### 一种崭新的实现方案

```go
package main

type Stack struct {
	num int // the current x
	min int // the current mininum
}

type MinStack struct {
	Items []Stack
}

func Constructor() MinStack {
	return MinStack{[]Stack{}}
}

func (this *MinStack) Push(x int) {
	if len(this.Items) == 0 {
		this.Items = append(this.Items, Stack{x, x})
	} else {
		lastMin := this.Items[len(this.Items)-1].min
		if x < lastMin {
			this.Items = append(this.Items, Stack{x, x})
		} else {
			this.Items = append(this.Items, Stack{x, lastMin})
		}
	}
}

func (this *MinStack) Pop() {
	this.Items = this.Items[:len(this.Items)-1]
}

func (this *MinStack) Top() int {
	return this.Items[len(this.Items)-1].num
}

func (this *MinStack) GetMin() int {
	return this.Items[len(this.Items)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
```