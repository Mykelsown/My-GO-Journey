package student

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	word := ""
	for _, r := range str {
		if r == ' ' {
			result[word]++
			word = ""
		} else {
			word += string(r)
		}
	}
	result[word]++

	return result
}
