package student

func Capitalize(s string) string {
	result := ""
	newWord := true

	for i := 0; i < len(s); i++ {
		c := s[i]

		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			if newWord && c >= 'a' && c <= 'z' {
				c = c - ('a' - 'A') // capitalize first letter
			} else if !newWord && c >= 'A' && c <= 'Z' {
				c = c + ('a' - 'A') // lowercase non-first letters
			}
			newWord = false
		} else {
			newWord = true
		}

		result += string(c)
	}

	return result
}
