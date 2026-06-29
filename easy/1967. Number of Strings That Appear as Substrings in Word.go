package easy

func numOfStrings260629(patterns []string, word string) int {
	ans := 0
	n := len(word)
	for _, pattern := range patterns {
		currN := len(pattern)
		for i := 0; i <= n-currN; i++ {
			if pattern == word[i:i+currN] {
				ans += 1
				break
			}
		}
	}
	return ans
}
