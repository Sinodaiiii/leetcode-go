package easy

func countPrimeSetBits260221(left int, right int) int {
	count := func(i int) int {
		ret := 0
		for i > 0 {
			if i&1 == 1 {
				ret += 1
			}
			i = i >> 1
		}
		return ret
	}
	ans := 0
	for i := left; i <= right; i++ {
		curr := count(i)
		// fmt.Println(i, curr)
		if curr == 2 || curr == 3 || curr == 5 || curr == 7 || curr == 11 || curr == 13 || curr == 17 || curr == 19 {
			ans += 1
		}
	}
	return ans
}

// 1 2 3 5 7 11 13 17
