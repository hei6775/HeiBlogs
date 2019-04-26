package common

//space O(1)
func RotateSlcesInt(nums []int,k int){
	n := len(nums) //数组长度
	k %= n //如果k>n的情况，则取k/n的余数

}

func reverse(nums []int,start,end int){
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}


