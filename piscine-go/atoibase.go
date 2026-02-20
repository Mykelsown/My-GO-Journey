package student

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return 0
		}
		seen[r] = true
	}

	value := make(map[rune]int)
	for i, r := range base {
		value[r] = i
	}

	result := 0
	baseLen := len(base)

	for _, r := range s {
		digit, ok := value[r]
		if !ok {
			return 0
		}
		result = result*baseLen + digit
	}

	return result
}
