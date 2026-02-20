package student

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}

	return max
}
