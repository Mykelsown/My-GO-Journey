package student

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	result := ""
	for i, s := range args {
		if i == 0 {
			result = s
		} else {
			result += "\n" + s
		}
	}
	return result
}
