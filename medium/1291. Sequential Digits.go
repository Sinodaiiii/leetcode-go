package medium

func sequentialDigits260713(low int, high int) []int {
	ans := make([]int, 0)
	pivot, gap := 12, 11
	for {
		curr := pivot
		for i := curr % 10; i <= 9; i++ {
			if curr > high {
				return ans
			} else if curr >= low {
				ans = append(ans, curr)
			}
			curr += gap
		}
		pivot = pivot*10 + pivot%10 + 1
		gap = gap*10 + 1
	}
}
