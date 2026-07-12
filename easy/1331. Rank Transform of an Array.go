package easy

import "sort"

func arrayRankTransform260712(arr []int) []int {
	n := len(arr)
	tmp := make([][2]int, n)
	for i := range n {
		tmp[i] = [2]int{arr[i], i}
	}
	sort.Slice(tmp, func(i, j int) bool { return tmp[i][0] <= tmp[j][0] })
	index := 1
	for i, num := range tmp {
		if i > 0 && num[0] > tmp[i-1][0] {
			index += 1
		}
		arr[num[1]] = index
	}
	return arr
}
