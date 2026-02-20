package student

func PodiumPosition(podium [][]string) [][]string {
	if len(podium) < 4 {
		return podium
	}

	result := [][]string{
		{podium[3][0]},
		{podium[2][0]},
		{podium[1][0]},
		{podium[0][0]},
	}

	return result
}
