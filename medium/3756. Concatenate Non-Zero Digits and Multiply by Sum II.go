package medium

func sumAndMultiply260708(s string, queries [][]int) []int {
	n := len(s)
	index, preSum := make([]int, n), make([]int, n)
	sum := 0
	num := make([]byte, 0, n)
	for i, c := range s {
		index[i] = len(num)
		sum += int(c - '0')
		preSum[i] = sum
		if c != '0' {
			num = append(num, byte(c))
		}
	}

	m := len(num)
	pow10 := make([]int, m+1)
	pow10[0] = 1
	pref := make([]int, m+1)
	mod := 1000000007
	for i := 0; i < m; i++ {
		pow10[i+1] = pow10[i] * 10 % mod
		pref[i+1] = (pref[i]*10 + int(num[i]-'0')) % mod
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		currSum := preSum[r] - preSum[l] + int(s[l]-'0')
		startIdx := index[l]
		endIdx := index[r]
		if s[r] == '0' {
			endIdx--
		}
		currNum := 0
		if startIdx <= endIdx {
			length := endIdx - startIdx + 1
			currNum = (pref[endIdx+1] - pref[startIdx]*pow10[length]%mod + mod) % mod
		}

		ans[i] = currSum * currNum % mod
	}
	return ans
}
