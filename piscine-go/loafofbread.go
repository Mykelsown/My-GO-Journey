package student

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	res := ""
	count := 0
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if count == 5 {
			res += " "
			count = 0
			continue
		}

		if runes[i] == ' ' {
			continue
		}

		res += string(runes[i])
		count++
	}

	if len(res) > 0 && res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}

	return res + "\n"
}
