package easy

func sumAndMultiply260707(n int) int64 {
	num := 0
	sum := 0
	multi := 1
	for n != 0 {
		curr := n % 10
		if curr != 0 {
			num += curr * multi
			sum += curr
			multi *= 10
		}
		n /= 10
		// fmt.Println(n, num, sum)
	}
	return int64(num * sum)
}
