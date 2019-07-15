package Lecture02

//Shell's Sort Algorithms
func Shell_Sort(args []int) []int {
	gap := 1
	for gap < (len(args) / 3) {
		gap = 3*gap + 1 //1,4,13,40,121,364,1093
	}
	for gap >= 1 {
		//1
		for i := gap; i < len(args); i++ {
			for j := i; j >= gap && args[j] < args[j-gap]; j -= gap {
				args[j], args[j-gap] = args[j-gap], args[j]
			}
		}
		gap /= 3
	}
	return args
}
