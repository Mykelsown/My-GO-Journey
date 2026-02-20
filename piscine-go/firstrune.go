package student

func FirstRune(s string) rune {
	if s == "" {
		return 0
	}
	for _, r := range s {
		return r
	}
	return 0
}
