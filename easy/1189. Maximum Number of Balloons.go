package easy

func maxNumberOfBalloons260622(text string) int {
	dict := make([]int, 5)
	for _, c := range text {
		switch c {
		case 'b':
			dict[0] += 1
		case 'a':
			dict[1] += 1
		case 'l':
			dict[2] += 1
		case 'o':
			dict[3] += 1
		case 'n':
			dict[4] += 1
		}
	}
	return min(dict[0], dict[1], dict[2]/2, dict[3]/2, dict[4])
}
