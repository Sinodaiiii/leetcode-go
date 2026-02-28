package medium

func concatenatedBinary260228(n int) int {
	mod := 1000000007
	ans := 0
	curr := 0
	next := 1
	for i := 1; i <= n; i++ {
		if i == next {
			next *= 2
			curr += 1
		}
		ans = (ans<<curr + i) % mod
	}
	return ans
}
