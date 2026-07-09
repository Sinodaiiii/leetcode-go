package medium

func pathExistenceQueries260709(n int, nums []int, maxDiff int, queries [][]int) []bool {
	father := make([]int, n)
	for i := range n {
		father[i] = i
	}
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] <= maxDiff {
			father[i] = father[i-1]
		}
	}
	ans := make([]bool, len(queries))
	for i, query := range queries {
		ans[i] = father[query[0]] == father[query[1]]
	}
	return ans
}
