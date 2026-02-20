package student

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb > 20 { // optional, prevent stack overflow
		return 0
	}
	return nb * RecursiveFactorial(nb-1)
}
