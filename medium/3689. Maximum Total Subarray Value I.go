package medium

import "math"

func maxTotalValue260609(nums []int, k int) int64 {
	minNum, maxNum := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		minNum = min(minNum, num)
		maxNum = max(maxNum, num)
	}
	return int64(k * (maxNum - minNum))
}
