package medium

func processStr260616(s string) string {
	curr := make([]byte, 0, 10000)
	for _, c := range s {
		switch c {
		case '*':
			if len(curr) > 0 {
				curr = curr[:len(curr)-1]
			}
		case '#':
			n := len(curr)
			for i := 0; i < n; i++ {
				curr = append(curr, curr[i])
			}
		case '%':
			n := len(curr) - 1
			if n < 0 {
				continue
			}
			for i := 0; i <= n/2; i++ {
				curr[i], curr[n-i] = curr[n-i], curr[i]
			}
		default:
			curr = append(curr, byte(c))
		}
		// fmt.Println(c, curr)
	}
	return string(curr)
}
