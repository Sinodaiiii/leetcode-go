package hard

func zigZagArrays260623(n int, l int, r int) int {
	mod := 1000000007
	width := r - l + 1
	pre := make([][2]int, width)
	for i := range width {
		pre[i] = [2]int{1, 1}
	}
	for i := 1; i < n; i++ {
		curr := make([][2]int, width)
		l2rG, r2lL := 0, 0
		for j := range width {
			curr[j][0] += l2rG
			l2rG = (l2rG + pre[j][1]) % mod

			curr[width-1-j][1] += r2lL
			r2lL = (r2lL + pre[width-1-j][0]) % mod
		}
		pre = curr
	}
	ans := 0
	for i := range width {
		ans = (ans + pre[i][0] + pre[i][1]) % mod
	}
	return ans
}
