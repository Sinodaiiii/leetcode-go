package easy

func leftRightDifference260606(nums []int) []int {
	ans := make([]int, len(nums))
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	total := 0
	for _, num := range nums {
		total += num
	}
	pre := 0
	for i, num := range nums {
		ans[i] = abs(total - num - 2*pre)
		pre += num
	}
	return ans
}
