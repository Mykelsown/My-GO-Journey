package student

func Split(s, sep string) []string {
	var result []string
	if sep == "" {
		return []string{s}
	}

	sepLen := len(sep)
	word := ""

	for i := 0; i < len(s); {
		// check if sep matches at position i
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, word)
			word = ""
			i += sepLen
		} else {
			word += string(s[i])
			i++
		}
	}

	result = append(result, word)
	return result
}
