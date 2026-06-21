package medium

import "sort"

func maxIceCream260621(costs []int, coins int) int {
	sort.Ints(costs)
	ans := 0
	for _, cost := range costs {
		if coins < cost {
			return ans
		}
		coins -= cost
		ans += 1
	}
	return ans
}
