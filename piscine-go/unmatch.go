package student

func Unmatch(a []int) int {
	var quant int
	for _, un := range a {
		quant = 0
		for _, v := range a {
			if v == un {
				quant++
			}
		}
		if quant%2 != 0 {
			return un
		}
	}
	return -1
}
