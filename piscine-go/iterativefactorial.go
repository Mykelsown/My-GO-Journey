package student

func IterativeFactorial(n int) int {
	if n < 0 || n > 20 {
		return 0
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}
