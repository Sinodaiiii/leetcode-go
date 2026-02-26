package medium

func numSteps260226(s string) int {
	index := len(s) - 1
	ans := 0
	for s[index] == '0' {
		index -= 1
		ans += 1
	}
	if index == 0 {
		return ans
	}
	ans += 2
	index -= 1
	for index > 0 {
		switch s[index] {
		case '0':
			ans += 2
		case '1':
			ans += 1
		}
		index -= 1
	}
	return ans + 1
}
