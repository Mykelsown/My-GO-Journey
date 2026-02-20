package student

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	i := 0
	sign := 1

	// Handle optional sign
	if s[0] == '+' {
		i++
	} else if s[0] == '-' {
		sign = -1
		i++
	}

	result := 0

	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}

	return sign * result
}
