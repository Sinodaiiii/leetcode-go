package easy

func gcdOfOddEvenSums260715(n int) int {
	a := (1 + n) * n
	b := a - n
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
