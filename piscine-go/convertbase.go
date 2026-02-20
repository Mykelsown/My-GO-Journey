package student

func ConvertBase(nbr, baseFrom, baseTo string) string {
	if nbr == "" {
		return ""
	}

	// Convert from baseFrom to decimal
	dec := 0
	baseFromLen := len(baseFrom)
	for _, r := range nbr {
		for i, br := range baseFrom {
			if r == br {
				dec = dec*baseFromLen + i
				break
			}
		}
	}

	// Convert from decimal to baseTo
	if dec == 0 {
		return string(baseTo[0])
	}

	baseToLen := len(baseTo)
	result := ""
	for dec > 0 {
		result = string(baseTo[dec%baseToLen]) + result
		dec /= baseToLen
	}

	return result
}
