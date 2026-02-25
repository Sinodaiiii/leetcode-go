package easy

func sortByBits260225(arr []int) []int {
	check1 := func(num int) (ret int) {
		for num != 0 {
			ret += num & 1
			num = num >> 1
		}
		return ret
	}
	var qs func(l, r int)
	qs = func(l, r int) {
		// fmt.Println(l, r, arr)
		if l >= r {
			return
		}
		pivot := arr[l]
		pivotC := check1(pivot)
		// fmt.Println(pivot, pivotC)
		i, j := l-1, r+1
		for i < j {
			i += 1
			for {
				iC := check1(arr[i])
				if iC < pivotC || iC == pivotC && arr[i] < pivot {
					i += 1
				} else {
					break
				}
			}
			j -= 1
			for {
				jC := check1(arr[j])
				if jC > pivotC || jC == pivotC && arr[j] > pivot {
					j -= 1
				} else {
					break
				}
			}
			// fmt.Println(i, j)
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		qs(l, j)
		qs(j+1, r)
	}
	qs(0, len(arr)-1)
	return arr
}
