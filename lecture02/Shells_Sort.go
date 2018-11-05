package lecture02

//Shell's Sort Algorithms
func Shell_Sort(args []int)([]int){
	gap := len(args)/2
	for gap >=1 {
		right := 0+gap
		for i:=0;right<len(args);right=i+gap{
			if args[i] > args[right]{
				temp := args[i]
				args[i] = args[right]
				args[right]= temp
			}
			i++
		}
		gap /= 2
	}
	return args
}