package student

func SortWordArr(a []string) {
	n := len(a)
	if n < 2 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
