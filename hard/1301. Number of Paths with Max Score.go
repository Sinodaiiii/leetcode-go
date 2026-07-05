package hard

func pathsWithMaxScore260705(board []string) []int {
	n := len(board)
	mod := 1000000007
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		curr := make([][2]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				curr[j] = [2]int{1, 1}
			} else if board[i][j] == 'X' {
				continue
			} else {
				var l, u, lu [2]int
				if i > 0 && j > 0 {
					lu = dp[j-1]
				}
				if i > 0 {
					u = dp[j]
				}
				if j > 0 {
					l = curr[j-1]
				}
				scoreMax := max(lu[0], l[0], u[0])
				if scoreMax == 0 {
					continue
				} else {
					curr[j][0] = (scoreMax + int(board[i][j]-'0')) % mod
					if lu[0] == scoreMax {
						curr[j][1] += lu[1]
					}
					if u[0] == scoreMax {
						curr[j][1] += u[1]
					}
					if l[0] == scoreMax {
						curr[j][1] += l[1]
					}
					curr[j][1] = curr[j][1] % mod
				}
			}
		}
		dp = curr
	}
	if dp[n-1][0] == 0 {
		return []int{0, 0}
	}
	return []int{(dp[n-1][0] + mod - 1 - 35) % mod, dp[n-1][1]}
}
