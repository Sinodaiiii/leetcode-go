package medium

import "math"

func earliestFinishTime260603(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	getAns := func(firstStart, firstDuration, secondStart, secondDuration []int) int {
		minEnd := math.MaxInt32
		for i := range firstStart {
			minEnd = min(minEnd, firstStart[i]+firstDuration[i])
		}
		ret := math.MaxInt32
		for i := range secondStart {
			if secondStart[i] > minEnd {
				ret = min(ret, secondStart[i]+secondDuration[i])
			} else {
				ret = min(ret, minEnd+secondDuration[i])
			}
		}
		return ret
	}

	return min(getAns(landStartTime, landDuration, waterStartTime, waterDuration), getAns(waterStartTime, waterDuration, landStartTime, landDuration))
}
