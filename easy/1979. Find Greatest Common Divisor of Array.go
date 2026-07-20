package easy

import "math"

func findGCD260718(nums []int) int {
	minNum := math.MaxInt
	maxNum := -1
	for _, num := range nums {
		minNum = min(minNum, num)
		maxNum = max(maxNum, num)
	}
	for minNum != 0 {
		maxNum, minNum = minNum, maxNum%minNum
	}
	return maxNum
}
