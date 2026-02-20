package student

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return nb
	}

	var rec func(int) int
	rec = func(i int) int {
		if i*i == nb {
			return i
		}
		if i*i > nb {
			return 0
		}
		return rec(i + 1)
	}

	return rec(2)
}
