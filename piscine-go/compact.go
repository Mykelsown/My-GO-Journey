package student

func Compact(ptr *[]string) int {
	a := *ptr
	j := 0

	for _, v := range a {
		if v != "" {
			a[j] = v
			j++
		}
	}

	*ptr = a[:j]
	return j
}
