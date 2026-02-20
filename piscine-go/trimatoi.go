package student

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	signApplied := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		if !signApplied && c == '-' {
			sign = -1
			signApplied = true
			continue
		}

		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
			signApplied = true
		}
	}

	if num == 0 {
		return 0
	}
	return num * sign
}
