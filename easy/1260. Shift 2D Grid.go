package easy

func shiftGrid260720(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		newColumn := (i + k) % n
		rowOffset := (i + k) / n
		for j := 0; j < m; j++ {
			ans[(j+rowOffset)%m][newColumn] = grid[j][i]
		}
	}
	return ans
}
