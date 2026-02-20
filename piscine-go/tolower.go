package student

func ToLower(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c = c + ('a' - 'A')
		}
		result += string(c)
	}
	return result
}
