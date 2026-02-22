package easy

func binaryGap260222(n int) int {
	ans, pre := 0, -1
	curr := 0
	for n > 0 {
		if n%2 == 1 {
			if pre != -1 {
				ans = max(ans, curr-pre)
			}
			pre = curr
		}
		curr += 1
		n = n >> 1
	}
	return ans
}
