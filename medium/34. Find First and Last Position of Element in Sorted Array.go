package medium

import "fmt"

func searchRange260221(nums []int, target int) []int {
	getLT := func() int {
		l, r := 0, len(nums)-1
		pivot := -1
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] < target {
				pivot = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return pivot + 1
	}
	getGT := func() int {
		l, r := 0, len(nums)-1
		pivot := len(nums)
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] > target {
				pivot = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return pivot - 1
	}
	l, r := getLT(), getGT()
	fmt.Println(l, r)
	if l > r {
		return []int{-1, -1}
	}
	return []int{l, r}
}
