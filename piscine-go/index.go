package student

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}

	lenS := len(s)
	lenF := len(toFind)

	if lenF > lenS {
		return -1
	}

	for i := 0; i <= lenS-lenF; i++ {
		j := 0
		for ; j < lenF; j++ {
			if s[i+j] != toFind[j] {
				break
			}
		}
		if j == lenF {
			return i
		}
	}

	return -1
}
