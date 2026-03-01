package medium

func minPartitions260301(n string) int {
	ans := '0'
	for _, c := range n {
		ans = max(ans, c)
	}
	return int(ans - '0')
}
