package myLecTest

//KMP算法
func KMP(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	//获取next数组
	next := getNext(needle)
	toreturn := -1

	i, j := 0, 0
	//haystack肯定比needle长
	// j模式串的当前位置，i字符串中 与模式串开始匹配的位置的头部索引
	for i <= (len(haystack) - len(needle)) {
		//当j<模式串的长度 且当前index的j相等的话 j++
		for (j < len(needle)) && (needle[j] == haystack[i+j]) {
			j++
		}
		//否则  j等于0则i++ 匹配下一个字符
		//如果匹配完成，则返回字符串的头部索引 i
		//否则 i位移j-next[j-1] j位移next[j]
		if j == 0 {
			i++
		} else {
			if j == len(needle) {
				return i
			}
			i = i + j - next[j-1]
			j = next[j-1]
		}
	}
	return toreturn
}

func getNext(p string) []int {
	l := len(p)         //模式字符串的长度
	t := make([]int, l) //返回最长前缀后缀数组，初始化都为0
	t[0] = 0
	i, j := 0, 1 //i用来记录长度，j是p的当前位置index
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
