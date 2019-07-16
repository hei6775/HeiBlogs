# LeetCode

## 目录

* [189.Rotate Array](#189.Rotate Array)
* [193. Valid Phone Numbers](#193.-Valid Phone Numbers)
* [198. House Robber](#198.-House Robber)

#### 189.Rotate Array

如果 k>n

```golang
//author:abbycoding
//comments:genius
//such as
// [1,2,3,4,5]  k=2
// output [4,5,1,2,3]
// n=5, k = k % n = 2,
//reverse(nums,0,4)  [1,2,3,4,5] ==> [5,4,3,2,1]
//reverse(nums,0,1)  [5,4,3,2,1] ==> [4,5,3,2,1]
//reverse(nums,2,4)  [4,5,3,2,1] ==> [4,5,1,2,3]

func rotate(nums []int, k int) {
    n := len(nums) //数组长度
	k %= n //如果k>n的情况，则取k/n的余数
    reverse(nums, 0, n - 1)
    reverse(nums, 0, k - 1)
    reverse(nums, k, n - 1)
}

//反转 交换
func reverse(nums []int, start int, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
```

#### 193. Valid Phone Numbers

file.txt 是每行都是一串的数字，
从 file.txt 选出合法的电话号码，电话号码的格式为：

(xxx) xxx-xxxx or xxx-xxx-xxxx.

```shell
grep -P '^(\d{3}-|\(\d{3}\) )\d{3}-\d{4}$' file.txt
#-r参数开启扩展正则模式，-n只打印被sed处理的行
sed -n -r '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/p' file.txt

awk '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/' file.txt
```

#### 198. House Robber

在这道题了解到动态规划（`dynamic-programming`）

```golang
func rob(nums []int) int {
    prevMax := 0
    currMax := 0

    for i:=0; i < len(nums); i++ {
        temp:=currMax
        if prevMax + nums[i] > currMax {
            currMax = prevMax + nums[i]
        }
        prevMax = temp
    }
    return currMax
}
// input [1,2,3,1]
// output 4

// temp = currMax = 0
// 0+1 = 1 > 0
// currMax = 1
// prevMax = 0

// temp = currMax = 1
// 0 + 2 > 1
// currMax = 2
// prevMax = 1

// temp = currMax = 2
// 1 + 3 = 4 > 2
// currMax = 4
// prevMax = 2

// temp = currMax = 4
// 2 + 1 > 4
// prevMax = 4
// return 4

```