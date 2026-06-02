package main

import (
	"math"
	"sort"
)

func earliestFinishTime260602(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	m, n := len(landStartTime), len(waterStartTime)
	list := make([]int, m)
	findLE := func(start, duration []int, target int) int {
		ret := -1
		l, r := 0, len(start)-1
		for l <= r {
			mid := (l + r) / 2
			i := list[mid]
			if start[i]+duration[i] <= target {
				ret = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return ret
	}

	for i := range list {
		list[i] = i
	}
	sort.Slice(list, func(i, j int) bool {
		indexI, indexJ := list[i], list[j]
		return landStartTime[indexI]+landDuration[indexI] <= landStartTime[indexJ]+landDuration[indexJ]
	})
	ans := math.MaxInt32
	for i := range waterStartTime {
		index := findLE(landStartTime, landDuration, waterStartTime[i])
		if index == -1 {
			ans = min(ans, landStartTime[list[0]]+landDuration[list[0]]+waterDuration[i])
		} else {
			ans = min(ans, waterStartTime[i]+waterDuration[i])
		}
	}

	list = make([]int, n)
	for i := range list {
		list[i] = i
	}
	sort.Slice(list, func(i, j int) bool {
		indexI, indexJ := list[i], list[j]
		return waterStartTime[indexI]+waterDuration[indexI] <= waterStartTime[indexJ]+waterDuration[indexJ]
	})
	for i := range landStartTime {
		index := findLE(waterStartTime, waterDuration, landStartTime[i])
		if index == -1 {
			ans = min(ans, waterStartTime[list[0]]+waterDuration[list[0]]+landDuration[i])
		} else {
			ans = min(ans, landStartTime[i]+landDuration[i])
		}
	}
	return ans
}
