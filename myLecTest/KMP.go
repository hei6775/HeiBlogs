package myLecTest

//func KMP(haystack string, needle string) int {
//	if needle == "" {
//		return 0
//	}
//}

func getNext(p string) []int {
	l := len(p) //模式字符串的长度
	t := make([]int, l)//返回最长前缀后缀数组，初始化都为0
	t[0] = 0
	i, j := 0, 1//i用来记录长度，j是p的当前位置index
	for j < l {
		//如果相等 赋值
		if p[i] == p[j] {
			t[j] = i + 1
			i++
			j++
		} else {
			//当p[i],p[j]不相等时
			//当i不等于0的时候取t[i-1]的值 否则t[j]=0
			if i != 0 {
				i = t[i-1]
			} else {
				t[j] = i
				j++
			}
		}
	}
	return t
}
