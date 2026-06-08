package medium

func pivotArray260608(nums []int, pivot int) []int {
	n := len(nums)
	ans := make([]int, n)
	left, right := 0, n-1
	for _, num := range nums {
		if num < pivot {
			ans[left] = num
			left += 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] > pivot {
			ans[right] = nums[i]
			right -= 1
		}
	}
	for i := left; i <= right; i++ {
		ans[i] = pivot
	}
	return ans
}
