package easy

func mapWordWeights260613(words []string, weights []int) string {
	ans := make([]byte, len(words))
	for i, word := range words {
		curr := 0
		for _, c := range word {
			curr += weights[int(c-'a')]
		}
		ans[i] = byte('a' + 25 - curr%26)
	}
	return string(ans)
}
