package myLecTest

//func KMP(haystack string, needle string) int {
//	if needle == "" {
//		return 0
//	}
//}

func getNext(p string) []int {
	l := len(p)
	t := make([]int, l)
	t[0] = 0
	i, j := 0, 1
	for j < l {
		if p[i] == p[j] {
			t[j] = i + 1
			i++
			j++
		} else {
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
