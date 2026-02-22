package easy

func hammingWeight260221(n int) int {
	ans := 0
	for n > 0 {
		if n&1 == 1 {
			ans += 1
		}
		n = n >> 1
	}
	return ans
}
