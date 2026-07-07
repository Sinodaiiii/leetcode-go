package medium

import "sort"

func removeCoveredIntervals260706(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] >= intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ans := 0
	rightMax := -1
	for _, interval := range intervals {
		if interval[1] <= rightMax {
			ans += 1
		} else {
			rightMax = interval[1]
		}
	}
	return len(intervals) - ans
}
