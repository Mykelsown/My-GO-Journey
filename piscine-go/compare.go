package student

func Compare(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	min := len(ra)
	if len(rb) < min {
		min = len(rb)
	}
	for i := 0; i < min; i++ {
		if ra[i] < rb[i] {
			return -1
		} else if ra[i] > rb[i] {
			return 1
		}
	}
	if len(ra) < len(rb) {
		return -1
	} else if len(ra) > len(rb) {
		return 1
	}
	return 0
}
