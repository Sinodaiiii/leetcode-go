package medium

import "sort"

func maximumElementAfterDecrementingAndRearranging260628(arr []int) int {
	sort.Ints(arr)
	minNum := 1
	ans := 0
	for _, num := range arr {
		if num-minNum+1 > ans {
			ans += 1
		}
	}
	// fmt.Println(arr)
	return minNum + ans - 1
}