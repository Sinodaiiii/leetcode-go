package medium

func combinationSum260102(candidates []int, target int) [][]int {
	dp := make([][][]int, target+1)
	for _, candidate := range candidates {
		if candidate > target {
			continue
		}
		dp[candidate] = append(dp[candidate], []int{candidate})
		for i := candidate + 1; i <= target; i++ {
			if len(dp[i-candidate]) != 0 {
				for j := 0; j < len(dp[i-candidate]); j++ {
					tmp := make([]int, len(dp[i-candidate][j]))
					copy(tmp, dp[i-candidate][j])
					tmp = append(tmp, candidate)
					dp[i] = append(dp[i], tmp)
				}
			}
		}
	}
	return dp[target]
}

func combinationSum260104(candidates []int, target int) [][]int {
	curr := make([]int, 0)
	sum := 0
	ans := [][]int{}
	n := len(candidates)
	var dfs func(index int)
	dfs = func(index int) {
		if sum == target {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			ans = append(ans, tmp)
			return
		} else if sum > target {
			return
		}
		j := len(curr)
		curr = append(curr, 0)
		for i := index; i < n; i++ {
			sum += candidates[i]
			curr[j] = candidates[i]
			dfs(i)
			sum -= candidates[i]
		}
		curr = curr[:len(curr)-1]
	}

	dfs(0)
	return ans
}

func combinationSum260221(candidates []int, target int) [][]int {
	ans := [][]int{}
	curr := []int{}
	n := len(candidates)
	var dfs func(index, left int)
	dfs = func(index, left int) {
		// fmt.Println(index, left, curr)
		if left == 0 {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			ans = append(ans, tmp)
			return
		}
		if index == n || left < 0 {
			return
		}
		for i := index; i < n; i++ {
			curr = append(curr, candidates[i])
			dfs(i, left-candidates[i])
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, target)
	return ans
}
