package lecture02

func BubbleAsort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		for j := i+1; j < len(v); j++ {
			if  v[i]>v[j]{
				v[i],v[j] = v[j],v[i]
			}
		}
	}
}

func BubbleZsort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		for j := i+1; j < len(v); j++ {
			if  v[i]<v[j]{
				v[i],v[j] = v[j],v[i]
			}
		}
	}
}
